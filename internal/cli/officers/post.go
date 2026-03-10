package officers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/client"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/forms"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var PostOfficer = &cobra.Command{
	Use:   "post [flags]",
	Short: "Post a new officer",

	Run: func(cmd *cobra.Command, args []string) {
		postOfficer(config.Cfg)
	},
}

func init() {
	// Officer flags
	PostOfficer.Flags().StringP("uuid", "u", "", "Set uuid of this officer")
	PostOfficer.Flags().StringP("name", "n", "", "Set the full name of this officer")
	PostOfficer.Flags().StringP("picture", "p", "", "Set the picture of this officer")
	PostOfficer.Flags().StringP("github", "g", "", "Set the github of this officer")
	PostOfficer.Flags().StringP("discord", "d", "", "Set the discord of this officer")
}

func postOfficer(cfg *config.Config) {
	postUrl := config.GetBaseURL(cfg).JoinPath("v1", "board", "officers")

	payload, err := postForm()
	cobra.CheckErr(err)
	b, err := json.Marshal(payload)
	cobra.CheckErr(err)

	if body, err := client.SendRequestAndReadResponse(postUrl, true, http.MethodPost,
		bytes.NewBuffer(b)); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		if body != nil {
			utils.PrettyPrintJSONErr(body)
		}
	} else {
		utils.PrettyPrintJSON(body)
	}
}

func postForm() (*dbmodels.CreateOfficerParams, error) {
	var payload dbmodels.CreateOfficerParams
	var err error
	var (
		picture string
		github  string
		discord string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Officer ID").
				Value(&payload.Uuid).
				Validate(forms.ValidateNonEmpty()),
			huh.NewInput().
				Title("Full Name").
				Value(&payload.FullName).
				Validate(forms.ValidateNonEmpty()),
			huh.NewInput().
				Title("Picture URL").
				Value(&picture),
			huh.NewInput().
				Title("GitHub Username").
				Value(&github),
			huh.NewInput().
				Title("Discord Username").
				Value(&discord),
		),
	)
	if err = form.Run(); err != nil {
		return nil, err
	}

	// HACK: conversions required here due to lack of DTO models
	payload.Picture = utils.StringtoNullString(picture)
	payload.Github = utils.StringtoNullString(github)
	payload.Discord = utils.StringtoNullString(discord)

	return &payload, nil
}
