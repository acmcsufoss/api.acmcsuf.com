package announcements

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

var PostAnnouncement = &cobra.Command{
	Use:   "post",
	Short: "post a new announcement",

	Run: func(cmd *cobra.Command, args []string) {
		postAnnouncement(config.Cfg)
	},
}

func postAnnouncement(cfg *config.Config) {
	postUrl := config.GetBaseURL(cfg).JoinPath("v1", "announcements")

	payload, err := postForm()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}
	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: could not marshal JSON:", err)
		return
	}

	if body, err := client.SendRequestAndReadResponse(postUrl, true, http.MethodPost,
		bytes.NewBuffer(b)); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		if body != nil {
			// NOTE: This isn't going to stderr. Should probably fix that at some point
			utils.PrettyPrintJSON(body)
		}
	} else {
		utils.PrettyPrintJSON(body)
	}
}

func postForm() (*dto.Announcement, error) {
	var payload dto.Announcement
	var err error
	var (
		announceAtStr string
		channelIDStr  string
		messageIDStr  string
	)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Announcement ID").
				Value(&payload.Uuid).
				Validate(forms.ValidateNonEmpty()),
			huh.NewInput().
				Title("Announcement Visibility").
				Value(&payload.Visibility).
				Validate(forms.ValidateNonEmpty()),
			huh.NewInput().
				Title("Announcement Time\n"+
					"Format:  \x1b[93mMM/DD/YY HH:MM[PM | AM]\x1b[0m\n"+
					"Example: \x1b[93m01/02/06 03:04PM\x1b[0m").
				Value(&announceAtStr).
				Validate(forms.ValidateNonEmpty()),
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

	payload.AnnounceAt, err = utils.ByteSlicetoUnix([]byte(announceAtStr))
	if err != nil {
		return nil, err
	}
	payload.DiscordChannelID = &channelIDStr
	payload.DiscordMessageID = &messageIDStr

	return &payload, nil
}
