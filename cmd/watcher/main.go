package main

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"

	"github.com/rafaelsanzio/go-stream-live/pkg/applog"
	"github.com/rafaelsanzio/go-stream-live/pkg/config"
	"github.com/rafaelsanzio/go-stream-live/pkg/config/key"
	"github.com/rafaelsanzio/go-stream-live/pkg/errs"
	"github.com/rafaelsanzio/go-stream-live/pkg/stream"
)

func main() {
	key.LoadEnvVars()

	err := Watcher("stockexchange", func(message []byte) {
		log.Printf("NATS - Received message: %s\n", string(message))
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Watcher(topic string, alertFunc func(message []byte)) error {
	ctx := context.Background()

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

	nc := natsStream.Connect(ctx)
	defer nc.Close()

	// Build an event channel to the topic
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe(topic, ch)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	// Wait for new messages and call the alert function
	for msg := range ch {
		alertFunc(msg.Data)
	}

	return nil
}
