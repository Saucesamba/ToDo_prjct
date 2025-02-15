package app

import (
	"MyProject/internal/db"
	"MyProject/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

func filterTasks(tasks []models.Task, filter string) []models.Task {
	if filter == "" {
		return tasks
	}
	var f bool
	if filter == "true" {
		f = true
	} else if filter == "false" {
		f = false
	} else {
		fmt.Errorf("Filter ircorrect")
	}
	var res []models.Task
	for _, task := range tasks {
		if task.Completed == f {
			res = append(res, task)
		}
	}
	return res
}

func mergeSort(tasks []models.Task, sort string) []models.Task {
	if len(tasks) <= 1 {
		return tasks
	}
	mid := len(tasks) / 2
	left := tasks[:mid]
	right := tasks[mid:]
	return merge(mergeSort(left, sort), mergeSort(right, sort), sort)
}
func merge(left, right []models.Task, sort string) []models.Task {
	li := 0
	ri := 0
	var res []models.Task
	for li < len(left) && ri < len(right) {
		switch sort {
		case "asc":
			{
				if right[ri].CreatedAt.After(left[li].CreatedAt) {
					res = append(res, left[li])
					li++
				} else {
					res = append(res, right[ri])
					ri++
				}
			}

		case "dec":
			{
				if left[li].CreatedAt.After(right[ri].CreatedAt) {
					res = append(res, left[li])
					li++
				} else {
					res = append(res, right[ri])
					ri++
				}
			}
		}
	}
	for li < len(left) {
		res = append(res, left[li])
		li++
	}
	for ri < len(right) {
		res = append(res, right[ri])
		ri++
	}

	return res
}

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

func GetUserTasks(data *sql.DB, id int, sort, filter string) ([]models.Task, error) {
	tasks, err := db.GetAllTasks(data, id)
	if err != nil {
		fmt.Errorf("failed to fetch tasks: %v", err)
	}
	res := filterTasks(tasks, filter)
	if sort == "" {
		return res, nil
	}
	res2 := mergeSort(res, sort)
	return res2, nil
}

func CreateTask(data *sql.DB, task *models.Task, userId int) error {
	id, timeCr, err := db.CreateTask(data, task, userId)
	task.Id = id
	task.CreatedAt = timeCr
	if err != nil {
		fmt.Errorf("failed to create task: %v", err)
	}
	return nil
}
func DeleteTask(data *sql.DB, id int) error {
	err := db.DeleteTask(data, id)
	if err != nil {
		fmt.Errorf("failed to delete task: %v", err)
	}
	return nil
}
func UpdateTaskStatus(data *sql.DB, id int) error {
	task, err := db.GetTaskById(data, id)
	if err != nil {
		fmt.Errorf("failed to fetch task: %v", err)
	}
	if !task.Completed {
		task.Completed = true
	} else {
		task.Completed = false
	}
	err = db.UpdateTask(data, task)
	if err != nil {
		fmt.Errorf("failed to update task: %v", err)
	}
	return nil
}
func GetTaskStat(data *sql.DB, id int) (models.UserTaskInfo, error) {
	tasks, err := db.GetAllTasks(data, id)
	if err != nil {
		fmt.Errorf("failed to fetch tasks: %v", err)
	}
	var res models.UserTaskInfo
	for _, task := range tasks {
		if task.Completed == true {
			res.CompletedCount++
		}
	}
	res.TaskCount = len(tasks)
	return res, nil
}
