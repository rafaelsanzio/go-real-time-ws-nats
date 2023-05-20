package config

import (
	"github.com/rafaelsanzio/go-stream-live/pkg/config/key"
	"github.com/rafaelsanzio/go-stream-live/pkg/errs"
)

type Service interface {
	Value(key.Key) (string, errs.AppError)
}
