package officers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/client"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var DeleteOfficers = &cobra.Command{
	Use:   "delete --id <uuid>",
	Short: "Delete an officer with their id",

	Run: func(cmd *cobra.Command, args []string) {
		var uuid string
		uuid, _ = cmd.Flags().GetString("id")
		deleteOfficer(uuid, config.Cfg)
	},
}

func init() {
	DeleteOfficers.Flags().String("id", "", "Delete an officer by their id")
	DeleteOfficers.MarkFlagRequired("id")
}

func deleteOfficer(id string, cfg *config.Config) {
	deleteUrl := config.GetBaseURL(cfg).JoinPath("v1", "board", "officers", id)

	if body, err := client.SendRequestAndReadResponse(deleteUrl, http.MethodDelete); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		utils.PrettyPrintJSON(body)
	}
}
