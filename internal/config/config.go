package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/viper"

	"github.com/jackdar/zettelkasten-cli/internal/utils"
)

type Configuration struct {
	ZkDir        string `mapstructure:"ZKDIR"`
	Editor       string `mapstructure:"EDITOR"`
	NewNoteDir   string `mapstructure:"NEW_NOTE_DIR"`
	TemplatesDir string `mapstructure:"TEMPLATES_DIR"`
	UseObsidian  bool   `mapstructure:"USE_OBSIDIAN"`
}

func NewConfig(cfgFile string) (*Configuration, error) {
	var config *Configuration

	setDefaults()

	if cfgFile != "" {
		viper.SetConfigType("env")
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("$HOME")
		viper.SetConfigType("env")
		viper.SetConfigName(".zk")
	}

	viper.ReadInConfig()

	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.New("Could not unmarshal config file!")
	}

	return config, nil
}

func setDefaults() {
	home, err := os.UserHomeDir()
	utils.CheckErr(err)

	viper.SetDefault("ZKDIR", home+"/Zettelkasten")
	viper.SetDefault("EDITOR", "vim")
	viper.SetDefault("NEW_NOTE_DIR", home+"/Zettelkasten/__Inbox")
	viper.SetDefault("TEMPLATES_DIR", home+"/Zettelkasten/Templates")
	viper.SetDefault("USE_OBSIDIAN", false)
}

func PrintConfiguration(config Configuration) {
	val := reflect.ValueOf(config)
	typ := reflect.TypeOf(config)

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i).Interface()
		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}
}
