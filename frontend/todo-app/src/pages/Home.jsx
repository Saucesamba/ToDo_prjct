// src/pages/Home.js
import React from 'react';
import image1 from '../photos/101.jpg';
import image2 from '../photos/102.jpg';
import { Link } from 'react-router-dom'

const Home = ({ isLoggedIn }) => {
    return (
        <div>
            <div className="home-content">
                        <section className="home-section">
                            <img
                                src = {image2}
                                alt="Планируйте свои задачи"
                                className="home-image"
                            />
                            <div className="home-text">
                                <h3>Планируйте свои задачи</h3>
                                <p>Мгновенно фиксируйте и упорядочивайте задачи, записывая их своими словами.</p>
                            </div>
                        </section>

                        <section className="home-section">
                            <img
                                src={image1}
                                alt="Достигайте своих целей"
                                className="home-image"
                            />
                            <div className="home-text">
                                <h3>Достигайте своих целей</h3>
                                <p>Наглядный список задач позволит Вам в разы быстрее достигать своих целей.</p>
                                <Link to="/tasks">
                                    <button>Посмотреть задачи</button>
                                </Link>
                            </div>
                        </section>


            </div>

        </div>
    );
};

export default Home;