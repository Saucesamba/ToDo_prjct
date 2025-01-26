// Слой работы с БД
//Реализация операций CRUD для Task

package db

//db.Query(): выполняет SQL-запрос, который возвращает набор строк (sql.Rows).
//rows.Next(): переходит к следующей строке результатов.
//rows.Scan(): читает значения из текущей строки и записывает их в переменные.
//db.QueryRow(): выполняет запрос, который возвращает одну строку.
//db.Exec(): выполняет SQL-запрос, который не возвращает строк (например, INSERT, UPDATE, DELETE).

import (
	"MyProject/internal/config"
	"MyProject/internal/models"

	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*sql.DB, error) {

	Source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	db, err := sql.Open("postgres", Source)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db: %w", err)
	}
	return db, nil
}

func CreateTask(db *sql.DB, task *models.Task) (int, error) {
	query := "INSERT INTO tasks (name, description,completed) VALUES ($1, $2, $3) returning id"
	var id int
	err := db.QueryRow(query, task.Name, task.Description, task.Completed).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %w", err)
	}
	return id, nil
}

func GetAllTasks(db *sql.DB) ([]models.Task, error) {
	rows, err := db.Query("SELECT id, name, description, completed FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("failed to query all tasks: %w", err)
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Completed); err != nil {
			return nil, fmt.Errorf("failed to scan task: %w", err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return tasks, nil
}

// получение таски по Id
func GetTaskById(db *sql.DB, id int) (*models.Task, error) {
	query := "SELECT id, name, description, completed FROM tasks WHERE id = $1"
	row := db.QueryRow(query, id)
	task := &models.Task{}

	err := row.Scan(&task.Id, &task.Name, &task.Description, &task.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan task with Id: %w", err)
	}
	return task, nil
}

// обновление в БД
func UpdateTask(db *sql.DB, task *models.Task) error {
	query := "UPDATE tasks SET name = $1, description = $2, completed = $3 WHERE id = $4"
	_, err := db.Exec(query, task.Name, task.Description, task.Completed, task.Id)
	if err != nil {
		return fmt.Errorf("error updating task: %w", err)
	}
	return nil
}

// Удаление из БД
func DeleteTask(db *sql.DB, id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}
	return nil
}
