package connections

import (
	"ginTemplate/internal/config"
	"ginTemplate/pkg/logging"
	"github.com/nats-io/nats.go"
)

func ConnectToNatsBroker() *nats.Conn {
	logging.DefaultLogger.Info().Msg("Connecting to nats...")
	if config.DefaultConfig.NatsDsn == nil {
		logging.DefaultLogger.Fatal().Msg("Nats DSN is nil")
		return nil
	}
	nc, _ := nats.Connect(*config.DefaultConfig.NatsDsn)
	return nc
}
