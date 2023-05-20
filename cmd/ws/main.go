package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rafaelsanzio/go-stream-live/pkg/applog"
	"github.com/rafaelsanzio/go-stream-live/pkg/config"
	"github.com/rafaelsanzio/go-stream-live/pkg/config/key"
	"github.com/rafaelsanzio/go-stream-live/pkg/errs"
	"github.com/rafaelsanzio/go-stream-live/pkg/websocket"
)

func main() {
	key.LoadEnvVars()

	wsPort, err_ := config.Value(key.WSPort)
	if err_ != nil {
		_ = errs.ErrGettingEnvWebSocketPort.Throwf(applog.Log, errs.ErrFmt, err_)
	}

	http.HandleFunc("/ws", websocket.Handler)
	log.Printf("WebSocket server running on port 8080...\n")

	wsCompletePort := fmt.Sprintf(":%s", wsPort)

	err := http.ListenAndServe(wsCompletePort, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
