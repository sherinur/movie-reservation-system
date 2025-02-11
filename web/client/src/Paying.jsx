import React, { useState } from "react";

const Paying = () => {
  return (
        <>
            <div className="background-wrapper"></div>
            
            <div class="container booking-container">
                <div className="order-info">
                    <h1 class="title">Booking Detail</h1>

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

                    <div class="transaction-detail">
                        <h2>Transaction Detail</h2>
                        <div class="price-row">
                            <p>REGULAR SEAT</p>
                            <p>1500 Tg x3</p>
                        </div>
                        <div class="price-row">
                            <p>Service Charge (6%)</p>
                            <p>90 Tg x3</p>
                        </div>
                        <hr></hr>
                        <div class="total-payment">
                            <p><strong>Total payment</strong></p>
                            <p class="total-amount">4770 Tg</p>
                        </div>
                        <p class="note">*Purchased ticket cannot be canceled</p>
                    </div>

                    <a href="/paid"><button class="checkout-btn">Checkout Ticket</button></a>
                </div>
            </div>   
        </>
    );
};

export default Paying;