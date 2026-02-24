package announcements

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/forms"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/oauth"
	dto_request "github.com/acmcsufoss/api.acmcsuf.com/internal/dto/request"
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
	payload, err := postForm()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error: could not marshal JSON:", err)
		return
	}

	postURL := config.GetBaseURL(cfg).JoinPath("v1", "announcements")
	client := http.Client{}
	req, err := oauth.NewRequestWithAuth(http.MethodPost, postURL.String(), strings.NewReader(string(jsonPayload)))
	if err != nil {
		fmt.Println("Error: could not create request:", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: could not send request:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP", res.Status)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: could not read response body:", err)
		return
	}
	utils.PrettyPrintJSON(body)
}

// TODO: Use DTO models instaad of dbmodels
func postForm() (*dto_request.Announcement, error) {
	var payload dto_request.Announcement
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

	// HACK: These conversions won't be necessary once we start using DTO models here
	payload.AnnounceAt, err = utils.ByteSlicetoUnix([]byte(announceAtStr))
	if err != nil {
		return nil, err
	}
	payload.DiscordChannelID = &channelIDStr
	payload.DiscordMessageID = &messageIDStr

	return &payload, err
}
