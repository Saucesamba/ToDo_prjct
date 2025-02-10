package app

import (
	"MyProject/internal/db"
	"MyProject/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

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
func AuthUser(data *sql.DB, email string, password string) (models.User, error) {
	foundUser, err := db.GetUserByEmail(data, email, password)
	if err != nil {
		fmt.Errorf("failed to fetch users: %v", err)
	}
	if foundUser.Email == "" || foundUser.Password == "" {
		return models.User{}, fmt.Errorf("User with email %s not found", email)
	}
	return foundUser, nil
}
func GetInfoUser(data *sql.DB, id int) (models.User, error) {
	user, err := db.GetUserById(data, id)
	if err != nil {
		return models.User{}, fmt.Errorf("User with id %v not found", id)
	}
	return user, nil
}
func UpdateUser(data *sql.DB, user models.User) error {
	findPassw, err := db.GetUserById(data, user.Id)
	if err != nil {
		fmt.Errorf("failed to fetch user: %v", err)
	}
	if user.Password == "" {
		return fmt.Errorf("Please, confirm your password")
	}
	if findPassw.Password != user.Password {
		return errors.New("Passwords aren't equal")
	}
	err = db.UpdateUser(data, &user)
	if err != nil {
		fmt.Errorf("failed to update user information: %v", err)
	}
	return nil
}
func DeleteUser(data *sql.DB, id int) error {
	err := db.DeleteUser(data, id)
	if err != nil {
		fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

// получение всех задач пользователя
func GetUserTasks(data *sql.DB, id int) ([]models.Task, error) {
	tasks, err := db.GetAllTasks(data, id)
	if err != nil {
		fmt.Errorf("failed to fetch tasks: %v", err)
	}
	return tasks, nil
}

func CreateTask(data *sql.DB, task models.Task, userId int) error {
	id, err := db.CreateTask(data, &task, userId)
	task.Id = id
	if err != nil {
		fmt.Errorf("failed to create task: %v", err)
	}
	return nil
}
