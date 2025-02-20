// src/pages/Login.js
import React from 'react';
import LoginForm from '../components/LoginForm';
import { useNavigate } from 'react-router-dom';

const Login = () => {
    const navigate = useNavigate();
    const handleLoginSuccess = () => {
        navigate('/profile/');
    };
    return (
        <div>
            <h2>Login</h2>

            <LoginForm onSuccess={handleLoginSuccess}/>

        </div>
    );
};

export default Login;