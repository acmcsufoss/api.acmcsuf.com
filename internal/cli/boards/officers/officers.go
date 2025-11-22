package officers

import (
	"github.com/spf13/cobra"
)

type officerFlags struct {
	uuid     bool
	fullname bool
	picture  bool
	github   bool
	discord  bool
}

var CLIOfficers = &cobra.Command{
	Use:   "officers HEADER",
	Short: "A command to manage officers.",
}

func init() {
	CLIOfficers.AddCommand(GetOfficers)
	CLIOfficers.AddCommand(DeleteOfficers)
	CLIOfficers.AddCommand(PostOfficer)
	CLIOfficers.AddCommand(PutOfficer)
}
