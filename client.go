package client

import (
	"fmt"

	rest "github.com/SKF/go-rest-utility/client"
	"github.com/SKF/go-utility/v2/stages"
)

type API interface{}

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
