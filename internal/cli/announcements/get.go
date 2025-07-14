package announcements

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

var GetAnnouncement = &cobra.Command{
	Use:   "get",
	Short: "Get an announcement",

	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		uuid, _ := cmd.Flags().GetString("id")

		getAnnouncement(host, port, uuid)
	},
}

func init() {

	// Url flags
	GetAnnouncement.Flags().String("host", "127.0.0.1", "Set a custom host (Defaults to: 127.0.0.1)")
	GetAnnouncement.Flags().String("port", "8080", "Set a custom port (Defaults to: 8080)")
	GetAnnouncement.Flags().String("id", "", "Get a specific event by it's id")

}

func getAnnouncement(host string, port string, uuid string) {
	// ----- Constructing the url -----
	host = fmt.Sprint(host, ":", port)
	path := "announcements"
	if uuid != "" {
		path = fmt.Sprint(path, "/", uuid)
	}

	getUrl := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// ----- Requesting Get -----
	request, err := http.Get(getUrl.String())
	if err != nil {
		fmt.Println("error with request:", err)
		return
	}
	defer request.Body.Close()

	fmt.Println("Response status:", request.Status)

	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("error reading body:", err)
		return
	}

	fmt.Println(string(body))

}
