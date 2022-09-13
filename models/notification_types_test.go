package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internal_models "github.com/SKF/go-notification-client/internal/models"
	"github.com/SKF/go-notification-client/models"
)

func Test_NotificationType(t *testing.T) {
	tests := []struct {
		name     string
		expected models.NotificationType
		given    internal_models.ModelsNotificationTypeResponse
	}{
		{
			name:     "Test the empty struct",
			expected: models.NotificationType{},
			given:    internal_models.ModelsNotificationTypeResponse{},
		},
		{
			name: "Verify all fields are mapped",
			expected: models.NotificationType{
				Name:                               "notification",
				QualifyingAction:                   "SERVICE::ACTION",
				DefaultDeliveryChannel:             "sns",
				DefaultDeliveryScheduleType:        "cron",
				DailyDeliveryTimeHour:              1,
				DailyDeliveryTimeMinute:            30,
				MinimumSecondsBetweenNotifications: 60,
			},
			given: internal_models.ModelsNotificationTypeResponse{
				Name:                               "notification",
				QualifyingAction:                   "SERVICE::ACTION",
				DefaultDeliveryChannel:             "sns",
				DefaultDeliveryScheduleType:        "cron",
				DailyDeliveryTimeHour:              1,
				DailyDeliveryTimeMinute:            30,
				MinimumSecondsBetweenNotifications: 60,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			var actual models.NotificationType

			err := actual.FromInternal(tt.given)
			require.NoError(t, err)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
