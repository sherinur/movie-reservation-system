import React, { useState, useEffect } from 'react';
import { Modal, Button, Form, Alert } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import './style.scss';
import { Link, useNavigate } from 'react-router-dom';

const CinemaAdminPanel = () => {
  const [cinemas, setCinemas] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [showHallModal, setShowHallModal] = useState(false);
  const [editMode, setEditMode] = useState(false);
  const [editHallMode, setEditHallMode] = useState(false);
  const jwtToken = localStorage.getItem("accessToken");
  const navigate = useNavigate();
  const [currentCinema, setCurrentCinema] = useState({
    id: '',
    name: '',
    city: '',
    address: '',
    rating: '',
    hall_list: []
  });
  const [currentHall, setCurrentHall] = useState({
    id: '',
    number: '',
    rows: 0,
    columns: 0,
    available_seats: '',
    seats: []
  });
  const [expandedCinemaId, setExpandedCinemaId] = useState(null);
  const [expandedHallNumber, setExpandedHallNumber] = useState(null);
  const [showSuccessMessage, setShowSuccessMessage] = useState(false);

  useEffect(() => {
    if (!jwtToken) {
      navigate('/admin/login');
    } else {
      fetchCinemas();
    }
  }, [jwtToken, navigate]);

  const fetchCinemas = async () => {
    try {
      const response = await fetch("http://localhost/movi/cinema");
      const data = await response.json();
      setCinemas(data || []);
    } catch (error) {
      console.error("Error fetching cinemas:", error);
    }
  };

  const handleShowModal = (cinema = {
    id: '',
    name: '',
    city: '',
    address: '',
    rating: '',
    hall_list: []
  }) => {
    setCurrentCinema(cinema);
    setEditMode(!!cinema.id);
    setShowModal(true);
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setCurrentCinema({
      id: '',
      name: '',
      city: '',
      address: '',
      rating: '',
      hall_list: []
    });
  };

  const handleShowHallModal = (cinema, hall = {
    id: '',
    number: '',
    rows: 0,
    columns: 0,
    available_seats: '',
    seats: []
  }) => {
    setCurrentCinema(cinema);
    setCurrentHall(hall);
    setEditHallMode(!!hall.number);
    setShowHallModal(true);
  };

  const handleCloseHallModal = () => {
    setShowHallModal(false);
    setCurrentHall({
      id: '',
      number: '',
      rows: 0,
      columns: 0,
      available_seats: '',
      seats: []
    });
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCurrentCinema({ ...currentCinema, [name]: value });
  };

  const handleHallChange = (e) => {
    const { name, value } = e.target;
    const updatedValue = name === "number" || name === "available_seats" || name === "id" || name === "rows" || name === "columns" ? parseInt(value, 10) : value;
    const updatedHall = { ...currentHall, [name]: updatedValue };

    setCurrentHall(updatedHall);
  };

  const handleCardClick = (cinemaId) => {
    setExpandedCinemaId(expandedCinemaId === cinemaId ? null : cinemaId);
  };

  const handleHallClick = (hallNumber) => {
    setExpandedHallNumber(expandedHallNumber === hallNumber ? null : hallNumber);
  };

  const generateSeats = (rows, columns) => {
    const seats = [];
    const rowLetters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
    for (let i = 0; i < rows; i++) {
      for (let j = 0; j < columns; j++) {
        seats.push({
          row: rowLetters[i],
          column: (j + 1).toString(),
          status: 'available'
        });
      }
    }
    return seats;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const url = editMode ? `http://localhost/movi/cinema/${currentCinema.id}` : 'http://localhost/movi/cinema';
    const method = editMode ? 'PUT' : 'POST';

    // Convert rating to float
    const cinemaToSubmit = { ...currentCinema, rating: parseFloat(currentCinema.rating) };

    try {
      const response = await fetch(url, {
        method,
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(cinemaToSubmit)
      });

      if (response.ok) {
        fetchCinemas();
        handleCloseModal();
        setShowSuccessMessage(true);
        setTimeout(() => setShowSuccessMessage(false), 3000); // Hide success message after 3 seconds
      } else {
        console.error("Error saving cinema:", response.statusText);
      }
    } catch (error) {    
      console.error("Error saving cinema:", error);
    }
  };

  const handleHallSubmit = async (e) => {
    e.preventDefault();

    // Set currentHall.id to currentHall.number for POST requests
    if (!editHallMode) {
      setCurrentHall((prevHall) => ({ ...prevHall, id: prevHall.number }));
    }

    const method = editHallMode ? 'PUT' : 'POST';
    const url = editHallMode
      ? `http://localhost/movi/cinema/${currentCinema.id}/hall/${currentHall.id}`
      : `http://localhost/movi/cinema/${currentCinema.id}/hall`;

    // Generate seats for the hall
    const seats = generateSeats(currentHall.rows, currentHall.columns);

    const updatedHall = { ...currentHall, seats, available_seats: seats.length, cinema_id: currentCinema.id };
    const updatedHallList = editHallMode
      ? currentCinema.hall_list.map(hall => hall.id === currentHall.id ? updatedHall : hall)
      : [...currentCinema.hall_list, updatedHall];

    const updatedCinema = { ...currentCinema, hall_list: updatedHallList };

    try {
      const response = await fetch(url, {
        method,
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(updatedHall)
      });

      if (response.ok) {
        fetchCinemas();
        handleCloseHallModal();
        setShowSuccessMessage(true);
        setTimeout(() => setShowSuccessMessage(false), 3000); // Hide success message after 3 seconds
      } else {
        console.error(`Error ${editHallMode ? 'updating' : 'saving'} hall:`, response.statusText);
      }
    } catch (error) {
      console.error(`Error ${editHallMode ? 'updating' : 'saving'} hall:`, error);
    }
  };

  const handleDelete = async (id) => {
    try {
      const response = await fetch(`http://localhost/movi/cinema/${currentCinema.id}/hall/${id}`, {
        method: 'DELETE'
      });

      if (response.ok) {
        fetchCinemas();
      } else {
        console.error("Error deleting cinema:", response.statusText);
      }
    } catch (error) {
      console.error("Error deleting cinema:", error);
    }
  };

  const handleHallDelete = async (hallNumber, cinemaID) => {
    const updatedHallList = currentCinema.hall_list.filter(hall => hall.number !== hallNumber);
    const updatedCinema = { ...currentCinema, hall_list: updatedHallList };

    try {
      const response = await fetch(`http://localhost/movi/cinema/${cinemaID}/hall/${hallNumber}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(updatedCinema)
      });

      if (response.ok) {
        fetchCinemas();
      } else {
        console.error("Error deleting hall:", response.statusText);
      }
    } catch (error) {
      console.error("Error deleting hall:", error);
    }
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
                <Link to="/admin/cinema" className="nav-link active text-dark fw-bold">Cinemas</Link>
              </li>
              <li className="nav-item">
                <Link to="/admin/session" className="nav-link text-dark">Sessions</Link>
              </li>
              <li>
              <button className="btn me-2 custom-green-btn" onClick={handleLogout}>Logout</button>
              </li>
            </ul>
          </div>
        </nav>

        {/* Main Content */}
        <main className="col-md-10 ms-sm-auto px-md-4">
          <div className="d-flex justify-content-between align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 className="h4">Cinemas</h1>
            <button className="btn btn-success" onClick={() => handleShowModal()}>+ Create new</button>
          </div>
          {showSuccessMessage && (
            <Alert variant="success">
              Cinema successfully updated!
            </Alert>
          )}
          <table className="table table-bordered">
            <thead>
              <tr>
                <th>#</th>
                <th>Name</th>
                <th>City</th>
                <th>Address</th>
                <th>Rating</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              {cinemas && cinemas.map((cinema, index) => (
                <React.Fragment key={cinema.id}>
                  <tr onClick={() => handleCardClick(cinema.id)}>
                    <td>{index + 1}</td>
                    <td>{cinema.name}</td>
                    <td>{cinema.city}</td>
                    <td>{cinema.address}</td>
                    <td>{cinema.rating}</td>
                    <td>
                      <button className="btn btn-sm btn-warning" onClick={(e) => { e.stopPropagation(); handleShowModal(cinema); }}>✏️</button>
                      <button className="btn btn-sm btn-danger" onClick={(e) => { e.stopPropagation(); handleDelete(cinema.id); }}>🗑️</button>
                    </td>
                  </tr>
                  {expandedCinemaId === cinema.id && (
                    <tr>
                      <td colSpan="6">
                        <div className="expanded-info">
                          <p><strong>Halls:</strong></p>
                          {(cinema.hall_list || []).map((hall, hallIndex) => (
                            <div key={hallIndex} onClick={() => handleHallClick(hall.number)}>
                              <p><strong>Hall Number:</strong> {hall.number}</p>
                              <p><strong>Seats:</strong> {hall.available_seats}</p>
                              {expandedHallNumber === hall.number && (
                                <div>
                                  <div className="seat-grid">
                                    {(hall.seats || []).map((seat, seatIndex) => (
                                      <div key={seatIndex} className={`seat ${seat.status}`}>
                                        {seat.row}{seat.column}
                                      </div>
                                    ))}
                                  </div>
                                  <button className="btn btn-sm btn-warning" onClick={(e) => { e.stopPropagation(); handleShowHallModal(cinema, hall); }}>✏️</button>
                                  <button className="btn btn-sm btn-danger" onClick={(e) => { e.stopPropagation(); handleHallDelete(hall.id, cinema.id); }}>🗑️</button>
                                </div>
                              )}
                            </div>
                          ))}
                          <button className="btn btn-success mt-2" onClick={(e) => { e.stopPropagation(); handleShowHallModal(cinema); }}>+ Add Hall</button>
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

      {/* Modal for Add/Edit Cinema */}
      <Modal show={showModal} onHide={handleCloseModal}>
        <Modal.Header closeButton>
          <Modal.Title>{editMode ? 'Edit Cinema' : 'Add Cinema'}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleSubmit}>
            <Form.Group controlId="formName">
              <Form.Label>Name</Form.Label>
              <Form.Control
                type="text"
                name="name"
                value={currentCinema.name}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formCity" className="mt-3">
              <Form.Label>City</Form.Label>
              <Form.Control
                type="text"
                name="city"
                value={currentCinema.city}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formAddress" className="mt-3">
              <Form.Label>Address</Form.Label>
              <Form.Control
                type="text"
                name="address"
                value={currentCinema.address}
                onChange={handleChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formRating" className="mt-3">
              <Form.Label>Rating</Form.Label>
              <Form.Control
                type="number"
                name="rating"
                value={currentCinema.rating}
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

      {/* Modal for Add/Edit Hall */}
      <Modal show={showHallModal} onHide={handleCloseHallModal}>
        <Modal.Header closeButton>
          <Modal.Title>{editHallMode ? 'Edit Hall' : 'Add Hall'}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form onSubmit={handleHallSubmit}>
            <Form.Group controlId="formHallId">
              <Form.Label>ID</Form.Label>
              <Form.Control
                type="number"
                name="id"
                value={currentHall.id}
                readOnly
              />
            </Form.Group>
            <Form.Group controlId="formHallNumber">
              <Form.Label>Hall Number</Form.Label>
              <Form.Control
                type="number"
                name="number"
                value={currentHall.number}
                onChange={handleHallChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formRows" className="mt-3">
              <Form.Label>Rows</Form.Label>
              <Form.Control
                type="number"
                name="rows"
                value={currentHall.rows}
                onChange={handleHallChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formColumns" className="mt-3">
              <Form.Label>Columns</Form.Label>
              <Form.Control
                type="number"
                name="columns"
                value={currentHall.columns}
                onChange={handleHallChange}
                required
              />
            </Form.Group>
            <Form.Group controlId="formAvailableSeats" className="mt-3">
              <Form.Label>Available Seats</Form.Label>
              <Form.Control
                type="number"
                name="available_seats"
                value={currentHall.available_seats}
                readOnly
              />
            </Form.Group>
            <Button variant="primary" type="submit" className="mt-3">
              {editHallMode ? 'Update' : 'Add'}
            </Button>
          </Form>
        </Modal.Body>
      </Modal>
    </div>
  );
};

export default CinemaAdminPanel;