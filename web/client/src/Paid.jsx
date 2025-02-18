import React, { useState, useEffect } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import { jsPDF } from "jspdf";


const Paid = () => {
    const navigate = useNavigate();
    const jwtToken = localStorage.getItem("accessToken")
    const id = useParams();
    const idString = id.id.toString();
    const [qr, setQr] = useState("")
    const [seats, setSeats] = useState([]);
    const [prices, setPrices] = useState([]);
    const [totalPrice, setTotalPrice] = useState(0);
    const [title, setTitle] = useState("");

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
                console.log(data)
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
                setTitle(data.MovieTitle)
                setQr(data.QRCode)
            })
            .catch((error) => console.error("Error getting reservation:", error));
        }, [idString, jwtToken]);

    const handleDownload = () => {
        const doc = new jsPDF();

        const base64QR = "data:image/png;base64," + qr; 
        
        // Add text to the PDF (x, y, text)
        doc.text("THANK YOU FOR YOUR PURCHASE!", 10, 10);
        doc.text(`Your tickets on the movie ${title}:`, 10, 20);
        doc.text(seats.join(", ") || "None", 10,30);

        doc.text("Show this QR code to guard before session", 10, 50);
        doc.addImage(base64QR, "PNG", 10, 60, 50, 50)
    
        // Save the file
        doc.save("ticket" + idString + ".pdf");
    };


    return (
        <>
            <div className="background-wrapper"></div>
            
            <div className="container paid-container">
                <div className="reservation-info">
                    <h1 className="title">Ticket Detail</h1>

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

                    <button onClick={handleDownload} className="checkout-btn">Download Ticket</button>
                    <Link to="/"><button className="back-btn">Back to Movies</button></Link>
                </div>
                
            </div>
        </>
    );
};

export default Paid;