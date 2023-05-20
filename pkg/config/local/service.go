package local

import (
	"github.com/rafaelsanzio/go-stream-live/pkg/config/key"
	"github.com/rafaelsanzio/go-stream-live/pkg/errs"
)

type Service struct{}

func (s Service) Value(key key.Key) (string, errs.AppError) {
	return Values[key], nil
}
