package client

import (
	"context"
	"fmt"

	internal_models "github.com/SKF/go-notification-client/internal/models"
	"github.com/SKF/go-notification-client/models"
	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
)

type API interface {
	GetNotificationType(context.Context, string) (models.NotificationType, error)
	GetNotificationTypes(context.Context) ([]models.NotificationType, error)
	PostInitiatedNotifiction(ctx context.Context, initialNotification models.InitiatedNotification) (string, error)
	GetInitiatedNotifiction(ctx context.Context, externalID string) (models.InitiatedNotification, error)
	DeleteInitiatedNotifiction(ctx context.Context, externalID string) error
}

type Client struct {
	*rest.Client
}

var _ API = &Client{Client: nil}

func WithStage(stage string) rest.Option {
	if stage == stages.StageProd {
		return rest.WithBaseURL("https://api.notification.iot.enlight.skf.com")
	}

	return rest.WithBaseURL(fmt.Sprintf("https://api.notification.%s.iot.enlight.skf.com", stage))
}

func New(opts ...rest.Option) *Client {
	restClient := rest.NewClient(
		append([]rest.Option{
			// Defaults to production stage if no option is supplied
			WithStage(stages.StageProd),
			rest.WithProblemDecoder(&ProblemDecoder{}),
		}, opts...)...,
	)

	return &Client{restClient}
}

func (c *Client) GetNotificationType(ctx context.Context, name string) (models.NotificationType, error) {
	request := rest.Get("v1/notification-types/{name}").
		Assign("name", name).
		SetHeader("Accept", "application/json")

	var response internal_models.ModelsNotificationTypeResponse

	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.NotificationType{}, fmt.Errorf("getting notification types failed: %w", err)
	}

	notificationType := models.NotificationType{} //nolint:exhaustruct

	if err := notificationType.FromInternal(response); err != nil {
		return models.NotificationType{}, fmt.Errorf("converting notification type failed: %w", err)
	}

	return notificationType, nil
}

func (c *Client) GetNotificationTypes(ctx context.Context) ([]models.NotificationType, error) {
	request := rest.Get("v1/notification-types").
		SetHeader("Accept", "application/json")

	var response []internal_models.ModelsNotificationTypeResponse

	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return nil, fmt.Errorf("getting notification types failed: %w", err)
	}

	notificationTypes := make([]models.NotificationType, 0, len(response))

	for i := range response {
		var notificationType models.NotificationType

		if err := notificationType.FromInternal(response[i]); err != nil {
			return nil, fmt.Errorf("converting notification type failed: %w", err)
		}

		notificationTypes = append(notificationTypes, notificationType)
	}

	return notificationTypes, nil
}

func (c Client) PostInitiatedNotifiction(ctx context.Context, initialNotification models.InitiatedNotification) (string, error) {
	request := rest.Post("v1/initiated-notifications").
		WithJSONPayload(initialNotification.ToInternal()).
		SetHeader("Accept", "application/json")

	var response internal_models.ModelsPostInitiatedNotificationResponse

	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return "", fmt.Errorf("posting initiated notification failed: %w", err)

	}

	return response.ExternalID, nil
}

func (c Client) GetInitiatedNotifiction(ctx context.Context, externalID string) (models.InitiatedNotification, error) {
	request := rest.Get("v1/initiated-notifications/{externalId}").
		Assign("externalId", externalID).
		SetHeader("Accept", "application/json")

	var response internal_models.ModelsGetInitiatedNotificationResponse

	if err := c.DoAndUnmarshal(ctx, request, &response); err != nil {
		return models.InitiatedNotification{}, fmt.Errorf("getting initiated notification failed: %w", err)
	}

	initiatedNotification := models.InitiatedNotification{}

	if err := initiatedNotification.FromInternal(response); err != nil {
		return models.InitiatedNotification{}, fmt.Errorf("converting initiated notification failed: %w", err)
	}

	return initiatedNotification, nil
}

func (c Client) DeleteInitiatedNotifiction(ctx context.Context, externalID string) error {
	request := rest.Delete("v1/initiated-notifications/{externalId}").
		Assign("externalId", externalID)

	if _, err := c.Do(ctx, request); err != nil {
		return fmt.Errorf("deleting initiated notification failed: %w", err)
	}

	return nil
}
