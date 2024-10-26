package cli

import (
	"github.com/spf13/cobra"
	"thumbnail-proxy/internal/config"
)

func (c *CLI) output(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "output [path]",
		Short: "Set path to save preview",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cfg.SetOutputDir(args[0])
		},
	}
	return cmd
}
