// src/components/LoginForm.js
import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { useContext } from 'react';
import { AuthContext } from '../context/AuthContext';

const LoginForm = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();
    const { login } = useContext(AuthContext);
    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.put(
                `http://localhost:8080/users/login`,
                {
                    email: email,
                    password: password,
                }
            );
            console.log("response", response)

            if (response.status === 200) {
                login(response.data.id)
                navigate(`/users/${response.data.id}`)
            }
            return response
        } catch (e){
            console.log("e", e)
            setError(e.response.data.error)
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            {error && <div style={{color: "red"}}>{error}</div>}
            <div>
                <label>Email:</label>
                <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
            </div>
            <div>
                <label>Password:</label>
                <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
            </div>
            <button type="submit">Login</button>
        </form>
    );
};

export default LoginForm;