import React, { useEffect, useState } from 'react';
import logo from '../src/logo.png';
import 'bootstrap/dist/css/bootstrap.min.css';
import './style.scss';

const MoviePage = () => {
  const [movies, setMovies] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("http://localhost/movi/movie")
      .then(response => response.json())
      .then(data => {
        setMovies(data); 
        setLoading(false);
      })
      .catch(error => {
        console.error("Movie download error:", error);
        setLoading(false);
      });
  }, []);

  return (
    <div>
      <div className="background-wrapper"></div>

      <header className="text-white py-3">
        <div className="container d-flex justify-content-between align-items-center">
          <div className="d-flex align-items-center">
            <img src={logo} className="card-img-top" alt="logo" />
          </div>
          <div> 
            <a href="/login" className="btn btn-outline-light custom-green-btn">Log In</a>
            <a href="/register" className="btn btn-outline-light me-2">Register</a>
          </div>
        </div>
      </header>

      <div className="container text-center">
        <div className="row justify-content-center">
          <h2 className="mb-5 custom-white">Now Showing</h2>

          {loading ? (
            <p className="text-white">Loading...</p>
          ) : movies.length === 0 ? (
            <p className="text-white">There are no available movies.</p>
          ) : (
            movies.map((movie) => (
              <div className="col-md-2 movie-card" key={movie.id}>
                <a href="#" className="movie-link">
                  <div className="poster-container">
                    <img 
                      src={movie.posterimage || "https://via.placeholder.com/140x207"} 
                      className="card-img-top" 
                      alt={movie.title} 
                    />
                    <div className="rating-badge">
                      <span>&#9733;</span> {movie.rating || "N/A"}
                    </div>
                    <div className="pg-rating">{movie.pgrating || "N/A"}+</div>
                  </div>
                  <div className="card-body">
                    <h6 className="card-title">{movie.title}</h6>
                  </div>
                </a>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
};

export default MoviePage;
