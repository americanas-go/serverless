package noop

import (
	"context"
	"github.com/americanas-go/faas/repository"
)

// NewEvent returns an initialized noop client that implements event repository.
func NewEvent(ctx context.Context) repository.Event {
	return NewClient()
}
