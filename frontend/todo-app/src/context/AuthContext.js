import React, { createContext, useState, useEffect } from 'react';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [userId, setUserId] = useState(() => {
        // Получаем userId из localStorage при инициализации
        const storedUserId = localStorage.getItem('userId');
        return storedUserId ? parseInt(storedUserId, 10) : null;
    });

    const login = (id) => {
        localStorage.setItem('userId', id); // Сохраняем userId в localStorage
        setUserId(id);
    };

    const logout = () => {
        localStorage.removeItem('userId'); // Удаляем userId из localStorage
        setUserId(null);
    };

    useEffect(() => {
        // Обновляем localStorage при изменении userId
        if (userId) {
            localStorage.setItem('userId', userId);
        } else {
            localStorage.removeItem('userId');
        }
    }, [userId]);

    const value = {
        userId,
        login,
        logout,
    };

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
};