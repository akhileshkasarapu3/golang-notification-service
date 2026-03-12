package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"golang-notification-service/internal/config"
)

func NewPostgresConnection(cfg config.Config) (*sql.DB, error) {			// Uses ptr so that the variable doesn't change always
	// creates DB connection string
	dsn := fmt.Sprintf(
		"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDB,
	)

	// opens connection pool
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	

	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(5)
	database.SetConnMaxLifetime(5 * time.Minute)


	// pings DB to confirm it actually works
	err= database.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL successfully!")
	return database, nil	// this function is to connect to db by passing config details
}

