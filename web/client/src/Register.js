import React from "react";
import { Button, Form, Container } from 'react-bootstrap';
import { Link } from 'react-router-dom';

const Register = () => {
    return (
        <Container fluid className="p-0">
            <div className="row no-gutters hero">
                <div className="col-lg-6 col-md-6 col-sm-12 hero-poster">
                    <div className="overlay">
                        <h1>Welcome.</h1>
                        <p>Begin your cinematic adventure now with our ticketing platform!</p>
                    </div>
                </div>
                
                <div className="col-lg-4 col-md-6 col-sm-12 hero-form mx-auto">
                    <div className="container">
                        <h2 className="mb-4 title">Create an account</h2>

                        {/* Email Input */}
                        <Form.Group className="mb-3">
                            <Form.Label htmlFor="email">Email</Form.Label>
                            <Form.Control type="email" id="email" placeholder="Enter your email" required />
                            <div id="email-error" className="error-message mt-2">Email error</div>
                        </Form.Group>

                        {/* Username Input */}
                        <Form.Group className="mb-3">
                            <Form.Label htmlFor="username">Username</Form.Label>
                            <Form.Control type="text" id="username" placeholder="Enter your username" required />
                            <div id="username-error" className="error-message mt-2">Fullname error</div>
                        </Form.Group>

                        {/* Password Input */}
                        <Form.Group className="mb-3">
                            <Form.Label htmlFor="password">Password</Form.Label>
                            <Form.Control type="password" id="password" placeholder="Create a password" required />
                            <div id="password-error" className="error-message mt-2">Password error</div>
                        </Form.Group>

                        {/* Confirm Password Input */}
                        <Form.Group className="mb-3">
                            <Form.Control type="password" id="confirm-password" placeholder="Confirm your password" required />
                            <div id="confirm-password-error" className="error-message mt-2">Confirm Password error</div>
                        </Form.Group>

                        {/* Sign Up Button */}
                        <div className="d-grid">
                            <Button type="submit" className="btn-accept" id="submit-btn">Create account</Button>
                        </div>

                        {/* Already have an account link */}
                        <div className="mt-3 text-center">
                            <span className="already-account">
                                Already have an account ?{' '}
                                <Link to="/login" className="text-decoration-none link">Sign In</Link>
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </Container>
    );
};

export default Register;
