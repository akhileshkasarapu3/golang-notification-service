package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"golang-notification-service/internal/model"
)

type NotificationRepository struct {
	DB *sql.DB
}

func (r NotificationRepository) Create(ctx context.Context, req model.CreateNotificationRequest) (int64, error) {
	payloadBytes, err := json.Marshal(req.Payload)
	if err != nil {
		return 0, err
	}

	query := `
		INSERT INTO notifications (recipient_email, template_id, payload_json, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var notificationID int64
	err = r.DB.QueryRowContext(
		ctx,
		query,
		req.RecipientEmail,
		req.TemplateID,
		payloadBytes,
		"PENDING",
	).Scan(&notificationID)
	if err != nil {
		return 0, err
	}

	return notificationID, nil
}

// Get By ID from the database
func (r NotificationRepository) GetByID(ctx context.Context, id int64) (model.NotificationResponse, error) {
	query := `
		SELECT id, recipient_email, template_id, payload_json, status, error_message
		FROM notifications
		WHERE id = $1
	`

	var result model.NotificationResponse
	var payloadBytes []byte

	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&result.ID,
		&result.RecipientEmail,
		&result.TemplateID,
		&payloadBytes,
		&result.Status,
		&result.ErrorMessage,
	)
	if err != nil {
		return model.NotificationResponse{}, err
	}

	err = json.Unmarshal(payloadBytes, &result.Payload)
	if err != nil {
		return model.NotificationResponse{}, err
	}

	return result, nil
}



