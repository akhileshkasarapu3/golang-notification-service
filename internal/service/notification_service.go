package service

import (
	"context"
	"errors"
	"strings"

	"golang-notification-service/internal/model"
	"golang-notification-service/internal/repository"
)

type NotificationService struct {
	NotificationRepository repository.NotificationRepository
}


// Create Service to handle missing fileld
func (s NotificationService) CreateNotification(ctx context.Context, req model.CreateNotificationRequest) (int64, error) {
	if strings.TrimSpace(req.RecipientEmail) == "" {
		return 0, errors.New("recipient_email is required")
	}

	if req.TemplateID <= 0 {
		return 0, errors.New("template_id must be greater than 0")
	}

	if req.Payload == nil {
		return 0, errors.New("payload is required")
	}

	return s.NotificationRepository.Create(ctx, req)
}


// Get Notification by ID
func (s NotificationService) GetNotificationByID(ctx context.Context, id int64) (model.NotificationResponse, error) {
	if id <= 0 {
		return model.NotificationResponse{}, errors.New("id must be greater than 0")
	}

	return s.NotificationRepository.GetByID(ctx, id)
}
