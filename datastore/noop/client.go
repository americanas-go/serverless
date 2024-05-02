package noop

import (
	"context"
	v2 "github.com/cloudevents/sdk-go/v2"
)

// Client represents a noop client that implements the event repository.
type Client struct {
}

// NewClient creates a new noop client.
func NewClient() *Client {
	return &Client{}
}

// Publish publishes an event slice.
func (p *Client) Publish(ctx context.Context, outs []*v2.Event) (err error) {
	return nil
}
