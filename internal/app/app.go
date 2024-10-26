package app

import (
	"fmt"
	"log/slog"
	"thumbnail-proxy/internal/app/cli"
	grpcApp "thumbnail-proxy/internal/app/grpc"
	"thumbnail-proxy/internal/config"
)

type App struct {
	CliClient  *cli.CLI
	GrpcServer *grpcApp.Server
}

func New(
	cfg *config.Config,
	log *slog.Logger,
) (*App, error) {
	const op = "Server.New"

	cliClient, err := cli.New(cfg.Address, cfg, log)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	//TODO: implement thumbnailService
	grpcServer := grpcApp.New(log, nil, cfg.Port)

	return &App{
		CliClient:  cliClient,
		GrpcServer: grpcServer,
	}, nil
}
