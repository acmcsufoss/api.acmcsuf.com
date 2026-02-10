package announcements

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/oauth"
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
	payload, err := form()
	if err != nil {
		fmt.Println("Error: could not read form data")
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
func form() (*dbmodels.CreateAnnouncementParams, error) {
	var payload dbmodels.CreateAnnouncementParams
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
				Validate(func(id string) error {
					if id == "" {
						return errors.New("ID must not be empty")
					}
					return nil
				}),
			huh.NewInput().
				Title("Announcement visibility").
				Value(&payload.Visibility),
			huh.NewInput().
				Title("Announcement time").
				Value(&announceAtStr),
			huh.NewInput().
				Title("Channel ID").
				Value(&channelIDStr),
			huh.NewInput().
				Title("Message ID").
				Value(&messageIDStr),
		),
	)

	err := form.Run()

	return &payload, err
}
