package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jackdar/zettelkasten-cli/internal/config"
	"github.com/jackdar/zettelkasten-cli/internal/utils"
)

var (
	cfgFile string
	conf    = &config.Configuration{}

	rootCmd = &cobra.Command{
		Use:   "zk <command> [flags]",
		Short: "zk is a tool for managing Zettelkasten notes",
		Long:  `zk is a tool for managing Zettelkasten notes allowing the creation, editing, and searching of notes.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(func() {
		var err error
		conf, err = config.NewConfig(cfgFile)
		utils.CheckErr(err)
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zk)")
	rootCmd.PersistentFlags().StringVarP(&conf.ZkDir, "directory", "d", "", "Zettelkasten home directory (default is $HOME/zettelkasten)")
}
