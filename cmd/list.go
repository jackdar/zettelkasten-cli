package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/jackdar/zettelkasten-cli/internal/utils"
)

func init() {

}

func initDirectory() {
	dirName := viper.GetString("ZKDIR")
	dir, err := os.ReadDir(dirName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Zettelkasten directory not found! Creating...\n")

		err := createDirectory(dirName)
		utils.CheckErr(err)
	}

	for _, entry := range dir {
		fmt.Printf(" - %s\n", entry.Name())
	}
}

func createDirectory(dir string) error {
	err := os.Mkdir(dir, 0755)
	utils.CheckErr(err)

	err = os.Mkdir(dir+"/__Inbox", 0755)
	utils.CheckErr(err)

	err = os.Mkdir(dir+"/_Zettelkasten", 0755)
	utils.CheckErr(err)

	return err
}
