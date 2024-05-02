package repository

import (
	"github.com/americanas-go/config"
)

const (
	root          = "faas.datastore"
	eventProvider = root + ".event.provider"
)

func init() {
	config.Add(eventProvider, "noop", "event provider")
}

// EventProviderValue returns the event provider configured via the "faas.datastore.event.provider" key.
// If not configured, the default is nats.
func EventProviderValue() string {
	return config.String(eventProvider)
}
