package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	zkDir   string

	rootCmd = &cobra.Command{
		Use:   "zk <command> [flags]",
		Short: "zk is a tool for managing Zettelkasten notes",
		Long:  `zk is a tool for managing Zettelkasten notes allowing the creation, editing, and searching of notes.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Add a command to run
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&zkDir, "dir", "", "Zettelkasten home directory (default is $HOME/Zettelkasten)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zk)")

	viper.SetDefault("ZKDIR", "$HOME/zettelkasten")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("$HOME")
		viper.SetConfigType("env")
		viper.SetConfigName(".zk")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
