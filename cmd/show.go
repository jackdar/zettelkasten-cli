package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current ZKDIR",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", viper.GetString("ZKDIR"))
	},
}
