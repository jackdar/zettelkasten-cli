package utils

import (
	"errors"
	"fmt"
	"os"
)

func CheckErr(msg interface{}) {
	if msg != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
		os.Exit(1)
	}
}

func CheckFileExists(fpath string) error {
	_, err := os.Stat(fpath)
	if os.IsNotExist(err) {
		return errors.New("File already exists!")
	}

	return nil

}
