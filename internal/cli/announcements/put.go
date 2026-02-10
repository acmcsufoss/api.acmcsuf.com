package announcements

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "github.com/charmbracelet/huh"
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
}

func putAnnouncements(id string, cfg *config.Config) {
	var payload dbmodels.UpdateAnnouncementParams
	oldResourceURL := config.GetBaseURL(cfg).JoinPath("v1", "announcements", id)

	// ----- Get the Announcement We Want to Update -----
	client := http.Client{}
	getReq, err := oauth.NewRequestWithAuth(http.MethodGet, oldResourceURL.String(), nil)
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
		fmt.Println("Error reading response body:", err)
		return
	}

	var oldPayload dbmodels.CreateAnnouncementParams
	err = json.Unmarshal(body, &oldPayload)
	if err != nil {
		fmt.Println("Error: coudln't unmarshal response body:", err)
		return
	}

	// TODO: implement form

	// ----- Marshal Payload to Json -----
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	putRequest, err := oauth.NewRequestWithAuth(http.MethodPut, oldResourceURL.String(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Problem with PUT:", err)
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
