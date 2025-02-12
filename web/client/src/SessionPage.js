// src/SessionPage.js
import React, { useEffect, useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "./style.scss";
import logo from "./logo.png";

const SessionPage = () => {
  // const [movie, setMovie] = useState(null);
  // const [sessions, setSessions] = useState([]);
  // const [loading, setLoading] = useState(true);

  // useEffect(() => {
  //   // Fetch movie data
  //   fetch("localhost/movi/movie/60d5ec49f9a1c72d4c8e4b8b")
  //     .then((response) => response.json())
  //     .then((data) => {
  //       setMovie(data);
  //       setLoading(false);
  //     })
  //     .catch((error) => {
  //       console.error("Error fetching movie:", error);
  //       setLoading(false);
  //     });

  //   // Fetch session data
  //   fetch("localhost/movi/sessions?movie_id=60d5ec49f9a1c72d4c8e4b8b")
  //     .then((response) => response.json())
  //     .then((data) => {
  //       setSessions(data);
  //       setLoading(false);
  //     })
  //     .catch((error) => {
  //       console.error("Error fetching sessions:", error);
  //       setLoading(false);
  //     });
  // }, []);

  return (
    <div>
      <div className="background-wrapper"></div>

      <header className="text-white py-3">
        <div className="container d-flex justify-content-between align-items-center">
          <div className="d-flex align-items-center">
            <img src={logo} className="card-img-top" alt="logo" />
          </div>
          <div>
            <a href="/logout" className="btn btn-outline-light custom-red-btn me-2">Logout</a>
          </div>
        </div>
      </header>

      <div className="container py-4">
        <div className="row">
          <div className="col-lg-4">
            {loading ? (
              <p className="text-white">Loading movie...</p>
            ) : movie ? (
              <div className="card bg-dark text-white">
                <img src={movie.PosterImage} className="card-img-top" alt={movie.Title} />
                <div className="card-body">
                  <h5 className="card-title">{movie.Title}</h5>
                  <p className="card-text">{movie.Description}</p>
                  <p><strong>Genre:</strong> {movie.Genre}</p>
                  <p><strong>Duration:</strong> {movie.Duration} minutes</p>
                  <p><strong>Language:</strong> {movie.Language}</p>
                  <p><strong>Release Date:</strong> {movie.ReleaseDate}</p>
                  <p><strong>Rating:</strong> {movie.Rating}</p>
                  <p><strong>PG Rating:</strong> {movie.PGrating}</p>
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
                  <div key={index} className="card bg-dark text-white mb-3 p-3">
                    <h5>{session.cinemaAddress} - Hall {session.hall_number}</h5>
                    <p><strong>Start:</strong> {new Date(session.start_time).toLocaleString()}</p>
                    <p><strong>End:</strong> {new Date(session.end_time).toLocaleString()}</p>
                    <button className="btn btn-success w-100 mt-2">Select</button>
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