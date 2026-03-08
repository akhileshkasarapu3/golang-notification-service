package main

import (
	"log"
	"net/http"
	"golang-notification-service/internal/config"
	"golang-notification-service/internal/db"
	"golang-notification-service/internal/handler"
)

func main(){
	cfg := config.Load()	// 1. Get the config details

	postgresDB, err := db.NewPostgresConnection(cfg)	// 2. get the database
	if err != nil {
		log.Fatalf("Failed to connect to Database")
	}
	defer postgresDB.Close()

	healthHandler := handler.HealthHandler{
		DB: postgresDB,
	}

	http.HandleFunc("/health", healthHandler.Check)		// 3. Ping Database and Check

	log.Println("API Server running on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start server: %w", err)
	}
}