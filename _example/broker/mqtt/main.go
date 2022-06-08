package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-sohunjug/kratos-transport/broker"
	"github.com/go-sohunjug/kratos-transport/broker/mqtt"
)

const (
	EmqxBroker        = "tcp://broker.emqx.io:1883"
	EmqxCnBroker      = "tcp://broker-cn.emqx.io:1883"
	EclipseBroker     = "tcp://mqtt.eclipseprojects.io:1883"
	MosquittoBroker   = "tcp://test.mosquitto.org:1883"
	HiveMQBroker      = "tcp://broker.hivemq.com:1883"
	LocalEmqxBroker   = "tcp://127.0.0.1:1883"
	LocalRabbitBroker = "tcp://user:bitnami@127.0.0.1:1883"
)

func main() {
	ctx := context.Background()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	b := mqtt.NewBroker(
		broker.Addrs(LocalEmqxBroker),
		broker.OptionContext(ctx),
		mqtt.WithCleanSession(false),
		mqtt.WithAuth("user", "bitnami"),
		mqtt.WithClientId("test-client-2"),
	)

	defer func(b broker.Broker) {
		err := b.Disconnect()
		if err != nil {

		}
	}(b)

	if err := b.Connect(); err != nil {
		fmt.Println(err)
	}

	topic := "topic/bobo/#"
	//topicSharedGroup := "$share/g1/topic/bobo/#"
	//topicSharedQueue := "$queue/topic/bobo/#"

	_, _ = b.Subscribe(topic, receive,
		broker.SubscribeContext(ctx),
	)

	<-interrupt
}

func receive(_ context.Context, event broker.Event) error {
	fmt.Println("Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))
	//_ = event.Ack()
	return nil
}
