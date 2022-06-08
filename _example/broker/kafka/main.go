package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-sohunjug/kratos-transport/broker"
	"github.com/go-sohunjug/kratos-transport/broker/kafka"
)

func main() {
	ctx := context.Background()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := kafka.NewBroker(
		broker.Addrs("localhost:9092"),
		broker.OptionContext(ctx),
	)

	_, _ = b.Subscribe("test_topic", receive,
		broker.SubscribeContext(ctx),
		broker.Queue("mt-group"),
	)

	<-interrupt
}

func receive(_ context.Context, event broker.Event) error {
	fmt.Println("Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))
	//_ = event.Ack()
	return nil
}
