package app

import (
	"MyProject/internal/db"
	"MyProject/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

// функция для регистрации пользователя, валидация на непустые поля
// добавить проверку на валидацию почты (что такой почты еще нет)
func RegistrUser(data *sql.DB, email, name, password string) (models.User, error) {
	if email == "" || password == "" || name == "" {
		return models.User{}, errors.New("All fields are required")
	}
	user := models.User{Email: email, Name: name, Password: password}
	id, err := db.CreateUser(data, &user)
	if err != nil {
		fmt.Errorf("failed to create user: %v", id)
	} else {
		fmt.Println("Registration was successful")
	}
	user.Id = id
	return user, nil
}

// функция для аутентификации пользователя
func AuthUser(data *sql.DB, email string, password string) (models.User, error) {
	foundUser, err := db.GetUserByEmail(data, email, password)
	if err != nil {
		fmt.Errorf("failed to fetch users: %v", err)
	}
	return foundUser, nil
}

// получение информации о пользователе по id
func GetInfoUser(data *sql.DB, id int) (models.User, error) {
	user, err := db.GetUserById(data, id)
	if err != nil {
		return models.User{}, fmt.Errorf("User with id %v not found", id)
	}
	return user, nil
}

// обновление информации о пользователе
func UpdateUser(data *sql.DB, user models.User) error {
	err := db.UpdateUser(data, &user)
	if err != nil {
		fmt.Errorf("failed to update user information: %v", err)
	}
	return nil
}

// удаление информации о пользователе
func DeleteUser(data *sql.DB, id int) error {
	err := db.DeleteUser(data, id)
	if err != nil {
		fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

// получение всех задач пользователя
func GetUserTasks(data *sql.DB, id int) ([]models.Task, error) {
	user := models.User{Id: id}
	tasks, err := db.GetAllTasks(data, &user)
	if err != nil {
		fmt.Errorf("failed to fetch tasks: %v", err)
	}
	return tasks, nil
}
