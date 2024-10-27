package app

import (
	"log/slog"
	grpcApp "thumbnail-proxy/internal/app/grpc"
	"thumbnail-proxy/internal/config/server"
	"thumbnail-proxy/internal/service"
)

type App struct {
	GrpcServer *grpcApp.Server
}

func New(
	cfg *server.Config,
	log *slog.Logger,
) *App {
	const op = "App.New"

	//TODO: implement storage
	tbService := service.New(log, nil, nil, cfg.Timeout)

	grpcServer := grpcApp.New(log, tbService, cfg.Port)

	return &App{GrpcServer: grpcServer}
}
