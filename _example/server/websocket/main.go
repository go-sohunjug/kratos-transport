package main

import (
	"fmt"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-sohunjug/kratos-transport/transport/websocket"
)

func main() {
	//ctx := context.Background()

	wsSrv := websocket.NewServer(
		websocket.Address(":8800"),
		websocket.ReadHandle("/ws", handleMessage),
		websocket.ConnectHandle(handleConnect),
	)

	app := kratos.New(
		kratos.Name("websocket"),
		kratos.Server(
			wsSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Println(err)
	}
}

func handleConnect(connectionId string, register bool) {
	if register {
		fmt.Printf("%s connected\n", connectionId)
	} else {
		fmt.Printf("%s disconnect\n", connectionId)
	}
}

func handleMessage(connectionId string, message *websocket.Message) (*websocket.Message, error) {
	fmt.Printf("[%s] Payload: %s\n", connectionId, string(message.Body))

	var relyMsg websocket.Message
	relyMsg.Body = []byte("hello")

	return &relyMsg, nil
}
