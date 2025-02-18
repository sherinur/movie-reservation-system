import React, { useState, useEffect } from "react";
import { useNavigate, useParams, Link } from "react-router-dom";


const SeatSelection = () => {
    const navigate = useNavigate();
    const sessionId = useParams().id.toString();
    const [movieId, setMovieId] = useState("");
    const jwtToken = localStorage.getItem("accessToken");

    const [seats, setSeats] = useState([]);
    const [bookedSeats, setBookedSeats] = useState([]);
    const [selectedSeats, setSelectedSeats] = useState([]);
    const [groupedSeats, setGroupedSeats] = useState({});

    const price = 1500;
    const [totalPrice, setTotalPrice] = useState(0); 
    const [title, setTitle] = useState("");

    // Fetch seat data from server
    useEffect(() => {
        fetch("http://localhost/movi/session/" + sessionId, {
            method: "GET",
            headers: { 
                "Content-Type": "application/json",
            },
        })
            .then((response) => response.json())
            .then((data) => {
                let seat = []

                for (let ticket of data.seats) {
                    seat.push(ticket)
                }

                setSeats(seat)
                setMovieId(data.movie_id)
            })
            .catch((error) => console.error("Error getting reservation:", error));
        }, [sessionId]);

    useEffect(() => {
        // Group seats by their letter and number them
        const groups = {}
        seats.forEach((seat) => {
        if (!groups[seat.row]) {
            groups[seat.row] = [];
        }
        groups[seat.row].push(seat);
        });
        
        setGroupedSeats(groups)
    }, [seats]);

    useEffect(() => {
        fetch("http://localhost/movi/movie/" + movieId, {
            method: "GET",
            headers: { 
                "Content-Type": "application/json",
            },
        })
            .then((response) => response.json())
            .then((data) => {
                setTitle(data.title)
            })
            .catch((error) => console.error("Error getting reservation:", error));
    }, [movieId]);
    

    // Toggle seat selection
    const toggleSeat = (seat) => {
        setSelectedSeats((prev) => {
            if (prev.includes(seat)) {
                setTotalPrice((new_price) => totalPrice - price); // Subtract price if deselected
                return prev.filter((s) => s !== seat);
            } else {
                setTotalPrice((new_price) => totalPrice + price); // Add price if selected
                return [...prev, seat];
            }
        });
    };

    // Handle payment (send selected seats to backend)
    const handleStartPayment = () => {
        if (selectedSeats.length === 0) {
        alert("Please, choose seats to reserve")
        return
        }
        let jsonData = {
            screening_id: sessionId, 
            movie_title: title,
            tickets: []
        }

        for (let ticket of selectedSeats) {
        jsonData.tickets.push({
            seat: ticket,
            price: price,
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
        .then((data) => navigate("/order/" + data.InsertedID))
        .catch((error) => console.error("Error confirming seats:", error));
    };

  return (
        <>
            <div className="background-wrapper"></div>

            <div className="seat-selection-container">
                <h1 className="title">{title}</h1>

                {/* Screen */}
                <div className="screen">SCREEN</div>

                <div className="hall">
                    {Object.entries(groupedSeats).map(([row, seatNumbers]) => (
                        <div key={row} className="seat-map">
                            {seatNumbers.map((seat) => (
                                <div onClick={() => toggleSeat(seat.row + seat.column)} key={seat.row + seat.column} className={`seat ${selectedSeats.includes(seat.row + seat.column) ? "selected" : ""}`}>
                                    {seat.row + seat.column}
                                </div>
                            ))}
                        </div>
                    ))}
                </div>
                
                <div className="seat-selection-footer">
                    <div className="total-price">
                        <p>Total:</p>
                        <strong>{totalPrice} Tg</strong>
                    </div>
                    <div className="selected-seats">
                      <p>Seats: </p>
                      <strong>{selectedSeats.join(", ") || "None"}</strong>
                    </div>
                    <div className="action-buttons">
                        <Link to="/"><button className="btn-cancel">Back</button></Link>
                        <button onClick={handleStartPayment} className="btn-proceed">Proceed Payment</button>
                    </div>
                </div>
            </div>    
        </>
    );
};

export default SeatSelection;