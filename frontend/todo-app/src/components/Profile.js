// src/components/Profile.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Profile = ({ userId }) => {
    const [user, setUser] = useState(null);
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/users/${userId}`);
                console.log("response", response)
                if(response.status === 200){
                    setUser(response.data)
                }

            } catch (e){
                if(e.response.status === 404){
                    setError("User not found!")
                }else{
                    setError(e.response.data.error)
                }

            } finally {
                setLoading(false);
            }
        };
        fetchUser();
    }, [userId]);

    if (loading) {
        return <div>Loading profile...</div>;
    }

    if (error) {
        return <div style={{color: "red"}}>{error}</div>;
    }

    if (!user) {
        return <div>No user information available.</div>;
    }

    return (
        <div>
            <h2>User Profile</h2>
            <p><strong>ID:</strong> {user.id}</p>
            <p><strong>Name:</strong> {user.name}</p>
            <p><strong>Email:</strong> {user.email}</p>
        </div>
    );
};

export default Profile;