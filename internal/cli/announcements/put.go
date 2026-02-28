package announcements

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/oauth"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/dto"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var PutAnnouncements = &cobra.Command{
	Use:   "put --id <uuid>",
	Short: "update an existing announcement by its id",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		putAnnouncements(id, config.Cfg)
	},
}

func init() {
	PutAnnouncements.Flags().String("id", "", "Get an announcement by its id")
	PutAnnouncements.MarkFlagRequired("id")
}

func putAnnouncements(id string, cfg *config.Config) {
	resourceUrl := config.GetBaseURL(cfg).JoinPath("v1", "announcements", id)

	// ----- Get announcement we want to update -----
	var oldPayload dbmodels.CreateAnnouncementParams
	if body, err := client.SendRequestAndReadResponse(resourceUrl, false, http.MethodGet, nil); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		err = json.Unmarshal(body, &oldPayload)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: failed to unmarshal response body:", err)
			return
		}
	}

	// ----- Update found announceement -----
	var oldPayload dto.UpdateAnnouncement
	err = json.Unmarshal(body, &oldPayload)
	if err != nil {
		fmt.Println("Error: failed to unmarshal response body:", err)
		return
	}
	newPayload, err := putForm(id)
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
	if body, err := client.SendRequestAndReadResponse(resourceUrl, true, http.MethodPut,
		bytes.NewBuffer(b)); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		utils.PrettyPrintJSON(body)
	}
}

// TODO: Use DTO models instaad of dbmodels
func putForm(uuid string) (*dto.UpdateAnnouncement, error) {
	var payload dto.UpdateAnnouncement
	var err error
	var (
		visibilityStr string = "poop"
		announceAtStr string
		channelIDStr  string
		messageIDStr  string
	)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Announcement Visibility").
				Value(&visibilityStr),
			huh.NewInput().
				Title("Announcement Time\n"+
					"Format:  \x1b[93mMM/DD/YY HH:MM[PM | AM]\x1b[0m\n"+
					"Example: \x1b[93m01/02/06 03:04PM\x1b[0m").
				Value(&announceAtStr),
			// TODO: write validator for time inputs
			huh.NewInput().
				Title("Channel ID").
				Value(&channelIDStr),
			huh.NewInput().
				Title("Message ID").
				Value(&messageIDStr),
		),
	)
	if err = form.Run(); err != nil {
		return nil, err
	}

	// HACK: These conversions won't be necessary once we start using DTO models here
	payload.Visibility = &visibilityStr
	if announceAtStr != "" {
		timestamp, err := utils.ByteSlicetoUnix([]byte(announceAtStr))
		if err != nil {
			return nil, err
		}
		payload.AnnounceAt = &timestamp
	}

	payload.DiscordChannelID = &channelIDStr
	payload.DiscordMessageID = &messageIDStr

	return &payload, nil
}
