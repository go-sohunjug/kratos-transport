package kafka

import (
	"context"

	"github.com/go-sohunjug/kratos-transport/broker"
	KAFKA "github.com/segmentio/kafka-go"
)

type publication struct {
	topic  string
	err    error
	m      *broker.Message
	ctx    context.Context
	reader *KAFKA.Reader
	km     KAFKA.Message
}

func (p *publication) Topic() string {
	return p.topic
}

func (p *publication) Message() *broker.Message {
	return p.m
}

func (p *publication) Ack() error {
	return p.reader.CommitMessages(p.ctx, p.km)
}

func (p *publication) Error() error {
	return p.err
}

func (p *publication) Reply(m *broker.Message) { }
