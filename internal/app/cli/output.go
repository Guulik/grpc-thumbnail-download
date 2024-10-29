package cli

import (
	"github.com/spf13/cobra"
	"thumbnail-proxy/internal/config/cli"
)

func (cli *CLI) outputCommand(cfg *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "output [path]",
		Short: "Set path to save preview. Use default to set ./downloads/",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "default" {
				cfg.SetOutputDir("./downloads")
			}
			cfg.SetOutputDir(args[0])
		},
	}
	return cmd
}
