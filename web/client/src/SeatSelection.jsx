import React, { useState } from "react";

const SeatSelection = () => {
  const rows = ["A", "B", "C", "D", "E", "F", "G"];
  const seatsPerRow = 10;
  const [selectedSeats, setSelectedSeats] = useState([]);

  const toggleSeat = (seat) => {
    setSelectedSeats((prev) =>
      prev.includes(seat)
        ? prev.filter((s) => s !== seat)
        : [...prev, seat]
    );
  };

  return (
        <>
            <div class="background-wrapper"></div>

            <div className="seat-selection-container">
                <h1 class="title">Seat Selection</h1>

                {/* Screen */}
                <div className="screen">SCREEN</div>

                <div className="hall">
                    {rows.map((row) => (
                        <div key={row} className="seat-map"> 
                        {Array(seatsPerRow)
                            .fill(0)
                            .map((_, i) => {
                            const seatNumber = `${row}${i + 1}`;
                            return (
                                <div
                                key={seatNumber}
                                className={`seat ${selectedSeats.includes(seatNumber) ? "selected" : ""}`}
                                onClick={() => toggleSeat(seatNumber)}
                                >
                                {seatNumber}
                                </div>
                            );
                            })}
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
                        <strong>C4, C5, C6</strong>
                    </div>
                    <div className="action-buttons">
                        <a href="booking.html"><button className="btn-cancel">Back</button></a>
                        <a href="/paying"><button className="btn-proceed">Proceed Payment</button></a>
                    </div>
                </div>
            </div>    
        </>
    );
};

export default SeatSelection;