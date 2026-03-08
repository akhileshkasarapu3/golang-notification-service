package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type HealthHandler struct {
	DB *sql.DB
}

func (h HealthHandler) Check(w http.ResponseWriter, r *http.Request){
	err := h.DB.Ping()  	// Ping the DB.		(struct for persistant details).(DB connection).Ping
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		_ = json.NewEncoder(w).Encode(map[string]string{		// Encode: convert maps/struct to json objects
																// NewEncoder: pipeline to client/user 
			"status": "error",
			"message": "database connection field",
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string {
		"status" : "ok",
		"message": "API and Database are healthy",
	})
}