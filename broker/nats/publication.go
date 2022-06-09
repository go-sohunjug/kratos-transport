package nats

import "github.com/go-sohunjug/kratos-transport/broker"

type publication struct {
	topic   string
	reply string
	err error
	s *natsBroker
	m   *broker.Message
}

func (p *publication) Topic() string {
	return p.topic
}

func (p *publication) Message() *broker.Message {
	return p.m
}

func (p *publication) Ack() error {
	return nil
}

func (p *publication) Error() error {
	return p.err
}

func (p *publication) Reply(m *broker.Message) {
	p.s.Publish(p.reply, m)
}
