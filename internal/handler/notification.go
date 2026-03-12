package handler

import (
	"encoding/json"
	"net/http"

	"golang-notification-service/internal/model"
	"golang-notification-service/internal/service"
)

type NotificationHandler struct {
	NotificationService service.NotificationService
}

func (h NotificationHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "method not allowed",
		})
		return
	}

	var req model.CreateNotificationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	notificationID, err := h.NotificationService.CreateNotification(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"notification_id": notificationID,
		"status":          "PENDING",
	})
}