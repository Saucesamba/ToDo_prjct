import React, { useState, useEffect, useContext } from 'react';
import { BrowserRouter, Route, Routes, Link } from 'react-router-dom';
import Home from './pages/Home';
import Auth from './pages/Auth';
import Registre from './pages/Registre';
import Profile from './pages/Profile';
import Footer from "./components/Footer";
import './index.css';
import TaskPage from "./components/TaskPage";
import { AuthProvider, AuthContext } from './context/AuthContext';

function App() {

    return (
        <BrowserRouter>
            <AuthProvider>
                <AppContent/>
            </AuthProvider>
        </BrowserRouter>
    );
}

function AppContent(){
    const { userId, logout } = useContext(AuthContext);

    return(
        <div className="app-container">
            <header className="app-header">
                <h1><Link to="/" style={{textDecoration: "none", color: "white"}}>To-Do List</Link></h1>
                <div className="auth-buttons">
                    {userId ? (
                        <button onClick={logout}>Sign Out</button>
                    ) : (
                        /*Здесь происходит навигация на страницы авторизации*/
                        <>
                            <Link to="users/login"><button>Sign In</button></Link>
                            <Link to="users/register"><button>Register</button></Link>
                        </>
                    )}
                </div>
            </header>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="users/login" element={<Auth />} />
                <Route path="users/register" element={<Registre />} />
                <Route path="users/:userId" element={<Profile />} />
                <Route path="tasks" element={<TaskPage />} />
            </Routes>
            <Footer />
        </div>
    )

}


export default App;