import React, { useEffect, useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import logo from '../src/logo.png';
import 'bootstrap/dist/css/bootstrap.min.css';
import './style.scss';

const HomePage = () => {
  const [movies, setMovies] = useState([]);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    fetch("http://localhost/movi/movie", {
      method: 'GET'
    })
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

  const handleMovieClick = (movieID) => {
    navigate(`/session/${movieID}`);
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
              <div className="col-md-2 movie-card" key={movie.id} onClick={() => handleMovieClick(movie.id)}>
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
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
};

export default HomePage;