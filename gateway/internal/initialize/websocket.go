package initialize

import (
	"github.com/bhtoan2204/gateway/global"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func InitWebSocket() {
	// Subscribe to a channel
	pubsub := global.Redis.Subscribe(global.Ctx, "socket_channel")
	// Initialize WebSocket
	defer pubsub.Close()

	_, err := pubsub.Receive(global.Ctx)
	if err != nil {
		panic(err)
	}

	// ch := pubsub.Channel()

}
