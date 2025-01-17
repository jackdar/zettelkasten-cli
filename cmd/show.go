package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jackdar/zettelkasten-cli/internal/config"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current config",
	Run: func(cmd *cobra.Command, args []string) {
		config.PrintConfiguration(*conf)
	},
}
