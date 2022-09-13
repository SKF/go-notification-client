package main

import (
	"context"
	"fmt"
	"os"
	"time"

	dd_http "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"

	notification "github.com/SKF/go-notification-client/internal/models"
	"github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-rest-utility/client/auth"
	"github.com/SKF/go-utility/v2/stages"
)

const serviceName = "example-service"

type tokenProvider struct{}

func (t *tokenProvider) GetRawToken(ctx context.Context) (auth.RawToken, error) {
	return auth.RawToken(mustGetEnv("TOKEN")), nil
}

func main() {
	c := notification.New(
		notification.WithStage(stages.StageSandbox),
		client.WithDatadogTracing(dd_http.RTWithServiceName(serviceName)),
		client.WithTokenProvider(&tokenProvider{}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	notificationType, err := c.GetNotificationType(ctx, "name")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", notificationType)

	notificationTypes, err := c.GetNotificationTypes(ctx)
	if err != nil {
		panic(err)
	}

	for _, notificationType = range notificationTypes {
		fmt.Printf("%+v\n", notificationType)
	}
}

func mustGetEnv(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Errorf("environment variable %q is not set", key))
	}

	return value
}
