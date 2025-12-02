package positions

import (
	"github.com/spf13/cobra"
)

type positionFlags struct {
	oid      bool
	semester bool
	tier     bool
	fullname bool
	title    bool
	team     bool
}

var CLIPositions = &cobra.Command{
	Use:   "positions HEADER",
	Short: "A command to manage positions.",
}

func init() {
	CLIPositions.AddCommand(GetPositions)
	CLIPositions.AddCommand(PostPosition)
	CLIPositions.AddCommand(PutPosition)
	CLIPositions.AddCommand(DeletePosition)
}
