package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new [<note name>] [flags]",
	Short: "Create a new Zettelkasten note.",
	Long:  "Create a new Zettelkasten note using the name provided in the Inbox.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		editor := os.Getenv("EDITOR")
		fpath := viper.GetString("ZKDIR") + "/__Inbox/" + args[0] + ".md"

		open := exec.Command(editor, fpath, "+normal i#  ", "+startinsert")

		open.Stdout = os.Stdout
		open.Stdin = os.Stdin
		open.Stderr = os.Stderr

		if err := open.Run(); err != nil {
			log.Fatal(err)
		}
	},
}
