package main

import (
	"context"
	"thumbnail-proxy/internal/app/cli"
	clicfg "thumbnail-proxy/internal/config/cli"
)

func main() {
	cfg := clicfg.MustLoad()

	cliClient, err := cli.New(cfg)
	if err != nil {
		panic(err)
	}

	cliClient.Execute(context.Background())
}
