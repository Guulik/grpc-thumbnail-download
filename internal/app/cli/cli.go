package cli

import (
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

func (c *CLI) Execute() error {
	const op = "cli.Execute"

	rootCmd := &cobra.Command{Use: "thumbnail"}
	rootCmd.AddCommand(c.get(c.cfg.OutputDir))
	rootCmd.AddCommand(c.output(c.cfg))

	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return rootCmd.Execute()
}

func (c *CLI) MustExecute() {
	if err := c.Execute(); err != nil {
		panic(err)
	}
}
