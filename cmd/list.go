package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/jackdar/zettelkasten-cli/internal/utils"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List notes created on the same day",
	Run: func(cmd *cobra.Command, args []string) {
		listNotes()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listNotes() {
	dir, err := os.ReadDir(filepath.Join(conf.ZkDir, "__Inbox"))
	utils.CheckErr(err)

	today := time.Now().Format("2006-01-02")

	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		utils.CheckErr(err)

		if info.ModTime().Format("2006-01-02") == today {
			fmt.Printf(" - %s\n", entry.Name())
		}
	}
}
