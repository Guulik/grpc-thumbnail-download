package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"thumbnail-proxy/internal/app"
	"thumbnail-proxy/internal/app/cli"
	clicfg "thumbnail-proxy/internal/config/cli"
	"thumbnail-proxy/internal/config/server"
	"thumbnail-proxy/internal/lib/logger/handlers/slogpretty"
)

const (
	envLocal = "local"
	envProd  = " prod"
)

func main() {
	cfg := server.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application",
		slog.Any("cfg", cfg),
	)

	a := app.New(cfg, log)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		a.GrpcServer.MustRun()
	}()
	wg.Wait()

	cliCfg := clicfg.MustLoad()
	cliClient, err := cli.New(cliCfg)
	if err != nil {
		panic(err)
	}

	go func() {
		cliClient.MustExecute()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	a.GrpcServer.Stop()

	fmt.Println("Gracefully stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
