package stomp

import (
	"github.com/go-sohunjug/kratos-transport/broker"
	"github.com/go-stomp/stomp/v3"
)

type publication struct {
	msg    *stomp.Message
	m      *broker.Message
	broker *stompBroker
	topic  string
	err    error
}

func (p *publication) Ack() error {
	return p.broker.stompConn.Ack(p.msg)
}

func (p *publication) Error() error {
	return p.err
}

func (p *publication) Topic() string {
	return p.topic
}

func (p *publication) Message() *broker.Message {
	return p.m
}

func (p *publication) Reply(m *broker.Message) { }
