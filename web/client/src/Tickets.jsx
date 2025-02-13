import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";


const Tickets = () => {
    const jwtToken = localStorage.getItem("accessToken")
    const [seats, setSeats] = useState([]);
    const [reservations, setReservations] = useState([]);

    // Fetch seat data from server
    useEffect(() => {
        fetch("http://localhost/res/booking/", {
            method: "GET",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": "Bearer " + jwtToken, 
            },
        })
            .then((response) => response.json())
            .then((data) => {
                console.log(data)
                let tickets = []

                for (let reservation of data) {
                    tickets.push(reservation)
                }

                setReservations(tickets)
            })
            .catch((error) => console.error("Error getting reservation:", error));
        }, []);

  return (
        <>
            <div className="background-wrapper"></div>

            <div className="container">
                <div className="row profile-section">
                
                    <div className="col-lg-12 col-md-6 col-sm-12">
                        <ul className="nav nav-pills justify-content-center">
                            <li className="nav-item">
                            <a className="nav-link active" id="upcoming-tab" data-bs-toggle="pill" href="#upcoming">Upcoming</a>
                            </li>
                            <li className="nav-item">
                            <a className="nav-link" id="history-tab" data-bs-toggle="pill" href="#history">History</a>
                            </li>
                        </ul>
                    </div>

                    <div className="col-lg-12 col-md-6 col-sm-12">
                        <div className="tab-content mt-3">
                            {/* Upcoming Tab */}
                            <div className="col-lg-12 col-sm-12 tab-pane fade show active" id="upcoming">
                                <div className="row">
                                    {reservations.map((reservation) => (
                                    <div key={reservation.id} className="col-lg-4 col-md-4 col-sm-6 mb-4">
                                        <div className="card">
                                            <div className="card-body">
                                                <p className="card-text"><strong>Date</strong></p>
                                                <p className="card-text">2025-02-14</p>
                                                <p className="card-text"><strong>Movie Title</strong></p>
                                                <p className="card-text">JOKER</p>
                                                <p className="card-text"><strong>Seats</strong></p>
                                                <p className="card-text">
                                                    {reservation.Tickets.map((ticket) => (
                                                        <p>REGULAR SEAT - {ticket.seat}</p>
                                                    ))}
                                                </p>
                                                <p className="card-text"><strong>Time</strong></p>
                                                <p className="card-text">14:40</p>  
                                                <Link to={`/paid/${reservation.ID}`}><button className="btn btn-primary mt-2">Download Ticket</button></Link>
                                            </div>
                                        </div>
                                    </div>
                                    ))}
                                </div>
                            </div>

                            {/* History Tab */}
                            <div className="col-lg-12 col-md-12 col-sm-12 tab-pane fade" id="history">
                                <p>History content goes here.</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>  
        </>
    );
};

export default Tickets;