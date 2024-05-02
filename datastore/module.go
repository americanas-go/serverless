package datastore

import (
	"github.com/americanas-go/faas/datastore/gcp/pubsub"
	"github.com/americanas-go/faas/datastore/noop"
	"sync"

	"github.com/americanas-go/faas/datastore/aws/kinesis"
	"github.com/americanas-go/faas/datastore/aws/sns"
	"github.com/americanas-go/faas/datastore/aws/sqs"
	"github.com/americanas-go/faas/datastore/nats"
	"github.com/americanas-go/faas/repository"
	"github.com/americanas-go/log"
	"go.uber.org/fx"
)

var eventOnce sync.Once

// EventModule returns fx module for initialization of event module based on the configured event provider.
// It can be: nats (default), kinesis, sns or sqs.
//
// The module is only loaded once.
func EventModule() fx.Option {

	options := fx.Options()

	eventOnce.Do(func() {

		value := repository.EventProviderValue()
		log.Debugf("Loading event provider %s", value)
		switch value {
		case "kinesis":
			options = kinesis.Module()
		case "sns":
			options = sns.Module()
		case "sqs":
			options = sqs.Module()
		case "pubsub":
			options = pubsub.Module()
		case "nats":
			options = nats.Module()
		default:
			options = noop.Module()
		}

	})

	return options
}
