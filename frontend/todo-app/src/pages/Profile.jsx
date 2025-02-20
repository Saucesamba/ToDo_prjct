// src/pages/UserPage.js
import React, { useState } from 'react';
import Profile from '../components/Profile';
import EditUser from '../components/EditUser';
import { useParams } from 'react-router-dom';

const UserPage = () => {
    const { userId } = useParams();
    const [isEditing, setIsEditing] = useState(false);

    const handleUpdateSuccess = () => {
        setIsEditing(false); // Закрываем форму редактирования после успешного обновления
    };

    const handleEditClick = () => {
        setIsEditing(true); // Открываем форму редактирования
    };

    return (
        <div>
            <Profile userId={userId} />
            {!isEditing && ( // Отображаем кнопку "Edit" только если isEditing === false
                <button onClick={handleEditClick}>Edit Profile</button>
            )}
            {isEditing && ( // Отображаем EditUser только если isEditing === true
                <EditUser userId={userId} onUpdateSuccess={handleUpdateSuccess} />
            )}
        </div>
    );
};

export default UserPage;