// src/components/Footer.js
import React from 'react';

const Footer = () => {
    return (
        <footer className="app-footer">
            <div className="footer-content">
                <div className="footer-section">
                    <h3>Контакты</h3>
                    <p>Email: arhipartyom@yandex.ru</p>
                    <p>Phone: +7-925-322-94-77</p>
                </div>
                <div className="footer-section">
                    <h3>Социальные сети</h3>
                    <a href="https://vk.com">ВКонтакте</a>
                </div>
                <div className="footer-section">
                    <h3>Информация</h3>
                    <p>Описание компании</p>
                </div>
            </div>
            <div className="footer-bottom">
                <p>{new Date().getFullYear()} To-Do list from SauceSamba</p>
            </div>
        </footer>
    );
};

export default Footer;