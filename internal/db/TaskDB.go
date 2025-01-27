// CRUD операции для задач
package db

import (
	"MyProject/internal/models"
	"database/sql"
	"fmt"
)

// Функция для создания задачи пользователю. На вход ID пользователя, а также параметры задачи.
// На выходе id созданной задачи
func CreateTask(db *sql.DB, task *models.Task, user *models.User) (int, error) {
	query := "INSERT INTO tasks (name, description,completed,userid) VALUES ($1, $2, $3, $4) returning id"
	var id int
	err := db.QueryRow(query, task.Name, task.Description, task.Completed, user.Id).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %w", err)
	}
	return id, nil
}

// Функция для получения всех задач пользователя
// Ищем по связанному ID пользователя
// На выходе массив всех задач пользователя
func GetAllTasks(db *sql.DB, user *models.User) ([]models.Task, error) {

	rows, err := db.Query("SELECT id, name, description, completed FROM tasks WHERE userid = $1;", user.Id)
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
// На выходе задача пользователя или ошибка
func GetTaskById(db *sql.DB, id int, user *models.User) (*models.Task, error) {
	query := "SELECT id, name, description, completed FROM tasks WHERE id = $1; userid = $2;"
	row := db.QueryRow(query, id, user.Id)
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
// проверка на то что пользователь может изменять только свою задачу
func UpdateTask(db *sql.DB, task *models.Task, user *models.User) error {
	if task.UserId != user.Id {
		return fmt.Errorf("invalid user")
	} else {
		query := "UPDATE tasks SET name = $1, description = $2, completed = $3 WHERE id = $4"

		_, err := db.Exec(query, task.Name, task.Description, task.Completed, task.Id)
		if err != nil {
			return fmt.Errorf("error updating task: %w", err)
		}
	}
	return nil
}

// Удаление из БД
// Также есть проверка на то, что пользователь не может удалять не свою задачу
func DeleteTask(db *sql.DB, id int, user *models.User) error {
	query := "DELETE FROM tasks WHERE id = $1; userid = $2;"
	_, err := db.Exec(query, id, user.Id)
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}
	return nil
}
