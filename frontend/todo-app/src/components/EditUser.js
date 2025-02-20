// src/components/EditUser.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';

const EditUser = ({ userId, onUpdateSuccess }) => {

    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/users/${userId}`);
                if(response.status === 200){
                    setName(response.data.name);
                    setEmail(response.data.email);
                }

            } catch (e){
                if(e.response.status === 404){
                    setError("User not found!")
                }else{
                    setError(e.response.data.error)
                }

            }finally {
                setLoading(false);
            }
        };
        fetchUser();
    }, [userId]);

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            const response = await axios.put(
                `http://localhost:8080/users/${userId}`,
                {
                    name: name,
                    email: email,
                    password: password,
                }
            );
            if(response.status === 200){
                onUpdateSuccess();
                alert("User update successful!")
            }
        } catch (e){
            if(e.response.status === 404){
                setError("User not found!")
            }else{
                setError(e.response.data.error)
            }
        }
    }

    if(loading){
        return <div>Loading form...</div>
    }
    if (error) {
        return <div style={{color: "red"}}>{error}</div>;
    }
    return(
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
                <label>Confirm password:</label>
                <input type="password" value={password} onChange={(e) => setPassword(e.target.value)}/>
            </div>
            <button type="submit">Update</button>
        </form>
    )
}
export default EditUser