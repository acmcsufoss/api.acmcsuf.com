package events

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

var PostEvent = &cobra.Command{
	Use:   "post",
	Short: "Post a new event",

	Run: func(cmd *cobra.Command, args []string) {
		postEvent(config.Cfg)
	},
}

func postEvent(cfg *config.Config) {
	postUrl := config.GetBaseURL(cfg).JoinPath("v1", "events")

	payload, err := postForm()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: failed to marshal data:", err)
		return
	}

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

func postForm() (*dbmodels.CreateEventParams, error) {
	var payload dbmodels.CreateEventParams
	var err error
	var (
		startAtStr string
		endAtStr   string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Event ID").
				Value(&payload.Uuid).
				Validate(forms.ValidateNonEmpty()),
			huh.NewInput().
				Title("Location").
				Value(&payload.Location).
				Validate(forms.ValidateNonEmpty()),
			huh.NewInput().
				Title("Starts At\n"+
					"Format:  \x1b[93mMM/DD/YY HH:MM[PM | AM]\x1b[0m\n"+
					"Example: \x1b[93m01/02/06 03:04PM\x1b[0m").
				Value(&startAtStr).
				Validate(forms.ValidateNonEmpty()),
			huh.NewInput().
				Title("Ends At\n"+
					"Format:  \x1b[93mMM/DD/YY HH:MM[PM | AM]\x1b[0m\n"+
					"Example: \x1b[93m01/02/06 04:04PM\x1b[0m").
				Value(&endAtStr).
				Validate(forms.ValidateNonEmpty()),
			huh.NewConfirm().
				Title("All day event?").
				Value(&payload.IsAllDay),
			huh.NewInput().
				Title("Host").
				Value(&payload.Host).
				Validate(forms.ValidateNonEmpty()),
		),
	)
	if err = form.Run(); err != nil {
		return nil, err
	}

	payload.StartAt, err = utils.ByteSlicetoUnix([]byte(startAtStr))
	if err != nil {
		return nil, err
	}
	payload.EndAt, err = utils.ByteSlicetoUnix([]byte(endAtStr))
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
