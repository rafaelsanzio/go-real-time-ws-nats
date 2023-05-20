package websocket

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/rafaelsanzio/go-stream-live/pkg/applog"
	"github.com/rafaelsanzio/go-stream-live/pkg/config"
	"github.com/rafaelsanzio/go-stream-live/pkg/config/key"
	"github.com/rafaelsanzio/go-stream-live/pkg/errs"
	"github.com/rafaelsanzio/go-stream-live/pkg/stream"
)

type WebSocket struct {
	URL string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (ws WebSocket) Connect() *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(ws.URL, nil)
	if err != nil {
		_ = errs.ErrWebSocketConnection.Throwf(applog.Log, errs.ErrFmt, err)
		return nil
	}

	return conn
}

func (ws WebSocket) WriteMessage(conn *websocket.Conn, data []byte) {
	err := conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		_ = errs.ErrWebSocketWriteMessage.Throwf(applog.Log, errs.ErrFmt, err)
		return
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		_ = errs.ErrWebSocketUpgrader.Throwf(applog.Log, errs.ErrFmt, err)
		return
	}

	natsURL, err_ := config.Value(key.NatsURL)
	if err_ != nil {
		_ = errs.ErrGettingEnvNatsURL.Throwf(applog.Log, errs.ErrFmt, err_)
	}

	natsPort, err_ := config.Value(key.NatsPort)
	if err_ != nil {
		_ = errs.ErrGettingEnvNatsPort.Throwf(applog.Log, errs.ErrFmt, err_)
	}

	natsStream := stream.Nats{
		URL:  natsURL,
		PORT: natsPort,
	}

	natsConn := natsStream.Connect(ctx)
	defer natsConn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			_ = errs.ErrWebSocketReadMessage.Throwf(applog.Log, errs.ErrFmt, err)
			break
		}

		log.Printf("WS - Received message: %s\n", message)

		// TODO: Send the message to NATS
		subject := "any-topic"
		natsConn.Publish(subject, message)
	}
}
