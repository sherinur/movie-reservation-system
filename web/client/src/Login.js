import React from 'react';

const Login = () => {
    return (
        <>
            <div className="background-wrapper"></div>

            <div className="container">
                <div className="row justify-content-center login-section">
                    <div className="col-lg-6 col-md-6 col-sm-8">
                        <div className="card py-5 px-5 login-card">
                            <h2 className="mb-4 title">Login to your account</h2>
                            <form id="loginForm">
                                {/* Error Message */}
                                <div id="error-message" className="error-message mb-3">
                                    Invalid email or password. Please try again.
                                </div>

                                {/* Email Input */}
                                <div className="mb-3">
                                    <label htmlFor="email" className="form-label">Email</label>
                                    <input type="email" className="form-control" id="email" placeholder="Enter your email" />
                                </div>

                                {/* Password Input */}
                                <div className="mb-3">
                                    <label htmlFor="password" className="form-label">Password</label>
                                    <input type="password" className="form-control" id="password" placeholder="Enter your password" />
                                </div>

                                {/* Forgot Password Link */}
                                <div className="mb-3 text-end">
                                    <a href="/forgot" className="forgot-password link text-decoration-none">Forgot Password?</a>
                                </div>

                                {/* Sign In Button */}
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
