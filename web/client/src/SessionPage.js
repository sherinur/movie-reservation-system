import React, { useEffect, useState } from "react";
import { useParams, Link, useNavigate } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import "./style.scss";
import logo from "./logo.png";

const SessionPage = () => {
  const { movieID } = useParams();
  const [movie, setMovie] = useState(null);
  const [sessions, setSessions] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();
  const jwtToken = localStorage.getItem("accessToken");

  const options = { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric', 
    hour: '2-digit', 
    // minute: '2-digit', 
    // second: '2-digit', 
    // timeZoneName: 'short'
  };

  useEffect(() => {
    // Fetch movie data
    fetch(`http://localhost/movi/movie/${movieID}`)
      .then((response) => response.json())
      .then((data) => {
        setMovie(data);
        setLoading(false);
      })
      .catch((error) => {
        console.error("Error fetching movie:", error);
        setLoading(false);
      });

    // Fetch session data
    fetch(`http://localhost/movi/session/movie/${movieID}`)
      .then((response) => response.json())
      .then((data) => {
        setSessions(data || []); // Ensure sessions is an array
        setLoading(false);
      })
      .catch((error) => {
        console.error("Error fetching sessions:", error);
        setLoading(false);
      });
  }, [movieID]);

  const handleLogout = () => {
    localStorage.removeItem("accessToken");
    navigate("/login");
  };

  const handleBooking = (sessionID) => {
    navigate(`/booking/${sessionID}`)
  }

  const convertDuration = (minutes) => {
    const hours = Math.floor(minutes / 60);
    const mins = minutes % 60;
    return `${hours}h ${mins}m`;
  };

  return (
    <div>
      <div className="background-wrapper"></div>

      <header className="text-white py-3">
        <div className="container d-flex justify-content-between align-items-center">
          <Link to="/" className="d-flex justify-content-center">
            <img src={logo} className="card-img-top" alt="logo"/>
          </Link>
          <div>
            {jwtToken ? (
              <>
                <a href="/profile/tickets" className="btn btn-outline-light custom-green-btn">My Tickets</a>
                <button className="btn btn-outline-light me-2" onClick={handleLogout}>Logout</button>
              </>
            ) : (
              <>
                <a href="/login" className="btn btn-outline-light custom-green-btn">Log In</a>
                <a href="/register" className="btn btn-outline-light me-2">Register</a>
              </>
            )}
          </div>
        </div>
      </header>

      <div className="container py-4">
        <div className="row">
          <div className="col-lg-3">
            {loading ? (
              <p className="text-white">Loading movie...</p>
            ) : movie ? (
              <div className="text-white movie-card">
                <img src={movie.posterimage} className="card-img-top" alt={movie.title} />
                <div className="card-body">
                  <h5 className="card-title">{movie.title}</h5>
                  <p className="card-text">{movie.description}</p>
                  <p><strong>Genre:</strong> {movie.genre}</p>
                  <p><strong>Duration:</strong> {convertDuration(movie.duration)}</p>
                  <p><strong>Rating:</strong> {movie.rating}</p>
                  <p><strong>PG Rating:</strong> {movie.pgrating}</p>
                </div>
              </div>
            ) : (
              <p className="text-white">Movie not found.</p>
            )}
          </div>

          <div className="col-lg-8">
            <h3 className="custom-white">Available Sessions</h3>
            <div className="row">
              {loading ? (
                <p className="text-white">Loading sessions...</p>
              ) : sessions.length === 0 ? (
                <p className="text-white">There are no available sessions.</p>
              ) : (
                sessions.map((session, index) => (
                  <div key={index} className="col-md-6 col-lg-4 mb-4">
                    <div className="card session-card bg-dark text-white">
                      <div className="card-body">
                        <h5 className="card-title">{session.cinemaAddress}</h5>
                        <p className="card-text">Hall {session.hall_number}</p>
                        <p className="card-text"><strong>Date:</strong> {new Date(session.start_time).toLocaleString('en-US', options)}</p>
                        <button className="btn btn-success w-100 mt-2" onClick={() => handleBooking(session._id)}>Select</button>
                      </div>
                    </div>
                  </div>
                ))
              )}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SessionPage;