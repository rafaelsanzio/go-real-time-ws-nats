package local

import (
	"os"

	"github.com/rafaelsanzio/go-stream-live/pkg/config/key"
)

var Values = map[key.Key]string{
	key.WSURL:  getDefaultOrEnvVar("ws://localhost", "WS_URL"),
	key.WSPort: getDefaultOrEnvVar("8080", "WS_PORT"),

	key.NatsURL:  getDefaultOrEnvVar("nats://go-stream-live_nats_1", "NATS_URL"),
	key.NatsPort: getDefaultOrEnvVar("4222", "NATS_PORT"),

	key.WatcherPort: getDefaultOrEnvVar("9090", "WATCHER_PORT"),
}

// Some of the db fields are set via env var in the makefile, so this optionally uses those to prevent test failures in jenkins
func getDefaultOrEnvVar(dfault, envVar string) string {
	val := os.Getenv(envVar)
	if val != "" {
		return val
	}
	return dfault
}
