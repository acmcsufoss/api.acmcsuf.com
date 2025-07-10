package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csuf",
	Short: "A CLI tool to help manage the API of the CSUF ACM website",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CSUF API TOOL")
	},
}

func Execute() {

	// Logging the error, prefix is date, time, and what file the log is from
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err := rootCmd.Execute(); err != nil {
		log.Println("Error with CLI:", err)
		os.Exit(1)
	}

}

func init() {
	//rootCmd.AddCommand
}
