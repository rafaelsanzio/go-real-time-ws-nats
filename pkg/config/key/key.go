package key

import (
	"github.com/joho/godotenv"
	"github.com/rafaelsanzio/go-stream-live/pkg/applog"
	"github.com/rafaelsanzio/go-stream-live/pkg/errs"
)

type Key struct {
	Name   string
	Secure bool
	Provider
}

type Provider string

var (
	ProviderStore  = Provider("store")
	ProviderEnvVar = Provider("env")
)

var (
	WSURL       = Key{Name: "WS_URL", Secure: false, Provider: ProviderEnvVar}
	WSPort      = Key{Name: "WS_PORT", Secure: false, Provider: ProviderEnvVar}
	NatsURL     = Key{Name: "NATS_URL", Secure: false, Provider: ProviderEnvVar}
	NatsPort    = Key{Name: "NATS_PORT", Secure: false, Provider: ProviderEnvVar}
	WatcherPort = Key{Name: "WATCHER_PORT", Secure: false, Provider: ProviderEnvVar}
)

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		_ = errs.ErrGettingEnv.Throwf(applog.Log, errs.ErrFmt, err)
	}
}
