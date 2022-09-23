package models

import (
	internal_models "github.com/SKF/go-notification-client/internal/models"
)

type InitiatedNotification struct {
	ID               string
	ExternalID       string
	Body             string
	CreatedBy        string
	Header           string
	NotificationType NotificationType
	ResourceID       string
	TriggeringUserID string
}

func (g *InitiatedNotification) FromInternal(internal internal_models.ModelsGetInitiatedNotificationResponse) error {
	g.ID = internal.ID
	g.ExternalID = internal.ExternalID
	g.Body = internal.Body
	g.CreatedBy = internal.CreatedBy
	g.Header = internal.Header
	g.ResourceID = internal.ResourceID
	g.TriggeringUserID = internal.TriggeringUserID

	if internal.NotificationType != nil {
		if err := g.NotificationType.FromInternal(*internal.NotificationType); err != nil {
			return err
		}
	}

	return nil
}

func (g *InitiatedNotification) ToInternal() internal_models.ModelsPostInitiatedNotificationRequest {
	internal := internal_models.ModelsPostInitiatedNotificationRequest{}

	internal.ID = g.ID
	internal.Body = g.Body
	internal.CreatedBy = g.CreatedBy
	internal.Header = g.Header
	internal.NotificationTypeName = g.NotificationType.Name
	internal.ResourceID = g.ResourceID

	return internal
}
