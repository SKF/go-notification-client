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
