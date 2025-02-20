import React, { useState, useEffect, useContext } from 'react';
import axios from 'axios';
import { AuthContext } from '../context/AuthContext';

const TaskPage = () => {
    const [tasks, setTasks] = useState(null);
    const [error, setError] = useState('');
    const [loading, setLoading] = useState(true);
    const { userId } = useContext(AuthContext);
    const [name, setName] = useState('');
    const [description, setDescription] = useState('');
    const [showForm, setShowForm] = useState(false);

    console.log(userId)
    const toggleTaskForm = () => {
        setShowForm((prevShowForm) => !prevShowForm);
    };


    useEffect(() => {
        const fetchTasks = async () => {
            try {
                const response = await axios.get(`http://localhost:8080/tasks?userId=${userId}`);
                console.log("responseTask", response)
                if (response.status === 200) {
                    if (response.data.tasks === null) {
                        setTasks(null); // Set tasks to null if server returns {"tasks": null}
                    } else {
                        setTasks(response.data);
                    }
                }
            } catch (e) {
                setError("Failed to fetch tasks.");
            } finally {
                setLoading(false);
            }
        };

        if (userId) {
            fetchTasks();
        } else {
            setLoading(false);
            setTasks([])
        }
    }, [userId]);

    const handleCreateSuccess = () => {
        setShowForm(false);
        // Optionally refresh the task list after successful creation
        // You could call fetchTasks here if needed.
    };
    const handleSubmit = async (e) => {
        e.preventDefault();

        try {

            const response = await axios.post(
                `http://localhost:8080/tasks?userId=${userId}`,
                {
                    name: name,
                    description: description,
                }
            );
            if(response.status === 200){
                setTasks((prev) => (Array.isArray(prev) ? prev.concat(response.data) : [response.data]))
                setShowForm(false)
            }

        } catch (e){

            setError("An unexpected error occurred.");

        }
    };
    if (loading) {
        return <div>Loading tasks...</div>;
    }

    if (error) {
        return <div className="error-message">{error}</div>;
    }

    return (
        <div>
            <h3>Tasks:</h3>

            <div>
                <button onClick={toggleTaskForm}>
                    {showForm ? 'Hide Task Form' : 'Show Task Form'}
                </button>
                {showForm &&
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
                    </form>}
            </div>

            <ul>
                {tasks && tasks.length > 0 ? (
                    tasks.map((task) => (
                        <li key={task.id}>
                            {task.name} - {task.description}
                        </li>
                    ))
                ) : (
                    <p>No tasks found.</p>
                )}
            </ul>
        </div>
    );
};

export default TaskPage;