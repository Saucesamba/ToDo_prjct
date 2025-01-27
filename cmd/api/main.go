package main

import (
	"MyProject/internal/config"
	"MyProject/internal/db"
	"MyProject/internal/models"
	"fmt"
	"log"
)

func main() {
	//Load config
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	dbConn, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Printf("failed to close db connection: %v", err)
		}
	}()
	firstUser := &models.User{
		Name:     "First User",
		Password: "123456",
		Email:    "first@email.com",
	}
	firstUserId, err := db.CreateUser(dbConn, firstUser)

	if err != nil {
		fmt.Printf("Failed to create first user: %v", err)
	} else {
		fmt.Printf("Created first user: %v", firstUserId)
	}

	newTask := &models.Task{
		Name:        "Починить стул",
		Description: "Очень важное дело",
		Completed:   false}

	taskId, err := db.CreateTask(dbConn, newTask, firstUser)
	if err != nil {
		log.Fatalf("failed to create task: %v", err)
	} else {
		log.Printf("task created: %v", taskId)
	}

	tasks, err := db.GetAllTasks(dbConn, firstUser)
	if err != nil {
		log.Fatalf("failed to get all tasks: %v", err)
	} else {
		for _, task := range tasks {
			log.Printf("task: %v", task.Id, task.Name)
		}
	}
}
