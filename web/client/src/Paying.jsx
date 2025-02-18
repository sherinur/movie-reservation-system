import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";

const Paying = () => {
    const navigate = useNavigate();
    const jwtToken = localStorage.getItem("accessToken")
    const reservationId = useParams().id.toString();

    const [sessionId, setSessionId] = useState("")
    const [seats, setSeats] = useState([]);

    const [prices, setPrices] = useState([]);
    const [sum, setSum] = useState(0);
    const servCharge = 6
    const [charge, setCharge] = useState(0)
    const [totalPrice, setTotalPrice] = useState(0)

    const [title, setTitle] = useState("");
    const [selectedPayment, setSelectedPayment] = useState("credit-card");
    const [promoCode, setPromoCode] = useState("");

    useEffect(() => {
        fetch("http://localhost/res/booking/" + reservationId, {
            method: "GET",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwtToken, 
            },
        })
          .then((response) => response.json())
          .then((data) => {
            let tempSeats = []
            let tempPrices = []
            let total = 0

            for (let ticket of data.Tickets) {
                tempSeats.push(ticket.seat)
                tempPrices.push(ticket.price)
                total += ticket.price
            }

            setSeats(tempSeats)
            setPrices(tempPrices)
            setSum(total)
            setSessionId(data.ScreeningID)
            setTitle(data.MovieTitle)
          })
          .catch((error) => console.error("Error getting reservation:", error));
      }, [reservationId, jwtToken]);

    useEffect(() => {
        setCharge((sum * servCharge) / 100)
        setTotalPrice(sum + charge)
    }, [charge, sum])

    const handleConfirmPayment = () => {
        let jsonData = {
            email: "testing@gmail.com",
            phone_number: "+71234567890",
            total_price: totalPrice,
        }

        fetch ("http://localhost/res/booking/" + reservationId, {
            method: "PUT",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwtToken, 
              },
            body: JSON.stringify(jsonData),
        })
            .catch((error) => console.error("Error paying reservation:", error));

        
        let jsonData2 = {
            "reservation_id": reservationId,
            "payment_price": totalPrice,
            "payment_method": selectedPayment,
            "status": "completed",
        }

        fetch("http://localhost/res/payments/", {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwtToken, 
            },
            body: JSON.stringify(jsonData2),
        })
            .then((response) => response.json())
            .then(() => navigate("/paid/" + reservationId))
            .catch((error) => console.error("Error paying reservation:", error));

    };


    const handleCancelPayment = () => {
        fetch("http://localhost/res/booking/delete/" + reservationId, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwtToken,
            },
        })
            .then((response) => response.json())
            .then(() => navigate("/booking/" + sessionId))
            .catch((error) => console.error("Error cancelling reservation:", error));
    }

    const handlePaymentChange = (method) => {
        setSelectedPayment(method);
    };

    const handleActivePromo = (discount) => {
        console.log("Before update:", totalPrice); 
        setTotalPrice((prevPrice) => {
            const newPrice = prevPrice - (prevPrice * discount) / 100;
            console.log("Updated price:", newPrice); 
            return newPrice;
        });
    }

    return (
        <>
            <div className="background-wrapper"></div>
            
            <div className="container booking-container">

                <div className="order-info">
                    <h1 className="title">Booking Detail</h1>

                    <div className="schedule-section">
                        <h2>Schedule</h2>
                        <p><strong>Movie Title</strong></p>
                        <p className="movie-title">{title}</p>
                        
                        <p><strong>Date</strong></p>
                        <p className="movie-date">Mon, 23 Oct 2023</p>

                        <div className="ticket-info">
                            <div className="seats">
                                <strong>Tickets ({seats.length})</strong>
                                <p>{seats.join(", ") || "None"}</p>
                            </div>
                            <div className="time">
                                <strong>Hours</strong>
                                <p>14:40</p>
                            </div>
                        </div>
                    </div>

                    <div className="transaction-detail">
                        <h2>Transaction Detail</h2>

                        {prices.map((price, i) => (
                            <div 
                            key={i}
                            className="price-row">
                                <p>REGULAR SEAT</p>
                                <p>{price} Tg </p>
                            </div>
                        ))}
                        <div className="price-row">
                            <p>Service Charge ({servCharge}%)</p>
                            <p>{charge} Tg </p>
                        </div>
                        <hr></hr>
                        <div className="total-payment">
                            <p><strong>Total payment</strong></p>
                            <p className="total-amount">{totalPrice.toFixed(2)} Tg</p>
                        </div>
                    </div>

                    <div className="payment-method my-5">
                        <h3>Select Payment Method</h3>
                        <div className="payment-options">
                        {["credit-card", "paypal", "bank-transfer"].map((method) => (
                            <label
                            key={method}
                            className={`payment-option ${selectedPayment === method ? "selected" : ""}`}
                            >
                            <input
                                type="radio"
                                name="payment"
                                value={method}
                                checked={selectedPayment === method}
                                onChange={() => handlePaymentChange(method)}
                            />
                            {method === "credit-card" ? "Credit Card" : 
                            method === "paypal" ? "PayPal" : 
                            "Bank Transfer"}
                            </label>
                        ))}
                        </div>
                    </div>

                    <div className="promo-code">
                        <input
                        type="text"
                        placeholder="#00000"
                        value={promoCode}
                        onChange={(e) => setPromoCode(e.target.value)}
                        />
                        <button onClick={() => handleActivePromo(10)} className="apply-btn">Apply</button>
                    </div>

                    <button onClick={handleCancelPayment} className="back-btn">Cancel</button>
                    <button onClick={handleConfirmPayment} className="checkout-btn">Checkout Ticket</button>
                    <p className="note">*Purchased ticket cannot be canceled</p>
                </div>
            </div>   
        </>
    );
};

export default Paying;