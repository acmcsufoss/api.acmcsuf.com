package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config <subcommand>",
	Short: "Manage this tool's persistent configuration.",
}

func init() {
	ConfigCmd.AddCommand(configInitCmd)
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create default configuration file",

	Run: func(cmd *cobra.Command, args []string) {
		if err := createDefaultConfigFile(); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating default config file: %v\n", err)
		}
	},
}

// TODO: implement `config show` subcommand
