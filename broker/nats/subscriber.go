package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/go-sohunjug/kratos-transport/broker"
)

type subscriber struct {
	s    *nats.Subscription
	opts broker.SubscribeOptions
}

func (s *subscriber) Options() broker.SubscribeOptions {
	return s.opts
}

func (s *subscriber) Topic() string {
	return s.s.Subject
}

func (s *subscriber) Unsubscribe() error {
	return s.s.Unsubscribe()
}
