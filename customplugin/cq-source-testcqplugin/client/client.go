package client

import (
	"context"
	"fmt"

	"github.com/akshat-crypto/cq-source-testcqplugin/internal/xkcd"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type Client struct {
	Logger zerolog.Logger
	XKCD   *xkcd.Client
}

func (c *Client) ID() string {
	// TODO: Change to either your plugin name or a unique dynamic identifier
	return "xkcddemo"
}

func New(ctx context.Context, logger zerolog.Logger, s specs.Source, opts source.Options) (schema.ClientMeta, error) {
	var pluginSpec Spec

	if err := s.UnmarshalSpec(&pluginSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin spec: %w", err)
	}
	// TODO: Add your client initialization here
	c, err := xkcd.NewClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		Logger: logger,
		XKCD:   c,
	}, nil
}
