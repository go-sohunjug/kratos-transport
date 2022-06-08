package nsq

import (
	NSQ "github.com/nsqio/go-nsq"
	"github.com/go-sohunjug/kratos-transport/broker"
)

type publication struct {
	topic string
	m     *broker.Message
	nm    *NSQ.Message
	opts  broker.PublishOptions
	err   error
}

func (p *publication) Topic() string {
	return p.topic
}

func (p *publication) Message() *broker.Message {
	return p.m
}

func (p *publication) Ack() error {
	p.nm.Finish()
	return nil
}

func (p *publication) Error() error {
	return p.err
}
