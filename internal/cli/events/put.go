package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/client"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/forms"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/dto"
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
	PutEvents.Flags().String("id", "", "ID of the event to update")
	PutEvents.MarkFlagRequired("id")
}

func putEvents(id string, cfg *config.Config) {
	resourceURL := config.GetBaseURL(cfg).JoinPath("v1", "events", id)

	var oldPayload dto.Event
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

func putForm(oldPayload *dto.Event) (dto.UpdateEvent, error) {
	var payload dto.UpdateEvent
	var err error
	var (
		locationStr string = oldPayload.Location
		startAtStr  string = utils.FormatUnix(oldPayload.StartAt)
		endAtStr    string = utils.FormatUnix(oldPayload.EndAt)
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
					"Format: \x1b[93m01/02/06 03:04PM\x1b[0m").
				Value(&startAtStr),
			huh.NewInput().
				Title("Ends At\n"+
					"Format: \x1b[93m01/02/06 03:04PM\x1b[0m").
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
		return dto.UpdateEvent{}, err
	}

	payload.Location = forms.NonEmptyPtr(locationStr)
	startUnix, err := utils.ParseTime(startAtStr)
	if err != nil {
		return dto.UpdateEvent{}, err
	}
	payload.StartAt = &startUnix
	endUnix, err := utils.ParseTime(endAtStr)
	if err != nil {
		return dto.UpdateEvent{}, err
	}
	payload.EndAt = &endUnix
	payload.IsAllDay = &isAllDay
	payload.Host = &hostStr

	return payload, nil
}
