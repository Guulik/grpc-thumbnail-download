package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"thumbnail-proxy/internal/config"
	"thumbnail-proxy/proto/gen/thumbnail"
)

type CLI struct {
	client thumbnailv1.ThumbnailClient
	cfg    *config.Config
	log    *slog.Logger
}

func New(
	addr string,
	cfg *config.Config,
	log *slog.Logger,
) (*CLI, error) {
	const op = "cli.New"

	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &CLI{
		client: thumbnailv1.NewThumbnailClient(cc),
		cfg:    cfg,
		log:    log,
	}, nil
}

func (c *CLI) Execute() error {
	const op = "cli.Execute"

	rootCmd := &cobra.Command{Use: "thumbnail"}
	rootCmd.AddCommand(c.get(c.cfg))
	rootCmd.AddCommand(c.output(c.cfg))

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return rootCmd.Execute()
}
