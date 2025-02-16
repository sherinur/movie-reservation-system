import React, { useState, useEffect } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";


const Paid = () => {
    const navigate = useNavigate();
    const jwtToken = localStorage.getItem("accessToken")
    const id = useParams();
    const idString = id.id.toString();
    const [seats, setSeats] = useState([]);
    const [prices, setPrices] = useState([]);
    const [totalPrice, setTotalPrice] = useState(0);

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

    return (
        <>
            <div className="background-wrapper"></div>
            
            <div className="container paid-container">
                <div className="reservation-info">
                    <h1 className="title">Ticket Detail</h1>

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

                    <button className="checkout-btn">Download Ticket</button>
                    <Link to="/"><button className="back-btn">Back to Movies</button></Link>
                </div>
                
            </div>
        </>
    );
};

export default Paid;