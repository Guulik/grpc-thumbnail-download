package cli

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"thumbnail-proxy/internal/config/cli"
	"thumbnail-proxy/proto/gen/thumbnail"
)

type CLI struct {
	client thumbnailv1.ThumbnailClient
	cfg    *cli.Config
}

func New(
	cfg *cli.Config,
) (*CLI, error) {
	const op = "cli.New"

	cc, err := grpc.NewClient(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &CLI{
		client: thumbnailv1.NewThumbnailClient(cc),
		cfg:    cfg,
	}, nil
}

func (cli *CLI) Execute() error {
	const op = "cli.Execute"

	ctx := context.Background()
	rootCmd := &cobra.Command{Use: "thumbnail"}
	rootCmd.AddCommand(cli.getCommand(ctx))
	rootCmd.AddCommand(cli.outputCommand(cli.cfg))

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (cli *CLI) MustExecute() {
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
