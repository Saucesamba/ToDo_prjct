package db

import (
	"MyProject/internal/models"
	"database/sql"
	"fmt"
)

func CreateUser(db *sql.DB, user *models.User) (int, error) {
	query := "INSERT INTO users (name, password, email) values ($1, $2, $3) RETURNING id"
	var id int
	err := db.QueryRow(query, user.Name, user.Password, user.Email).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Error creating user: %v", err)
	}
	return id, nil
}

func GetUserById(db *sql.DB, id int) (models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE id=$1"
	row := db.QueryRow(query, id)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("Error getting user: %v", err)
	}
	return user, nil
}
func GetUserByEmail(db *sql.DB, email, password string) (models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE email=$1 AND password=$2"
	row := db.QueryRow(query, email, password)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("Error getting user: %v", err)
	}
	return user, nil
}
func GetAllUsers(db *sql.DB) ([]models.User, error) {
	var users []models.User
	query := "select * from users"
	rows, err := db.Query(query)
	if err != nil {
		return users, fmt.Errorf("Error getting users: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("Error getting users: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}
	return users, nil
}

func UpdateUser(db *sql.DB, user *models.User) error {
	query := "UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4"
	_, err := db.Exec(query, user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		return fmt.Errorf("Error updating user: %v", err)
	}
	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Error deleting user: %v", err)
	}
	return nil
}
