package client_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	client "github.com/SKF/go-notification-client"
	internal_models "github.com/SKF/go-notification-client/internal/models"
	"github.com/SKF/go-notification-client/models"
	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/problems"
)

func Test_BaseURL(t *testing.T) {
	t.Parallel()

	c := client.New()

	require.NotNil(t, c.Client.BaseURL)
	assert.Equal(t, "api.notification.iot.enlight.skf.com", c.Client.BaseURL.Host)

	c = client.New(client.WithStage("sandbox"))

	require.NotNil(t, c.Client.BaseURL)
	assert.Equal(t, "api.notification.sandbox.iot.enlight.skf.com", c.Client.BaseURL.Host)
}

func Test_GetNotificationType(t *testing.T) {
	given := internal_models.ModelsNotificationTypeResponse{
		Name:                               "notification",
		QualifyingAction:                   "SERVICE::ACTION",
		DefaultDeliveryChannel:             "sns",
		DefaultDeliveryScheduleType:        "cron",
		DailyDeliveryTimeHour:              1,
		DailyDeliveryTimeMinute:            30,
		MinimumSecondsBetweenNotifications: 60,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(given)
		require.NoError(t, err)
	}))
	defer server.Close()

	c := client.New(rest.WithBaseURL(server.URL))

	actual, err := c.GetNotificationType(context.TODO(), "notification")
	require.NoError(t, err)

	expected := models.NotificationType{
		Name:                               "notification",
		QualifyingAction:                   "SERVICE::ACTION",
		DefaultDeliveryChannel:             "sns",
		DefaultDeliveryScheduleType:        "cron",
		DailyDeliveryTimeHour:              1,
		DailyDeliveryTimeMinute:            30,
		MinimumSecondsBetweenNotifications: 60,
	}

	assert.Equal(t, expected, actual)
}

func Test_GetNotificationType_ErrorResponse(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		problem := problems.Validation(problems.ValidationReason{
			Name:   "body",
			Reason: "invalid",
			Cause:  fmt.Errorf("this is the issue"),
		})

		problems.WriteResponse(context.TODO(), problem, w, r)
	}))
	defer server.Close()

	c := client.New(rest.WithBaseURL(server.URL))

	_, err := c.GetNotificationType(context.TODO(), "notification")
	require.Error(t, err)

	var problem problems.ValidationProblem

	if assert.ErrorAs(t, err, &problem) {
		assert.Equal(t, http.StatusBadRequest, problem.BasicProblem.Status)

		if assert.Len(t, problem.Reasons, 1) {
			assert.Equal(t, "body", problem.Reasons[0].Name)
		}
	}
}

func Test_GetNotificationTypes(t *testing.T) {
	given := []internal_models.ModelsNotificationTypeResponse{
		{
			Name:                               "notification",
			QualifyingAction:                   "SERVICE::ACTION",
			DefaultDeliveryChannel:             "sns",
			DefaultDeliveryScheduleType:        "cron",
			DailyDeliveryTimeHour:              1,
			DailyDeliveryTimeMinute:            30,
			MinimumSecondsBetweenNotifications: 60,
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(given)
		require.NoError(t, err)
	}))
	defer server.Close()

	c := client.New(rest.WithBaseURL(server.URL))

	actual, err := c.GetNotificationTypes(context.TODO())
	require.NoError(t, err)

	expected := []models.NotificationType{
		{
			Name:                               "notification",
			QualifyingAction:                   "SERVICE::ACTION",
			DefaultDeliveryChannel:             "sns",
			DefaultDeliveryScheduleType:        "cron",
			DailyDeliveryTimeHour:              1,
			DailyDeliveryTimeMinute:            30,
			MinimumSecondsBetweenNotifications: 60,
		},
	}

	require.Len(t, actual, len(expected))

	assert.Equal(t, expected, actual)
}

func Test_GetNotificationTypes_ErrorResponse(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		problem := problems.Validation(problems.ValidationReason{
			Name:   "body",
			Reason: "invalid",
			Cause:  fmt.Errorf("this is the issue"),
		})

		problems.WriteResponse(context.TODO(), problem, w, r)
	}))
	defer server.Close()

	c := client.New(rest.WithBaseURL(server.URL))

	_, err := c.GetNotificationTypes(context.TODO())
	require.Error(t, err)

	var problem problems.ValidationProblem

	if assert.ErrorAs(t, err, &problem) {
		assert.Equal(t, http.StatusBadRequest, problem.BasicProblem.Status)

		if assert.Len(t, problem.Reasons, 1) {
			assert.Equal(t, "body", problem.Reasons[0].Name)
		}
	}
}
