// src/components/RegisterForm.js
import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { useContext } from 'react';
import { AuthContext } from '../context/AuthContext';

const RegisterForm = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();
    const { login } = useContext(AuthContext);
    const handleSubmit = async (e) => {
        e.preventDefault();
        console.log("name", name, "email", email, "password", password)
        try {

            const response = await axios.post(
                `http://localhost:8080/users/register`,
                {
                    name: name,
                    email: email,
                    password: password,
                }
            );
            if(response.status === 201){
                login(response.data.id)
                navigate(`/users/${response.data.id}`)

            }

        } catch (e){
            setError(e.response.data.error)
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            {error && <div style={{color: "red"}}>{error}</div>}
            <div>
                <label>Name:</label>
                <input type="text" value={name} onChange={(e) => setName(e.target.value)} required />
            </div>
            <div>
                <label>Email:</label>
                <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
            </div>
            <div>
                <label>Password:</label>
                <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
            </div>
            <button type="submit">Register</button>
        </form>
    );
};

export default RegisterForm;