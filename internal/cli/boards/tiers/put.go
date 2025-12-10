package tiers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var PutTier = &cobra.Command{
	Use:   "put --id <uuid> [flags]",
	Short: "update an existing tier by id",

	Run: func(cmd *cobra.Command, args []string) {
		// ----- Populate Payload if Flag Data Given -----
		payload := models.UpdateTierParams{}

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		title, _ := cmd.Flags().GetString("")
		tindex, _ := cmd.Flags().GetInt64("tindex")
		team, _ := cmd.Flags().GetString("team")

		payload.Tier, _ = cmd.Flags().GetInt64("tier")
		payload.Title = utils.StringtoNullString(title)
		payload.TIndex = utils.Int64toNullInt64(tindex)
		payload.Team = utils.StringtoNullString(team)

		tier := strconv.Itoa(int(payload.Tier))

		// ----- Check for Flags Used -----
		flags := tierFlags{
			tier:   cmd.Flags().Lookup("tier").Changed,
			title:  cmd.Flags().Lookup("title").Changed,
			tindex: cmd.Flags().Lookup("tindex").Changed,
			team:   cmd.Flags().Lookup("team").Changed,
		}

		putTier(host, port, tier, &payload, flags)
	},
}

func init() {
	// ----- URL Flags -----
	PutTier.Flags().String("host", "127.0.0.1", "Set a custom host")
	PutTier.Flags().String("port", "8080", "Set a custom port")

	// ----- Tier Flags -----
	PutTier.Flags().Int64P("tier", "i", 0, "Set tier")
	PutTier.Flags().StringP("title", "t", "", "Set the tier's title")
	PutTier.Flags().Int64P("tindex", "T", 0, "Set the tier index")
	PutTier.Flags().StringP("team", "a", "", "Set the tier's team")

	PutTier.MarkFlagRequired("tier")
}

func putTier(host, port, id string, payload *models.UpdateTierParams, flags tierFlags) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	if id == "" {
		fmt.Println("Tier required for put! Use --tier")
		return
	}

	// ----- Construct Url -----
	hostPort := fmt.Sprint(host, ":", port)
	path := "v1/board/tiers/" + id

	u := &url.URL{
		Scheme: "http",
		Host:   hostPort,
		Path:   path,
	}

	// ----- Getting Old Tiers -----
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("error retrieving %s: %s\n", id, err)
		return
	}
	if resp == nil {
		fmt.Println("no response received")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return
	}

	var old models.CreateTierParams
	if err := json.Unmarshal(body, &old); err != nil {
		fmt.Println("error unmarshaling previous tier's data:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	// ----- Tier -----
	for {
		if flags.tier {
			break
		}

		change, err := utils.ChangePrompt("identifier", strconv.Itoa(int(old.Tier)), scanner, "tier")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			newTier, err := strconv.Atoi(string(change))
			if err != nil {
				fmt.Println(err)
				continue
			}
			payload.Tier = int64(newTier)
		} else {
			payload.Tier = old.Tier
		}
		break
	}

	// ----- Title -----
	for {
		if flags.title {
			break
		}

		change, err := utils.ChangePrompt("title", old.Title.String, scanner, "tier")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			payload.Title = utils.StringtoNullString(string(change))
		} else {
			payload.Title = old.Title
		}
		break
	}

	// ----- TIndex -----
	for {
		if flags.tindex {
			break
		}

		change, err := utils.ChangePrompt("index", strconv.Itoa(int(old.TIndex.Int64)), scanner, "tier")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			newTIndex, err := strconv.Atoi(string(change))
			if err != nil {
				fmt.Println(err)
				continue
			}
			payload.TIndex = utils.Int64toNullInt64(int64(newTIndex))
		} else {
			payload.TIndex = old.TIndex
		}
		break
	}

	// ----- Team -----
	for {
		if flags.team {
			break
		}

		change, err := utils.ChangePrompt("team", old.Team.String, scanner, "tier")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			payload.Team = utils.StringtoNullString(string(change))
		} else {
			payload.Team = old.Team
		}
		break
	}

	// ----- Confirm -----
	for {
		fmt.Println("Is the tier data correct? (y/n)")
		utils.PrintStruct(payload, false)
		scanner.Scan()
		confirmation := scanner.Bytes()

		ok, err := utils.YesOrNo(confirmation, scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !ok {
			return
		}
		break
	}

	// ----- Marshal Payload -----
	jsonPayload, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	// ----- PUT -----
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, u.String(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Problem with PUT:", err)
		return
	}

	putResp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error with response:", err)
		return
	}
	if putResp == nil {
		fmt.Println("no response received")
		return
	}
	defer putResp.Body.Close()

	fmt.Println("PUT status:", putResp.Status)

	body, err = io.ReadAll(putResp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Println(string(body))
}
