package officers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/client"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var PutOfficer = &cobra.Command{
	Use:   "put --id <uuid> [flags]",
	Short: "update an existing officer by id",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		putOfficer(id, config.Cfg)
	},
}

func init() {
	PutOfficer.Flags().String("id", "", "Officer ID to update")
	PutOfficer.MarkFlagRequired("id")
}

func putOfficer(id string, cfg *config.Config) {
	url := config.GetBaseURL(cfg).JoinPath("v1", "board", "officers", id)

	// ----- Get officer we want to update
	var oldPayload dbmodels.CreateOfficerParams
	if body, err := client.SendRequestAndReadResponse(url, false, http.MethodGet, nil); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		if body != nil {
			utils.PrettyPrintJSONErr(body)
		}
		return
	} else {
		if err := json.Unmarshal(body, &oldPayload); err != nil {
			fmt.Fprintln(os.Stderr, "Error: failed to unmarshal response body:", err)
			return
		}
	}

	// ----- Update found officer -----
	// Read new data
	newPayload, err := putForm(&oldPayload)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	b, err := json.Marshal(newPayload)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: failed to marshal data:", err)
		return
	}

	// Update remote resource with new data
	if body, err := client.SendRequestAndReadResponse(url, true, http.MethodPut,
		bytes.NewBuffer(b)); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		if body != nil {
			utils.PrettyPrintJSON(body)
		}
	} else {
		utils.PrettyPrintJSON(body)
	}
}

func putForm(oldPayload *dbmodels.CreateOfficerParams) (*dbmodels.UpdateOfficerParams, error) {
	var payload dbmodels.UpdateOfficerParams
	var err error
	var (
		name    string = oldPayload.FullName
		picture string = oldPayload.Picture.String
		github  string = oldPayload.Github.String
		discord string = oldPayload.Discord.String
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Full Name").
				Value(&name),
			huh.NewInput().
				Title("Picture URL").
				Value(&picture),
			huh.NewInput().
				Title("GitHub URL").
				Value(&github),
			huh.NewInput().
				Title("Discord URL").
				Value(&discord),
		),
	)
	if err = form.Run(); err != nil {
		return nil, err
	}

	payload.Uuid = oldPayload.Uuid
	payload.FullName = name
	payload.Picture = utils.StringtoNullString(picture)
	payload.Github = utils.StringtoNullString(github)
	payload.Discord = utils.StringtoNullString(discord)

	return &payload, nil
}
