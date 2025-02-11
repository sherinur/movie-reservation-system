import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Register from './Register';
import Login from './Login';
import SeatSelection from './SeatSelection';
import Paying from './Paying';
import Paid from './Paid';

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route path="/booking" element={<SeatSelection />} />
        <Route path="/paying" element={<Paying />} />
        <Route path="/paid" element={<Paid />} />
      </Routes>
    </Router>
  );
};

export default App;
