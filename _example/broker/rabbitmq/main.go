package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-sohunjug/kratos-transport/broker"
	"github.com/go-sohunjug/kratos-transport/broker/rabbitmq"
)

func main() {
	ctx := context.Background()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := rabbitmq.NewBroker(
		broker.Addrs("amqp://user:bitnami@127.0.0.1:5672"),
		broker.OptionContext(ctx),
	)

	_ = b.Init()

	if err := b.Connect(); err != nil {
		fmt.Println(err)
	}

	_, _ = b.Subscribe("test_queue.*", receive,
		broker.SubscribeContext(ctx),
		broker.Queue("test_queue.*"),
		// broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
	)

	<-interrupt
}

func receive(_ context.Context, event broker.Event) error {
	fmt.Println("Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))
	//_ = event.Ack()
	return nil
}
