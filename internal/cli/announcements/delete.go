package announcements

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var DeleteAnnouncements = &cobra.Command{
	Use:   "delete",
	Short: "delete an announcement by its id",

	Run: func(cmd *cobra.Command, args []string) {
		uuid, _ := cmd.Flags().GetString("id")
		deleteAnnouncement(uuid, config.Cfg)
	},
}

func init() {
	DeleteAnnouncements.Flags().String("id", "", "delete an announcement by its id")
	DeleteAnnouncements.MarkFlagRequired("id")
}

func deleteAnnouncement(id string, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	if id == "" {
		fmt.Println("ID is required to delete an announcement! Please use the --id flag")
		return
	}

	// ----- Constructing Url -----
	deleteUrl := baseURL.JoinPath("v1/announcements/", id)

	// ----- Delete -----
	request, err := http.NewRequest(http.MethodDelete, deleteUrl.String(), nil)
	if err != nil {
		fmt.Println("error with delete request:", err)
		return
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error with delete response:", err)
		return
	}

	if response == nil {
		fmt.Println("no response received")
		return
	}

	defer response.Body.Close()

	// ----- Read Response Information -----
	fmt.Println("Request status:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error with reading body:", err)
		return
	}

	fmt.Println(string(body))
}
