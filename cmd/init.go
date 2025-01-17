package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Zettelkasten directory",
	Long:  "Initializes the home Zettlekasten directory (default is $HOME/zettlekasten)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Initializing!\n")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
