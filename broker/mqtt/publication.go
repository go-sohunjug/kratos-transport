package mqtt

import "github.com/go-sohunjug/kratos-transport/broker"

type publication struct {
	topic string
	msg   *broker.Message
	err   error
}

func (m *publication) Ack() error {
	return nil
}

func (m *publication) Error() error {
	return m.err
}

func (m *publication) Topic() string {
	return m.topic
}

func (m *publication) Message() *broker.Message {
	return m.msg
}

func (m *publication) Reply(msg *broker.Message) { }
