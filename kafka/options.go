package kafka

import (
	"github.com/americanas-go/config"
)

// Options can be used to create customized handler.
type Options struct {
	Brokers       []string
	Subjects      []string
	GroupId       string
	Concurrency   int
	QueueCapacity int
	MinBytes      int
	MaxBytes      int
	StartOffset   int64
}

// DefaultOptions returns options based in config.
func DefaultOptions() (*Options, error) {

	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
