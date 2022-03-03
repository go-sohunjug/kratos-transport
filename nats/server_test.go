package nats

import (
	"context"
	"fmt"
	"github.com/tx7do/kratos-transport/common"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	ctx := context.Background()

	srv := NewServer(
		common.Addrs("nats://127.0.0.1:4222"),
		common.OptionContext(ctx),
	)

	if err := srv.Connect(); err != nil {
		panic(err)
	}

	if err := srv.Start(ctx); err != nil {
		panic(err)
	}

	_ = srv.RegisterSubscriber("test_topic", receive,
		common.SubscribeContext(ctx),
		common.Queue("test_topic"),
	)

	time.Sleep(time.Second * 60)

	if srv.Stop(ctx) != nil {
		t.Errorf("expected nil got %v", srv.Stop(ctx))
	}
}

func receive(event common.Event) error {
	fmt.Println("Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))
	return nil
}
