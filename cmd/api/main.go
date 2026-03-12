package main

import (
	"log"
	"net/http"

	"golang-notification-service/internal/config"
	"golang-notification-service/internal/db"
	"golang-notification-service/internal/handler"
	"golang-notification-service/internal/repository"
	"golang-notification-service/internal/service"
)

func main() {
	cfg := config.Load()

	postgresDB, err := db.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer postgresDB.Close()

	healthHandler := handler.HealthHandler{
		DB: postgresDB,
	}

	notificationRepository := repository.NotificationRepository{
		DB: postgresDB,
	}

	notificationService := service.NotificationService{
		NotificationRepository: notificationRepository,
	}

	notificationHandler := handler.NotificationHandler{
		NotificationService: notificationService,
	}

	http.HandleFunc("/health", healthHandler.Check)
	http.HandleFunc("/notifications", notificationHandler.Create)
	http.HandleFunc("/notifications/", notificationHandler.GetByID)

	log.Println("API server running on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}