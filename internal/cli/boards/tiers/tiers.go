package tiers

import (
	"github.com/spf13/cobra"
)

type tierFlags struct {
	tier   bool
	title  bool
	tindex bool
	team   bool
}

var CLITiers = &cobra.Command{
	Use:   "tiers HEADER",
	Short: "A command to manage tiers.",
}

func init() {
	CLITiers.AddCommand(GetTiers)
	CLITiers.AddCommand(PostTier)
	CLITiers.AddCommand(PutTier)
	CLITiers.AddCommand(DeleteTier)
}
