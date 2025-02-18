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
    
            navigate("/admin/movie");
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

                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
};

export default Login;
