import React, { useState } from "react";
import { useNavigate, Link} from "react-router-dom";
import { Alert } from "react-bootstrap";
import axios from "axios";
import logo from '../src/logo.png';

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState(null);
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        setError(null);

        try {
            const response = await axios.post("http://localhost/auth/users/login", {
                email,
                password
            });
    
            localStorage.setItem("accessToken", response.data.accessToken);
    
            navigate("/");
        } catch (err) {
            if (err.response) {
                setError(err.response.data.message || "Login failed");
            } else {
                setError("Server error, please try again later.");
            }
        }
    };

    return (
        <>
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

            <div className="container">
                <div className="row justify-content-center login-section">
                    <div className="col-lg-6 col-md-6 col-sm-8">
                        <div className="card py-5 px-5 login-card">
                            <h2 className="mb-4 title">Login to your account</h2>
                            {error && <Alert variant="danger">{error}</Alert>}
                            <form onSubmit={handleLogin}>
                                <div className="mb-3">
                                    <label htmlFor="email" className="form-label">Email</label>
                                    <input 
                                        type="email" 
                                        className={`form-control ${error ? "error-input" : ""}`} 
                                        id="email" 
                                        placeholder="Enter your email" 
                                        value={email} 
                                        onChange={(e) => setEmail(e.target.value)} 
                                    />
                                </div>

                                <div className="mb-3">
                                    <label htmlFor="password" className="form-label">Password</label>
                                    <input 
                                        type="password" 
                                        className={`form-control ${error ? "error-input" : ""}`} 
                                        id="password" 
                                        placeholder="Enter your password" 
                                        value={password} 
                                        onChange={(e) => setPassword(e.target.value)} 
                                    />
                                </div>

                                <div className="mb-3 text-end">
                                    <a href="/forgot" className="forgot-password link text-decoration-none">Forgot Password?</a>
                                </div>

                                <div className="d-grid">
                                    <button type="submit" className="btn-accept">Login now</button>
                                </div>

                                <div className="mt-3 text-center">
                                    <span className="already-account">
                                        Don't have an account? <a href="/register" className="link text-decoration-none">Register here</a>
                                    </span>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
};

export default Login;
