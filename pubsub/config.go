package pubsub

import (
	"github.com/americanas-go/config"
)

const (
	root                   = "faas.pubsub"
	subscription           = root + ".subscription"
	synchronous            = root + ".synchronous"
	concurrency            = root + ".concurrency"
	maxoutstandingmessages = root + ".maxoutstandingmessages"
)

func init() {
	config.Add(subscription, "changeme", "pubsub subscription name of target topic")
	config.Add(synchronous, false, "set false to enable concurrency pulling of messages. Otherwise, NumGoroutines will be set to 1")
	config.Add(concurrency, 16, "pubsub goroutine concurrency")
	config.Add(maxoutstandingmessages, 8, "limits the number of concurrent handlers of messages")
}
