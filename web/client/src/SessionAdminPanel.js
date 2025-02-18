import React, { useState, useEffect } from 'react';
import { Modal, Button, Form, Alert } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import './style.scss';

const SessionAdminPanel = () => {
  const [sessions, setSessions] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [editMode, setEditMode] = useState(false);
  const jwtToken = localStorage.getItem("accessToken");
  const navigate = useNavigate();
  const [currentSession, setCurrentSession] = useState({
    id: '',
    movie_id: '',
    cinema_id: '',
    address: '',
    hall_number: '',
    start_time: '',
    end_time: '',
    available_seats: 0,
    seats: []
  });
  const [movies, setMovies] = useState([]);
  const [cinemas, setCinemas] = useState([]);
  const [expandedSessionId, setExpandedSessionId] = useState(null);
  const [showSuccessMessage, setShowSuccessMessage] = useState(false);

  useEffect(() => {
    if (!jwtToken) {
      navigate('/admin/login');
    } else {
      fetchSessions();
      fetchMovies();
      fetchCinemas();
    }
  }, [jwtToken, navigate]);

  const fetchSessions = async () => {
    try {
      const response = await fetch("http://localhost/movi/session", {
        method: 'GET',
      });
      const data = await response.json();
      setSessions(data);
    } catch (error) {
      console.error("Error fetching sessions:", error);
    }
  };

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

  const fetchCinemas = async () => {
    try {
      const response = await fetch("http://localhost/movi/cinema");
      const data = await response.json();
      setCinemas(data);
    } catch (error) {
      console.error("Error fetching cinemas:", error);
    }
  };

  const handleShowModal = (session = {
    id: '',
    movie_id: '',
    cinema_id: '',
    address: '',
    hall_number: '',
    start_time: '',
    end_time: '',
    seats: [],
    available_seats: 0
  }) => {
    setCurrentSession(session);
    setEditMode(!!session.id);
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setCurrentSession({
      id: '',
      movie_id: '',
      cinema_id: '',
      address: '',
      hall_number: '',
      start_time: '',
      end_time: '',
      seats: [],
      available_seats: 0
    });
  };

  const options = {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit',
    timeZoneName: 'short'
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCurrentSession({ ...currentSession, [name]: name === 'hall_number' ? parseInt(value, 10) : value });
  };

  const handleCardClick = (sessionId) => {
    setExpandedSessionId(expandedSessionId === sessionId ? null : sessionId);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const url = editMode ? `http://localhost/movi/session/${currentSession.id}` : 'http://localhost/movi/session';
    const method = editMode ? 'PUT' : 'POST';

    // Format start_time and end_time to include timezone information
    const formattedSession = {
      ...currentSession,
      start_time: new Date(currentSession.start_time).toISOString(),
      end_time: new Date(currentSession.end_time).toISOString()
    };

    try {
      const response = await fetch(url, {
        method,
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(formattedSession),
      });

      if (response.ok) {
        fetchSessions();
        handleCloseModal();
        setShowSuccessMessage(true);
        setTimeout(() => setShowSuccessMessage(false), 3000); // Hide success message after 3 seconds
      } else {
        console.error("Error saving session:", response.statusText);
      }
    } catch (error) {
      console.error("Error saving session:", error);
    }
  };

  const handleDelete = async (id) => {
    try {
      const response = await fetch(`http://localhost/movi/session/${id}`, {
        method: 'DELETE'
      });

      if (response.ok) {
        fetchSessions();
      } else {
        console.error("Error deleting session:", response.statusText);
      }
    } catch (error) {
      console.error("Error deleting session:", error);
    }
  };

  const getMovieTitle = (movieId) => {
    const movie = movies.find(movie => movie.id === movieId);
    return movie ? movie.title : 'Unknown';
  };

  const getCinemaName = (cinemaId) => {
    const cinema = cinemas.find(cinema => cinema.id === cinemaId);
    return cinema ? cinema.name : 'Unknown';
  };

  const handleLogout = () => {
    localStorage.removeItem("accessToken");
    navigate("/admin/login");
  };

  return (
    <div className="container-fluid">
      <div className="row">
        {/* Sidebar */}
        <nav className="col-md-2 d-md-block bg-light sidebar vh-100">
          <div className="position-sticky">
            <ul className="nav flex-column">
              <li className="nav-item">
                <Link to="/admin/movie" className="nav-link text-dark">Movies</Link>
              </li>
              <li className="nav-item">
                <Link to="/admin/cinema" className="nav-link text-dark">Cinemas</Link>
              </li>
              <li className="nav-item">
                <Link to="/admin/session" className="nav-link active text-dark fw-bold">Sessions</Link>
              </li>
              <li>
              <button className="btn  me-2 custom-green-btn" onClick={handleLogout}>Logout</button>
              </li>
            </ul>
          </div>
        </nav>

        {/* Main Content */}
        <main className="col-md-10 ms-sm-auto px-md-4">
          <div className="d-flex justify-content-between align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 className="h4">Sessions</h1>
            <button className="btn btn-success" onClick={() => handleShowModal()}>+ Create new</button>
          </div>
          {showSuccessMessage && (
            <Alert variant="success">
              Session successfully updated!
            </Alert>
          )}
          <table className="table table-bordered">
            <thead>
              <tr>
                <th>#</th>
                <th>Movie</th>
                <th>Cinema</th>
                <th>Hall Number</th>
                <th>Start Time</th>
                <th>End Time</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              {sessions.map((session, index) => (
                <React.Fragment key={session.id}>
                  <tr onClick={() => handleCardClick(session.id)}>
                    <td>{index + 1}</td>
                    <td>{getMovieTitle(session.movie_id)}</td>
                    <td>{getCinemaName(session.cinema_id)}</td>
                    <td>{session.hall_number}</td>
                    <td>{new Date(session.start_time).toLocaleString('en-US', options)}</td>
                    <td>{new Date(session.end_time).toLocaleString('en-US', options)}</td>
                    <td>
                      <button className="btn btn-sm btn-warning" onClick={() => handleShowModal(session)}>✏️</button>
                      <button className="btn btn-sm btn-danger" onClick={() => handleDelete(session.id)}>🗑️</button>
                    </td>
                  </tr>
                  {expandedSessionId === session.id && (
                    <tr>
                      <td colSpan="7">
                        <div className="expanded-info">
                          <p><strong>Seats:</strong> {session.seats.length}</p>
                          <p><strong>Available Seats:</strong> {session.available_seats}</p>
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

      {/* Modal for Add/Edit Session */}
      <Modal show={showModal} onHide={handleCloseModal}>
        <Modal.Header closeButton>
          <Modal.Title>{editMode ? 'Edit Session' : 'Add Session'}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleSubmit}>
            <Form.Group controlId="formMovieID">
              <Form.Label>Movie</Form.Label>
              <Form.Control
                as="select"
                name="movie_id"
                value={currentSession.movie_id}
                onChange={handleChange}
                required
              >
                <option value="">Select a movie</option>
                {movies.map(movie => (
                  <option key={movie.id} value={movie.id}>{movie.title}</option>
                ))}
              </Form.Control>
            </Form.Group>
            <Form.Group controlId="formCinemaID" className="mt-3">
              <Form.Label>Cinema</Form.Label>
              <Form.Control
                as="select"
                name="cinema_id"
                value={currentSession.cinema_id}
                onChange={handleChange}
                required
              >
                <option value="">Select a cinema</option>
                {cinemas.map(cinema => (
                  <option key={cinema.id} value={cinema.id}>{cinema.name}</option>
                ))}
              </Form.Control>
            </Form.Group>
            <Form.Group controlId="formHallNumber" className="mt-3">
              <Form.Label>Hall Number</Form.Label>
              <Form.Control
                type="number"
                name="hall_number"
                value={currentSession.hall_number}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formStartTime" className="mt-3">
              <Form.Label>Start Time</Form.Label>
              <Form.Control
                type="datetime-local"
                name="start_time"
                value={currentSession.start_time}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formEndTime" className="mt-3">
              <Form.Label>End Time</Form.Label>
              <Form.Control
                type="datetime-local"
                name="end_time"
                value={currentSession.end_time}
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

export default SessionAdminPanel;