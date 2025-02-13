import React, { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";

const Paying = () => {
    const navigate = useNavigate();
    const jwtToken = localStorage.getItem("accessToken")
    const id = useParams();
    const idString = id.id.toString();
    const [seats, setSeats] = useState([]);
    const [prices, setPrices] = useState([]);
    const [totalPrice, setTotalPrice] = useState(0);
    const servCharge = 6

    useEffect(() => {
        fetch("http://localhost/res/booking/" + idString, {
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
            setTotalPrice(total)
          })
          .catch((error) => console.error("Error getting reservation:", error));
      }, [idString]);

    const handleConfirmPayment = () => {
        let jsonData = {
            email: "testing@gmail.com",
            phone_number: "+71234567890",
        }

        fetch ("http://localhost/res/booking/" + idString, {
            method: "PUT",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwtToken, 
              },
            body: JSON.stringify(jsonData),
        })
            .then((response) => response.json())
            .then(() => navigate("/paid/" + idString))
            .catch((error) => console.error("Error paying reservation:", error));
    };

    return (
        <>
            <div className="background-wrapper"></div>
            
            <div className="container booking-container">
                <div className="order-info">
                    <h1 className="title">Booking Detail</h1>

                    <div className="schedule-section">
                        <h2>Schedule</h2>
                        <p><strong>Movie Title</strong></p>
                        <p className="movie-title">SPIDERMAN NO WAY HOME</p>
                        
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
                            <p>Service Charge (6%)</p>
                            <p>{(totalPrice * servCharge) / 100} Tg </p>
                        </div>
                        <hr></hr>
                        <div className="total-payment">
                            <p><strong>Total payment</strong></p>
                            <p className="total-amount">{totalPrice + ((totalPrice * servCharge) / 100)} Tg</p>
                        </div>
                        <p className="note">*Purchased ticket cannot be canceled</p>
                    </div>

                    <button onClick={handleConfirmPayment} className="checkout-btn">Checkout Ticket</button>
                </div>
            </div>   
        </>
    );
};

export default Paying;