package stream

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"

	"github.com/rafaelsanzio/go-stream-live/pkg/applog"
	"github.com/rafaelsanzio/go-stream-live/pkg/errs"
)

type Nats struct {
	URL  string
	PORT string
}

func (n Nats) Connect(ctx context.Context) *nats.Conn {
	nc, err := nats.Connect(fmt.Sprintf("%s:%s", n.URL, n.PORT))
	if err != nil {
		_ = errs.ErrNatsConnection.Throwf(applog.Log, errs.ErrFmt, err)
		return nil
	}

	return nc
}
