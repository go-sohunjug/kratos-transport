package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"log"

	"github.com/go-sohunjug/kratos-transport/broker"
	"github.com/go-sohunjug/kratos-transport/transport/kafka"
)

func main() {
	ctx := context.Background()

	kafkaSrv := kafka.NewServer(
		kafka.Address("127.0.0.1:9092"),
		kafka.Subscribe(ctx, "test_topic", "a-group", false, receive),
	)

	app := kratos.New(
		kratos.Name("kafka"),
		kratos.Server(
			kafkaSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Println(err)
	}
}

func receive(_ context.Context, event broker.Event) error {
	fmt.Println("Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))
	return nil
}
