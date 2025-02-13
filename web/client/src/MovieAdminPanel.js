import React, { useState, useEffect } from 'react';
import { Modal, Button, Form, Alert } from 'react-bootstrap';
import {Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import './style.scss';

const AdminPanel = () => {
  const [movies, setMovies] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [editMode, setEditMode] = useState(false);
  const [currentMovie, setCurrentMovie] = useState({
    id: '',
    title: '',
    genre: '',
    description: '',
    posterimage: '',
    duration: 0,
    language: '',
    releasedate: '',
    rating: '',
    pgrating: '',
    production: '',
    producer: '',
    status: ''
  });
  const [expandedMovieId, setExpandedMovieId] = useState(null);
  const [showSuccessMessage, setShowSuccessMessage] = useState(false);

  useEffect(() => {
    fetchMovies();
  }, []);

  const fetchMovies = async () => {
    try {
      const response = await fetch("http://localhost/movi/movie", {
        method: 'GET',
      });
      const data = await response.json();
      setMovies(data);
    } catch (error) {
      console.error("Error fetching movies:", error);
    }
  };

  const handleShowModal = (movie = {
    id: '',
    title: '',
    genre: '',
    description: '',
    posterimage: '',
    duration: 0,
    language: '',
    releasedate: '',
    rating: '',
    pgrating: '',
    production: '',
    producer: '',
    status: ''
  }) => {
    setCurrentMovie(movie);
    setEditMode(!!movie.id);
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setCurrentMovie({
      id: '',
      title: '',
      genre: '',
      description: '',
      posterimage: '',
      duration: 0,
      language: '',
      releasedate: '',
      rating: '',
      pgrating: '',
      production: '',
      producer: '',
      status: ''
    });
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCurrentMovie({ ...currentMovie, [name]: value });
  };

  const handleCardClick = (movieId) => {
    setExpandedMovieId(expandedMovieId === movieId ? null : movieId);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const url = editMode ? `http://localhost/movi/movie/${currentMovie.id}` : 'http://localhost/movi/movie';
    const metho = editMode ? 'PUT' : 'POST';
  
    try {
      const movieToSubmit = { ...currentMovie, duration: parseInt(currentMovie.duration, 10) };
      const response = await fetch(url, {
        method: metho,
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(movieToSubmit),
      });
  
      if (response.ok) {
        fetchMovies();
        handleCloseModal();
        setShowSuccessMessage(true);
        setTimeout(() => setShowSuccessMessage(false), 3000); // Hide success message after 3 seconds
      } else {
        console.error("Error saving movie:", response.statusText);
      }
    } catch (error) {    
      console.error("Error saving movie:", error);
    }
  };
  

  const handleDelete = async (id) => {
    try {
      const response = await fetch(`http://localhost/movi/movie/${id}`, {
        method: 'DELETE'
      });

      if (response.ok) {
        fetchMovies();
      } else {
        console.error("Error deleting movie:", response.statusText);
      }
    } catch (error) {
      console.error("Error deleting movie:", error);
    }
  };

  return (
    <div className="container-fluid">
      <div className="row">
        {/* Sidebar */}
        <nav className="col-md-2 d-md-block bg-light sidebar vh-100">
          <div className="position-sticky">
            <ul className="nav flex-column">
              <li className="nav-item">
                <Link to="/admin/movie" className="nav-link active text-dark fw-bold">Movies</Link>
              </li>
              <li className="nav-item">
              <Link to="/admin/cinema" className="nav-link text-dark">Cinema</Link>
              </li>
              <li className="nav-item">
                <Link to="/admin/session" className="nav-link text-dark">Sessions</Link>
              </li>
              <li className="nav-item">
                <a className="nav-link text-dark" href="#">Users</a>
              </li>
              <li className="nav-item">
                <a className="nav-link text-dark" href="#">Orders</a>
              </li>
              <li className="nav-item">
                <a className="nav-link text-dark" href="#">Report</a>
              </li>
            </ul>
          </div>
        </nav>

        {/* Main Content */}
        <main className="col-md-10 ms-sm-auto px-md-4">
          <div className="d-flex justify-content-between align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 className="h4">Movies</h1>
            <button className="btn btn-success" onClick={() => handleShowModal()}>+ Create new</button>
          </div>
          {showSuccessMessage && (
            <Alert variant="success">
              Movie successfully updated!
            </Alert>
          )}
          <table className="table table-bordered">
            <thead>
              <tr>
                <th>#</th>
                <th>Cover</th>
                <th>Title</th>
                <th>Genre</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              {movies.map((movie, index) => (
                <React.Fragment key={movie.id}>
                  <tr onClick={() => handleCardClick(movie.id)}>
                    <td>{index + 1}</td>
                    <td><img src={movie.posterimage || "https://via.placeholder.com/50"} alt="Movie Cover" className="img-fluid" style={{ width: '50px' }} /></td>
                    <td>{movie.title}</td>
                    <td>{movie.genre}</td>
                    <td>
                      <button className="btn btn-sm btn-warning" onClick={() => handleShowModal(movie)}>✏️</button>
                      <button className="btn btn-sm btn-danger" onClick={() => handleDelete(movie.id)}>🗑️</button>
                    </td>
                  </tr>
                  {expandedMovieId === movie.id && (
                    <tr>
                      <td colSpan="5">
                        <div className="expanded-info">
                          <p><strong>Description:</strong> {movie.description}</p>
                          <p><strong>Duration:</strong> {movie.duration} minutes</p>
                          <p><strong>Language:</strong> {movie.language}</p>
                          <p><strong>Release Date:</strong> {movie.releasedate}</p>
                          <p><strong>Rating:</strong> {movie.rating}</p>
                          <p><strong>PG Rating:</strong> {movie.pgrating}</p>
                          <p><strong>Production:</strong> {movie.production}</p>
                          <p><strong>Producer:</strong> {movie.producer}</p>
                          <p><strong>Status:</strong> {movie.status}</p>
                        </div>
                      </td>
                    </tr>
                  )}
                </React.Fragment>
              ))}
            </tbody>
          </table>
        </main>
      </div>

      {/* Modal for Add/Edit Movie */}
      <Modal show={showModal} onHide={handleCloseModal}>
        <Modal.Header closeButton>
          <Modal.Title>{editMode ? 'Edit Movie' : 'Add Movie'}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleSubmit}>
            <Form.Group controlId="formTitle">
              <Form.Label>Title</Form.Label>
              <Form.Control
                type="text"
                name="title"
                value={currentMovie.title}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formGenre" className="mt-3">
              <Form.Label>Genre</Form.Label>
              <Form.Control
                type="text"
                name="genre"
                value={currentMovie.genre}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formDescription" className="mt-3">
              <Form.Label>Description</Form.Label>
              <Form.Control
                as="textarea"
                rows={3}
                name="description"
                value={currentMovie.description}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formPosterImage" className="mt-3">
              <Form.Label>Poster Image URL</Form.Label>
              <Form.Control
                type="text"
                name="posterimage"
                value={currentMovie.posterimage}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formDuration" className="mt-3">
              <Form.Label>Duration</Form.Label>
              <Form.Control
                type="number"
                name="duration"
                value={currentMovie.duration}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formLanguage" className="mt-3">
              <Form.Label>Language</Form.Label>
              <Form.Control
                type="text"
                name="language"
                value={currentMovie.language}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formReleaseDate" className="mt-3">
              <Form.Label>Release Date</Form.Label>
              <Form.Control
                type="text"
                name="releasedate"
                value={currentMovie.releasedate}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formRating" className="mt-3">
              <Form.Label>Rating</Form.Label>
              <Form.Control
                type="text"
                name="rating"
                value={currentMovie.rating}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formPGRating" className="mt-3">
              <Form.Label>PG Rating</Form.Label>
              <Form.Control
                type="text"
                name="pgrating"
                value={currentMovie.pgrating}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formProduction" className="mt-3">
              <Form.Label>Production</Form.Label>
              <Form.Control
                type="text"
                name="production"
                value={currentMovie.production}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formProducer" className="mt-3">
              <Form.Label>Producer</Form.Label>
              <Form.Control
                type="text"
                name="producer"
                value={currentMovie.producer}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formStatus" className="mt-3">
              <Form.Label>Status</Form.Label>
              <Form.Control
                type="text"
                name="status"
                value={currentMovie.status}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Button variant="primary" type="submit" className="mt-3">
              {editMode ? 'Update' : 'Add'}
            </Button>
          </Form>
        </Modal.Body>
      </Modal>
    </div>
  );
};

export default AdminPanel;