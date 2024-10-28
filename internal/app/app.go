package app

import (
	"log/slog"
	grpcApp "thumbnail-proxy/internal/app/grpc"
	"thumbnail-proxy/internal/config/server"
	"thumbnail-proxy/internal/service"
	"thumbnail-proxy/internal/storage/redis"
)

type App struct {
	GrpcServer *grpcApp.Server
}

func New(
	cfg *server.Config,
	log *slog.Logger,
) *App {
	redisClient := redis.InitRedis(cfg)
	cache := redis.New(log, redisClient, cfg)

	tbService := service.New(log, cache, cache, cfg.Timeout)

	grpcServer := grpcApp.New(log, tbService, cfg.Port)

	return &App{GrpcServer: grpcServer}
}
