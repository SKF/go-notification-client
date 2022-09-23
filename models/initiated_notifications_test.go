package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internal_models "github.com/SKF/go-notification-client/internal/models"
	"github.com/SKF/go-notification-client/models"
)

func Test_InitiatedNotification_FromInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected models.InitiatedNotification
		given    internal_models.ModelsGetInitiatedNotificationResponse
	}{
		{
			name:     "Test the empty struct",
			expected: models.InitiatedNotification{},
			given:    internal_models.ModelsGetInitiatedNotificationResponse{},
		},
		{
			name: "Verify all fields are mapped",
			expected: models.InitiatedNotification{
				ID:         "foo",
				ExternalID: "external_foo",
				Body:       "body",
				CreatedBy:  "waldo",
				Header:     "header",
				NotificationType: models.NotificationType{
					Name:                               "bar",
					QualifyingAction:                   "SERVICE::ACTION",
					DefaultDeliveryChannel:             "sns",
					DefaultDeliveryScheduleType:        "cron",
					MinimumSecondsBetweenNotifications: 60,
				},
				ResourceID:       "foo_resource",
				TriggeringUserID: "waldo_id",
			},
			given: internal_models.ModelsGetInitiatedNotificationResponse{
				Body:       "body",
				CreatedBy:  "waldo",
				ExternalID: "external_foo",
				Header:     "header",
				ID:         "foo",
				NotificationType: &internal_models.ModelsNotificationTypeResponse{
					DefaultDeliveryChannel:             "sns",
					DefaultDeliveryScheduleType:        "cron",
					MinimumSecondsBetweenNotifications: 60,
					Name:                               "bar",
					QualifyingAction:                   "SERVICE::ACTION",
				},
				ResourceID:       "foo_resource",
				TriggeringUserID: "waldo_id",
			},
		},
	}
	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			var actual models.InitiatedNotification

			err := actual.FromInternal(test.given)
			require.NoError(t, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_InitiatedNotification_ToInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expected internal_models.ModelsPostInitiatedNotificationRequest
		given    models.InitiatedNotification
	}{
		{
			name:     "Test the empty struct",
			expected: internal_models.ModelsPostInitiatedNotificationRequest{},
			given:    models.InitiatedNotification{},
		},
		{
			name: "Verify all fields are mapped",
			expected: internal_models.ModelsPostInitiatedNotificationRequest{
				Body:                 "body",
				CreatedBy:            "waldo",
				Header:               "header",
				ID:                   "foo",
				NotificationTypeName: "bar",
				ResourceID:           "foo_id",
			},
			given: models.InitiatedNotification{
				ID:         "foo",
				ExternalID: "",
				Body:       "body",
				CreatedBy:  "waldo",
				Header:     "header",
				NotificationType: models.NotificationType{
					Name:                               "bar",
					QualifyingAction:                   "",
					DefaultDeliveryChannel:             "",
					DefaultDeliveryScheduleType:        "",
					MinimumSecondsBetweenNotifications: 0,
				},
				ResourceID:       "foo_id",
				TriggeringUserID: "",
			},
		},
	}
	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			actual := test.given.ToInternal()

			assert.Equal(t, test.expected, actual)
		})
	}
}
