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