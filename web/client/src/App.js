import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Register from './Register';
import Login from './Login';
import Tickets from './Tickets';
import HomePage from './HomePage';
import SessionPage from './SessionPage';
import SeatSelection from './SeatSelection';
import Paying from './Paying';
import Paid from './Paid';
import MovieAdminPanel from './MovieAdminPanel';
import CinemaAdminPanel from './CinemaAdminPanel';

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route path="/profile/tickets" element={<Tickets />} />
        <Route path="/booking/:id" element={<SeatSelection />} />
        <Route path="/order/:id" element={<Paying />} />
        <Route path="/paid/:id" element={<Paid />} />
        <Route path="/session/:movieID" element={<SessionPage />} />
        <Route path="/admin/movie" element={<MovieAdminPanel />} />
        <Route path="/admin/cinema" element={<CinemaAdminPanel />} />
      </Routes>
    </Router>
  );
};

export default App;