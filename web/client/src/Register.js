import React, { useState } from "react";
import { Button, Form, Container, Alert } from "react-bootstrap";
import { Link, useNavigate } from "react-router-dom";

const Register = () => {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [error, setError] = useState(null);
    const [success, setSuccess] = useState(null);
    const navigate = useNavigate();

    const handleRegister = async (e) => {
        e.preventDefault();
        setError(null);
        setSuccess(null);

        if (password !== confirmPassword) {
            setError("Passwords do not match");
            return;
        }

        try {
            const response = await fetch("http://localhost/auth/users/register", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ email, username, password }),
            });

            const data = await response.json();
            
            if (!response.ok) {
                throw new Error(data.message || "Registration failed");
            }

            setSuccess("User registered successfully");
            setTimeout(() => navigate("/login"), 2000);
        } catch (err) {
            setError(err.message);
        }
    };

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
                        {error && <Alert variant="danger">{error}</Alert>}
                        {success && <Alert variant="success">{success}</Alert>}

                        <Form onSubmit={handleRegister}>
                            <Form.Group className="mb-3">
                                <Form.Label>Email</Form.Label>
                                <Form.Control 
                                    type="email" 
                                    value={email} 
                                    onChange={(e) => setEmail(e.target.value)}
                                    placeholder="Enter your email" 
                                    required 
                                />
                            </Form.Group>

                            <Form.Group className="mb-3">
                                <Form.Label>Username</Form.Label>
                                <Form.Control 
                                    type="text" 
                                    value={username} 
                                    onChange={(e) => setUsername(e.target.value)}
                                    placeholder="Enter your username" 
                                    required 
                                />
                            </Form.Group>

                            <Form.Group className="mb-3">
                                <Form.Label>Password</Form.Label>
                                <Form.Control 
                                    type="password" 
                                    value={password} 
                                    onChange={(e) => setPassword(e.target.value)}
                                    placeholder="Create a password" 
                                    required 
                                />
                            </Form.Group>

                            <Form.Group className="mb-3">
                                <Form.Label>Confirm Password</Form.Label>
                                <Form.Control 
                                    type="password" 
                                    value={confirmPassword} 
                                    onChange={(e) => setConfirmPassword(e.target.value)}
                                    placeholder="Confirm your password" 
                                    required 
                                />
                            </Form.Group>

                            <div className="d-grid">
                                <Button type="submit" className="btn-accept">Create account</Button>
                            </div>
                        </Form>

                        <div className="mt-3 text-center">
                            <span className="already-account">
                                Already have an account? <Link to="/login" className="text-decoration-none link">Sign In</Link>
                            </span>
                        </div>
                    </div>
                </div>
            </div>
        </Container>
    );
};

export default Register;
