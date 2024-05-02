package noop

import (
	"sync"

	"go.uber.org/fx"
)

var once sync.Once

// Module loads the noop module providing an initialized client.
//
// The module is only loaded once.
func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			fx.Provide(
				NewEvent,
			),
		)
	})

	return options
}
