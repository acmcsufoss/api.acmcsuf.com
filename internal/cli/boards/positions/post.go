package positions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var PostPosition = &cobra.Command{
	Use:   "post [flags]",
	Short: "Post a new tier",

	Run: func(cmd *cobra.Command, args []string) {

		// NOTE: Using update positions params since it covers all possible fields position in DB has,
		// compared to CreatePositionParams which only carries: oid, semester, and tier
		var payload models.UpdatePositionParams

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		title, _ := cmd.Flags().GetString("title")
		team, _ := cmd.Flags().GetString("team")

		payload.FullName, _ = cmd.Flags().GetString("name")
		payload.Title = utils.StringtoNullString(title)
		payload.Team = utils.StringtoNullString(team)
		payload.Oid, _ = cmd.Flags().GetString("oid")
		payload.Semester, _ = cmd.Flags().GetString("semester")
		payload.Tier, _ = cmd.Flags().GetInt64("tier")

		changedFlags := positionFlags{
			oid:      cmd.Flags().Lookup("oid").Changed,
			semester: cmd.Flags().Lookup("semester").Changed,
			tier:     cmd.Flags().Lookup("tier").Changed,
			fullname: cmd.Flags().Lookup("name").Changed,
			title:    cmd.Flags().Lookup("title").Changed,
			team:     cmd.Flags().Lookup("team").Changed,
		}

		postPosition(&payload, &changedFlags, host, port)
	},
}

func init() {
	// Url flags
	PostPosition.Flags().String("host", "127.0.0.1", "Set a custom host")
	PostPosition.Flags().String("port", "8080", "Set a custom port")

	// Position flags
	PostPosition.Flags().StringP("oid", "o", "", "Set the oid for position")
	PostPosition.Flags().StringP("semester", "s", "", "Set the semester for position")
	PostPosition.Flags().Int64P("tier", "t", 0, "Set the tier for position")
	PostPosition.Flags().StringP("name", "n", "", "Set the name of this position")
	PostPosition.Flags().StringP("title", "T", "", "Set the title of this position")
	PostPosition.Flags().StringP("team", "a", "", "Set the team  of this position")
}

func postPosition(payload *models.UpdatePositionParams, cf *positionFlags, host, port string) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	// Name
	for {
		if cf.fullname {
			break
		}

		fmt.Println("Please enter the full name of position holder:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.FullName = string(scanner.Bytes())
		break
	}

	// title
	for {
		if cf.title {
			break
		}

		fmt.Println("Please enter position title:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.Title = utils.StringtoNullString(string(scanner.Bytes()))
		break
	}

	// team
	for {
		if cf.team {
			break
		}

		fmt.Println("Please enter position team:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.Team = utils.StringtoNullString(string(scanner.Bytes()))
		break
	}
	// Oid
	for {
		if cf.oid {
			break
		}

		fmt.Println("Please enter position oid:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.Oid = string(scanner.Bytes())
		break
	}

	// Title
	for {
		if cf.semester {
			break
		}

		fmt.Println("Please enter the positions's semester:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.Semester = string(scanner.Bytes())
		break
	}

	// TIndex
	for {
		if cf.tier {
			break
		}

		fmt.Println("Please enter the positon's tier:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		tier, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			fmt.Println(err)
			continue
		}

		payload.Tier = int64(tier)
		break
	}

	// confirmation
	for {
		fmt.Println("Is your position data correct? If not, type n or no.")
		utils.PrintStruct(payload, false)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}

		confirmationBuffer := scanner.Bytes()
		confirmationBool, err := utils.YesOrNo(confirmationBuffer, scanner)
		if err != nil {
			fmt.Println("error with reading confirmation:", err)
		}
		if !confirmationBool {
			// Sorry :(
			return
		} else {
			break
		}
	}

	// marshal to json, and prepare url
	jsonPayload, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println("error formating payload to json: ", err)
		return
	}

	host = fmt.Sprint(host, ":", port)
	path := "v1/board/positions/"

	postURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// post payload
	response, err := http.Post(postURL.String(), "application/json", strings.NewReader(string(jsonPayload)))
	if err != nil {
		fmt.Println("error with post: ", err)
		return
	}

	if response == nil {
		fmt.Println("error, no response recieved")
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading body: ", err)
		return
	}

	fmt.Println(string(body))
}
