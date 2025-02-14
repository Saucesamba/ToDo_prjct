package db

import (
	"MyProject/internal/models"
	"database/sql"
	"fmt"
	"time"
)

func CreateTask(db *sql.DB, task *models.Task, userId int) (int, time.Time, error) {
	query := "INSERT INTO tasks (name, description,completed,user_id) VALUES ($1, $2, $3, $4) returning id,created_at;"
	var id int
	var createdTime time.Time
	err := db.QueryRow(query, task.Name, task.Description, task.Completed, userId).Scan(&id, &createdTime)
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("failed to create task: %w", err)
	}
	return id, createdTime, nil
}
func GetAllTasks(db *sql.DB, id int) ([]models.Task, error) {
	query := "SELECT * FROM tasks WHERE user_id = $1;"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query all tasks: %w", err)
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Completed, &task.UserId, &task.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}
	return tasks, nil
}
func GetTaskById(db *sql.DB, id int) (*models.Task, error) {
	query := "SELECT * FROM tasks WHERE id = $1;"
	row := db.QueryRow(query, id)
	task := &models.Task{}
	err := row.Scan(&task.Id, &task.Name, &task.Description, &task.Completed, &task.UserId, &task.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan task with Id: %w", err)
	}
	return task, nil
}
func UpdateTask(db *sql.DB, task *models.Task) error {
	query := "UPDATE tasks SET completed = $1 WHERE id = $2"
	_, err := db.Exec(query, task.Completed, task.Id)
	if err != nil {
		return fmt.Errorf("error updating task: %w", err)
	}
	return nil
}
func DeleteTask(db *sql.DB, id int) error {
	query := "DELETE FROM tasks WHERE id = $1;"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}
	return nil
}
