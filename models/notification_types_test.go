package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internal_models "github.com/SKF/go-notification-client/internal/models"
	"github.com/SKF/go-notification-client/models"
)

func Test_NotificationType(t *testing.T) {
	t.Parallel()

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
				MinimumSecondsBetweenNotifications: 60,
			},
			given: internal_models.ModelsNotificationTypeResponse{
				Name:                               "notification",
				QualifyingAction:                   "SERVICE::ACTION",
				DefaultDeliveryChannel:             "sns",
				DefaultDeliveryScheduleType:        "cron",
				MinimumSecondsBetweenNotifications: 60,
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var actual models.NotificationType

			err := actual.FromInternal(test.given)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}
