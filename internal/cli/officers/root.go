package officers

import (
	"github.com/spf13/cobra"
)

var CLIOfficers = &cobra.Command{
	Use:   "officers",
	Short: "A command to manage officers.",
}

func init() {
	CLIOfficers.AddCommand(GetOfficers)
	CLIOfficers.AddCommand(DeleteOfficers)
	CLIOfficers.AddCommand(PostOfficer)
	CLIOfficers.AddCommand(PutOfficer)
}
