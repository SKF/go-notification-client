package models

import (
	internal_models "github.com/SKF/go-notification-client/internal/models"
)

type NotificationType struct {
	Name                               string
	QualifyingAction                   string
	DefaultDeliveryChannel             string
	DefaultDeliveryScheduleType        string
	MinimumSecondsBetweenNotifications int64
}

func (n *NotificationType) FromInternal(internal internal_models.ModelsNotificationTypeResponse) error {
	n.Name = internal.Name
	n.QualifyingAction = internal.QualifyingAction
	n.DefaultDeliveryChannel = internal.DefaultDeliveryChannel
	n.DefaultDeliveryScheduleType = internal.DefaultDeliveryScheduleType
	n.MinimumSecondsBetweenNotifications = internal.MinimumSecondsBetweenNotifications

	return nil
}
