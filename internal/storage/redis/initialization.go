package redis

import (
	"github.com/redis/go-redis/v9"
	"log/slog"
	serverCfg "thumbnail-proxy/internal/config/server"
)

type Cache struct {
	log   *slog.Logger
	redis *redis.Client
	cfg   *serverCfg.Config
}

func New(log *slog.Logger, redis *redis.Client, cfg *serverCfg.Config) *Cache {
	return &Cache{
		log:   log,
		redis: redis,
		cfg:   cfg,
	}
}

func InitRedis(c *serverCfg.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
	return client
}
