import React, { useState } from "react";

const Paid = () => {
  return (
        <>
            <div class="background-wrapper"></div>
            
            <div class="container paid-container">
                <div className="reservation-info">
                    <h1 class="title">Ticket Detail</h1>

                    <div class="schedule-section">
                        <h2>Schedule</h2>
                        <p><strong>Movie Title</strong></p>
                        <p class="movie-title">SPIDERMAN NO WAY HOME</p>
                        
                        <p><strong>Date</strong></p>
                        <p class="movie-date">Mon, 23 Oct 2023</p>

                        <div class="ticket-info">
                            <div class="seats">
                                <strong>Ticket (3)</strong>
                                <p>C8, C9, C10</p>
                            </div>
                            <div class="time">
                                <strong>Hours</strong>
                                <p>14:40</p>
                            </div>
                        </div>
                    </div>

                    <button class="checkout-btn">Download Ticket</button>
                    <a href="/booking"><button class="back-btn">Back to Movies</button></a>
                </div>
                
            </div>
        </>
    );
};

export default Paid;