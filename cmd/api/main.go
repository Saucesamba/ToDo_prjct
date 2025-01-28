package main

import (
	"MyProject/internal/config"
	"MyProject/internal/db"
	"MyProject/internal/handlers"
	"log"
	"net/http"
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

	var handler = handlers.NewHandler(*dbConn)
	http.HandleFunc("/register", handler.HandleUserRegister)
	http.HandleFunc("/login", handler.HandleUserLogin)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
