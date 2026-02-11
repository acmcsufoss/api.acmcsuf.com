package announcements

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/oauth"
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
	resourceURL := config.GetBaseURL(cfg).JoinPath("v1", "announcements", id)

	// ----- Get the Announcement We Want to Update -----
	client := http.Client{}
	getReq, err := oauth.NewRequestWithAuth(http.MethodGet, resourceURL.String(), nil)
	if err != nil {
		fmt.Printf("Error: couldn't retrieve resource %s: %s", id, err)
		return
	}
	getRes, err := client.Do(getReq)
	if err != nil {
		fmt.Println("Error: failed to send request:", err)
		return
	}
	defer getRes.Body.Close()
	if getRes.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP", getRes.Status)
		return
	}
	body, err := io.ReadAll(getRes.Body)
	if err != nil {
		fmt.Println("Error: failed to read response body:", err)
		return
	}

	// ----- Update found announceement -----
	var oldPayload dbmodels.CreateAnnouncementParams
	err = json.Unmarshal(body, &oldPayload)
	if err != nil {
		fmt.Println("Error: failed to unmarshal response body:", err)
		return
	}
	newPayload, err := putForm(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	jsonPayload, err := json.Marshal(newPayload)
	if err != nil {
		fmt.Println("Error: failed to marshal data:", err)
		return
	}
	putRequest, err := oauth.NewRequestWithAuth(http.MethodPut, resourceURL.String(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error: failed to contruct request:", err)
		return
	}
	putResponse, err := client.Do(putRequest)
	if err != nil {
		fmt.Println("Error: failed to send request: ", err)
		return
	}
	defer putResponse.Body.Close()

	if putResponse.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP", putResponse.Status)
		return
	}
	body, err = io.ReadAll(putResponse.Body)
	if err != nil {
		fmt.Println("Error: failed to read response body", err)
		return
	}
	utils.PrettyPrintJSON(body)
}

// TODO: Use DTO models instaad of dbmodels
func putForm(uuid string) (*dbmodels.UpdateAnnouncementParams, error) {
	var payload dbmodels.UpdateAnnouncementParams
	var err error
	var (
		visibilityStr string
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

	payload.Uuid = uuid
	// HACK: These conversions won't be necessary once we start using DTO models here
	if visibilityStr != "" {
		payload.Visibility = utils.StringtoNullString(visibilityStr)
	}
	if announceAtStr != "" {
		timestamp, err := utils.ByteSlicetoUnix([]byte(announceAtStr))
		if err != nil {
			return nil, err
		}
		payload.AnnounceAt = utils.Int64toNullInt64(timestamp)
	}
	if channelIDStr != "" {
		payload.DiscordChannelID = utils.StringtoNullString(channelIDStr)
	}
	if messageIDStr != "" {
		payload.DiscordMessageID = utils.StringtoNullString(messageIDStr)
	}
	return &payload, nil
}
