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

var PutEvents = &cobra.Command{
	Use:   "put --id <uuid>",
	Short: "Update an existing event by its id",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		putEvents(id, config.Cfg)
	},
}

func init() {
	PutEvents.Flags().String("id", "", "Get an event by its id")
	PutEvents.MarkFlagRequired("id")
}

func putEvents(id string, cfg *config.Config) {
	resourceURL := config.GetBaseURL(cfg).JoinPath("v1", "events", id)

	var oldPayload dbmodels.CreateEventParams
	if body, err := client.SendRequestAndReadResponse(resourceURL, false, http.MethodGet, nil); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		if body != nil {
			utils.PrettyPrintJSONErr(body)
		}
		return
	} else {
		err = json.Unmarshal(body, &oldPayload)
		cobra.CheckErr(err)
	}

	newPayload, err := putForm(&oldPayload)
	cobra.CheckErr(err)
	b, err := json.Marshal(newPayload)
	cobra.CheckErr(err)

	if body, err := client.SendRequestAndReadResponse(resourceURL, true, http.MethodPut,
		bytes.NewBuffer(b)); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		if body != nil {
			utils.PrettyPrintJSONErr(body)
		}
	} else {
		utils.PrettyPrintJSON(body)
	}
}

func putForm(oldPayload *dbmodels.CreateEventParams) (*dbmodels.UpdateEventParams, error) {
	var payload dbmodels.UpdateEventParams
	var err error
	var (
		locationStr string = oldPayload.Location
		startAtStr  string
		endAtStr    string
		isAllDay    bool   = oldPayload.IsAllDay
		hostStr     string = oldPayload.Host
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Location").
				Value(&locationStr).
				Validate(forms.ValidateNonEmpty()),
			huh.NewInput().
				Title("Starts At\n"+
					"Format:  \x1b[93mMM/DD/YY HH:MM[PM | AM]\x1b[0m\n"+
					"Leave empty to keep existing value").
				Value(&startAtStr),
			huh.NewInput().
				Title("Ends At\n"+
					"Format:  \x1b[93mMM/DD/YY HH:MM[PM | AM]\x1b[0m\n"+
					"Leave empty to keep existing value").
				Value(&endAtStr),
			huh.NewConfirm().
				Title("All day event?").
				Value(&isAllDay),
			huh.NewInput().
				Title("Host").
				Value(&hostStr).
				Validate(forms.ValidateNonEmpty()),
		),
	)
	if err = form.Run(); err != nil {
		return nil, err
	}

	payload.Uuid = oldPayload.Uuid
	payload.Location = utils.StringtoNullString(locationStr)
	payload.IsAllDay = utils.BooltoNullBool(isAllDay)
	payload.Host = utils.StringtoNullString(hostStr)
	if startAtStr != "" {
		timestamp, err := utils.ByteSlicetoUnix([]byte(startAtStr))
		if err != nil {
			return nil, err
		}
		payload.StartAt = utils.Int64toNullInt64(timestamp)
	}
	if endAtStr != "" {
		timestamp, err := utils.ByteSlicetoUnix([]byte(endAtStr))
		if err != nil {
			return nil, err
		}
		payload.EndAt = utils.Int64toNullInt64(timestamp)
	}

	return &payload, nil
}
