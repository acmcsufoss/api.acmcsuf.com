package tiers

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

var PostTier = &cobra.Command{
	Use:   "post [flags]",
	Short: "Post a new tier",

	Run: func(cmd *cobra.Command, args []string) {
		var payload models.CreateTierParams

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		title, _ := cmd.Flags().GetString("")
		tindex, _ := cmd.Flags().GetInt64("tindex")
		team, _ := cmd.Flags().GetString("team")

		payload.Tier, _ = cmd.Flags().GetInt64("tier")
		payload.Title = utils.StringtoNullString(title)
		payload.TIndex = utils.Int64toNullInt64(tindex)
		payload.Team = utils.StringtoNullString(team)

		changedFlags := tierFlags{
			tier:   cmd.Flags().Lookup("tier").Changed,
			title:  cmd.Flags().Lookup("title").Changed,
			tindex: cmd.Flags().Lookup("tindex").Changed,
			team:   cmd.Flags().Lookup("team").Changed,
		}

		postTier(&payload, &changedFlags, host, port)
	},
}

func init() {
	// Url flags
	PostTier.Flags().String("host", "127.0.0.1", "Set a custom host")
	PostTier.Flags().String("port", "8080", "Set a custom port")

	// Tier flags
	PostTier.Flags().Int64P("tier", "i", 0, "Set tier")
	PostTier.Flags().StringP("title", "t", "", "Set the tier's title")
	PostTier.Flags().Int64P("tindex", "T", 0, "Set the tier index")
	PostTier.Flags().StringP("team", "a", "", "Set the tier's team")
}

func postTier(payload *models.CreateTierParams, cf *tierFlags, host, port string) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	// Tier
	for {
		if cf.tier {
			break
		}

		fmt.Println("Please enter tier number:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		tierNum, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			fmt.Println(err)
			continue
		}

		payload.Tier = int64(tierNum)
		break
	}

	// Title
	for {
		if cf.title {
			break
		}

		fmt.Println("Please enter the tier's title:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.Title = utils.StringtoNullString(string(scanner.Bytes()))
		break
	}

	// TIndex
	for {
		if cf.tindex {
			break
		}

		fmt.Println("Please enter the tier index:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		tindex, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			fmt.Println(err)
			continue
		}

		payload.TIndex = utils.Int64toNullInt64(int64(tindex))
		break
	}

	// Team
	for {
		if cf.team {
			break
		}

		fmt.Println("Please enter the tier's team:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.Team = utils.StringtoNullString(string(scanner.Bytes()))
		break
	}

	// confirmation
	for {
		fmt.Println("Is your tier data correct? If not, type n or no.")
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
	path := "v1/board/tiers/"

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
