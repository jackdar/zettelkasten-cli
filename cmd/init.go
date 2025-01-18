package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Zettelkasten directory",
	Long:  "Initializes the home Zettlekasten directory (default is $HOME/zettlekasten)",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			return
		}

		if _, err := os.ReadDir(filepath.Join(home, "Zettelkasten")); err == nil {
			fmt.Println("Zettelkasten directory already exists!")
			return
		}

		dirs := []string{
			filepath.Join(home, "Zettelkasten"),
			filepath.Join(home, "Zettelkasten", "__Inbox"),
			filepath.Join(home, "Zettelkasten", "_Zettelkasten"),
			filepath.Join(home, "Zettelkasten", "Templates"),
		}

		for _, dir := range dirs {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
		}

		fmt.Println("Zettelkasten directory structure initialized successfully!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
