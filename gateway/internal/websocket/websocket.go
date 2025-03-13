package websocket

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/consul"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ProxyWebsocketWithConsul(c *gin.Context, serviceName string, pathPrefix string) {
	dialer := websocket.Dialer{
		NetDialContext:   consul.ConsulDialContext(serviceName),
		HandshakeTimeout: 10 * time.Second,
	}
	targetURL, err := url.Parse("ws://" + serviceName)
	if err != nil {
		log.Printf("Error parsing target URL: %v", err)
		return
	}
	targetURL.Path = strings.TrimPrefix(c.Request.URL.Path, pathPrefix)

	targetConn, resp, err := dialer.Dial(targetURL.String(), c.Request.Header)
	if err != nil {
		global.Logger.Error("Error dialing target websocket", zap.Error(err), zap.Any("response", resp))
		return
	}
	defer targetConn.Close()
	clientConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Logger.Error("Error upgrading client websocket", zap.Error(err))
		return
	}
	defer clientConn.Close()

	errChan := make(chan error, 2)

	go func() {
		for {
			messageType, message, err := clientConn.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			if err := targetConn.WriteMessage(messageType, message); err != nil {
				errChan <- err
				return
			}
		}
	}()

	go func() {
		for {
			messageType, message, err := targetConn.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			if err := clientConn.WriteMessage(messageType, message); err != nil {
				errChan <- err
				return
			}
		}
	}()

	<-errChan

}
