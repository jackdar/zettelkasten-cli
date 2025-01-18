package cmd

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jackdar/zettelkasten-cli/internal/utils"
)

var obsidian bool

func init() {
	newCmd.Flags().BoolVarP(&obsidian, "obsidian", "o", false, "use Obsidian to open the new note (default false).")
	viper.BindPFlag("USE_OBSIDIAN", newCmd.Flags().Lookup("obsidian"))

	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   "new [<note name>] [flags]",
	Short: "Create a new Zettelkasten note.",
	Long:  "Create a new Zettelkasten note using the name provided in the Inbox.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := getTitle(args)

		fpath := conf.NewNoteDir + "/" + title + ".md"

		if !obsidian {
			runEditor(exec.Command(conf.Editor, fpath, "+normal i# "+title+"\n\n", "+startinsert"))
		} else {
			runObsidian(title)
		}
	},
}

func getTitle(args []string) string {
	if len(args) > 0 {
		return args[0]
	}

	form := huh.NewForm(huh.NewGroup(huh.NewInput().Key("title").Title("What is the note title?")))

	if err := form.Run(); err != nil {
		utils.CheckErr(err)
	}

	return form.GetString("title")
}

func runEditor(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		utils.CheckErr(err)
	}
}

func runObsidian(title string) {
	vault := filepath.Base(conf.ZkDir)
	dir := filepath.Base(conf.NewNoteDir)

	newNote, err := os.Create(conf.NewNoteDir + "/" + title + ".md")
	utils.CheckErr(err)

	_, err = io.WriteString(newNote, "# "+title+"\n\n")
	utils.CheckErr(err)

	cmd := exec.Command(os.Getenv("SHELL"), "-c", "open 'obsidian://open?vault="+vault+"&file="+dir+"/"+title+"'")

	if err := cmd.Run(); err != nil {
		utils.CheckErr(err)
	}
}
