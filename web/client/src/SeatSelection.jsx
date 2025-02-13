import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";


const SeatSelection = () => {
    const navigate = useNavigate();
    const jwtToken = localStorage.getItem("accessToken")
    const [seats, setSeats] = useState([]);
    const [bookedSeats, setBookedSeats] = useState([]);
    const [selectedSeats, setSelectedSeats] = useState([]);
    const [groupedSeats, setGroupedSeats] = useState({});

    // Fetch seat data from server
    useEffect(() => {
        fetch("http://localhost/movi/session/pLdVhypfRGCA9K4X7zy9eg==", {
            method: "GET",
            headers: { 
                "Content-Type": "application/json",
            },
        })
            .then((response) => response.json())
            .then((data) => {
            let seat = []

            for (let ticket of data.seats) {
                seat.push(ticket.row)
            }

            setSeats(seat)
            })
            .catch((error) => console.error("Error getting reservation:", error));
        }, []);

    useEffect(() => {
        // Group seats by their letter and number them
        const groups = {}
        seats.forEach((seat) => {
        if (!groups[seat]) {
            groups[seat] = [];
        }
        groups[seat].push(`${seat}${groups[seat].length + 1}`);
        });
        
        setGroupedSeats(groups)
        console.log(jwtToken)
    }, [seats]);

    // Toggle seat selection
    const toggleSeat = (seat) => {
        if (bookedSeats.includes(seat)) return; // Prevent selecting booked seats
        setSelectedSeats((prev) =>
        prev.includes(seat) ? prev.filter((s) => s !== seat) : [...prev, seat]
        );
    };

    // Handle payment (send selected seats to backend)
    const handleStartPayment = () => {
        if (selectedSeats.length === 0) {
        alert("Please, choose seats to reserve")
        return
        }
        let jsonData = {
        screening_id: "679bg09j53ae5cc94c021c5d", 
        tickets: []
        }

        for (let ticket of selectedSeats) {
        jsonData.tickets.push({
            seat: ticket,
            price: 1500,
            seat_type: "common",
            user_type: "adult"
        })}

        fetch("http://localhost/res/booking/", {
        method: "POST",
        headers: { 
            "Content-Type": "application/json",
            "Authorization": "Bearer " + jwtToken, 
        },
        body: JSON.stringify(jsonData),
        })
        .then((res) => res.json())
        .then((data) => navigate("/booking/" + data.InsertedID))
        .catch((error) => console.error("Error confirming seats:", error));
    };

  return (
        <>
            <div className="background-wrapper"></div>

            <div className="seat-selection-container">
                <h1 className="title">Seat Selection</h1>

                {/* Screen */}
                <div className="screen">SCREEN</div>

                <div className="hall">
                    {Object.entries(groupedSeats).map(([row, seatNumbers]) => (
                        <div key={row} className="seat-map">
                            {seatNumbers.map((seat) => (
                                <div onClick={() => toggleSeat(seat)} key={seat} className="seat">
                                    {seat}
                                </div>
                            ))}
                        </div>
                    ))}
                </div>
                
                <div className="seat-selection-footer">
                    <div className="total-price">
                        <p>Total:</p>
                        <strong>4500 Tg</strong>
                    </div>
                    <div className="selected-seats">
                      <p>Seats: </p>
                      <strong>{selectedSeats.join(", ") || "None"}</strong>
                    </div>
                    <div className="action-buttons">
                        <a href="booking.html"><button  className="btn-cancel">Back</button></a>
                        <button onClick={handleStartPayment} className="btn-proceed">Proceed Payment</button>
                    </div>
                </div>
            </div>    
        </>
    );
};

export default SeatSelection;