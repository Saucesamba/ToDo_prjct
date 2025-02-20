// src/components/TaskForm.js
import React, { useState } from 'react';
import axios from 'axios';
import { useContext } from 'react';
import { AuthContext } from '../context/AuthContext';


const TaskForm = ({onCreateSuccess}) => {
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [error, setError] = useState('');
    const { userId } = useContext(AuthContext);

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            //const requestBody = {
            //  name: name,
            //  description: description,
            //  userId: userId, // Ensure you're passing the userId
            //};
            const response = await axios.post(
                `http://localhost:8080/tasks?userId=${userId}`,
                {
                    name: name,
                    description: description,
                }

            );
            if(response.status === 200){
                onCreateSuccess();
                alert("task created successfully!")
            }

        } catch (e){
            if (e.response && e.response.data && e.response.data.error) {
                setError(e.response.data.error);
            } else {
                setError("An unexpected error occurred.");
            }

        }
    };

    return (
        <form onSubmit={handleSubmit}>
            {error && <div className="error-message">{error}</div>}
            <div>
                <label>Name:</label>
                <input
                    type="text"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    required
                />
            </div>
            <div>
                <label>Description:</label>
                <textarea
                    value={description}
                    onChange={(e) => setDescription(e.target.value)}
                    required
                />
            </div>
            <button type="submit">Create Task</button>
        </form>
    );
};

export default TaskForm;