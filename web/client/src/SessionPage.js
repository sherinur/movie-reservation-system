// src/HomePage.js
import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import './style.scss';
import logo from './logo.png'; // Ensure the logo image is in the correct path

const SessionPage = () => {
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
        <div className="row justify-content-center">
          <div className="row">
            <div className="col-lg-8">
              <h3 className="custom-white">Theater</h3>
              <div className="btn-group mb-3" role="group">
                <button className="btn btn-outline-light">Bukit Bintang</button>
                <button className="btn btn-outline-light selected">IOI Putrajaye</button>
                <button className="btn btn-outline-light">KB Mall</button>
              </div>

              <h3 className="custom-white">Date</h3>
              <div className="d-flex flex-wrap gap-2 mb-3">
                <button className="btn btn-outline-light">22 Oct Mon</button>
                <button className="btn btn-outline-light selected">22 Oct Mon</button>
                <button className="btn btn-outline-light">22 Oct Mon</button>
                <button className="btn btn-outline-light">22 Oct Mon</button>
                <button className="btn btn-outline-light">22 Oct Mon</button>
              </div>

              <h3 className="custom-white">Time</h3>
              <div className="d-flex flex-wrap gap-2 mb-3">
                <button className="btn btn-outline-light">15:40</button>
                <button className="btn btn-outline-light selected">15:40</button>
                <button className="btn btn-outline-light">15:40</button>
                <button className="btn btn-outline-light">15:40</button>
              </div>
            </div>

            <div className="col-lg-3">
              <div className="card bg-dark text-white">
                <img src="movie.jpeg" className="card-img-top" alt="Movie Poster" />
                <div className="card-body">
                  <h5 className="card-title">SPIDERMAN ACROSS THE SPIDERVERSE</h5>
                  <p className="card-text">Movie description here...</p>
                  <p><strong>Duration:</strong> 2h 30m</p>
                  <p><strong>Type:</strong> Cartoon</p>
                </div>
              </div>
              <div className="card bg-dark text-white mt-3 p-3">
                <h5>IOI Putrajaye</h5>
                <p>28 October 2023</p>
                <p>15:40</p>
                <small>*Seat selection can be done after this</small>
                <button className="btn btn-success w-100 mt-2">Proceed</button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha3/dist/js/bootstrap.bundle.min.js"></script>
    </div>
  );
};

export default SessionPage;