package connections

import (
	"ginTemplate/internal/config"
	"ginTemplate/pkg/logging"
	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() *redis.Client {
	logging.DefaultLogger.Info().Msg("Connecting to redis...")
	if config.DefaultConfig.RedisDsn == nil {
		logging.DefaultLogger.Fatal().Msg("Redis DSN is nil")
		return nil
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: *config.DefaultConfig.RedisDsn,
	})
	return rdb
}
