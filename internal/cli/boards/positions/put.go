package positions

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

var PutPosition = &cobra.Command{
	Use:   "put --id <uuid> [flags]",
	Short: "update an existing position by id",

	Run: func(cmd *cobra.Command, args []string) {
		// ----- Populate Payload if Flag Data Given -----
		payload := models.UpdatePositionParams{}

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		payload.Oid, _ = cmd.Flags().GetString("oid")
		payload.Semester, _ = cmd.Flags().GetString("semester")
		payload.Tier, _ = cmd.Flags().GetInt64("tier")

		// ----- Check for Flags Used -----
		changedFlags := positionFlags{
			oid:      cmd.Flags().Lookup("oid").Changed,
			semester: cmd.Flags().Lookup("semester").Changed,
			tier:     cmd.Flags().Lookup("tier").Changed,
		}

		putPosition(host, port, payload.Oid, &payload, changedFlags)
	},
}

func init() {
	// ----- URL Flags -----
	PutPosition.Flags().String("host", "127.0.0.1", "Set a custom host")
	PutPosition.Flags().String("port", "8080", "Set a custom port")

	// ----- Position Flags -----
	PutPosition.Flags().StringP("oid", "o", "", "Set the oid for position")
	PutPosition.Flags().StringP("semester", "s", "", "Set the semester for position")
	PutPosition.Flags().Int64P("tier", "t", 0, "Set the tier for position")

	PutPosition.MarkFlagRequired("oid")
}

func putPosition(host, port, id string, payload *models.UpdatePositionParams, flags positionFlags) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	if id == "" {
		fmt.Println("Oid required for put! Use --oid")
		return
	}

	// ----- Construct url -----
	hostPort := fmt.Sprint(host, ":", port)
	path := "v1/board/positions/" + id

	u := &url.URL{
		Scheme: "http",
		Host:   hostPort,
		Path:   path,
	}

	// ----- Getting old positions -----
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

	var old models.CreatePositionParams
	if err := json.Unmarshal(body, &old); err != nil {
		fmt.Println("error unmarshaling previous postion's data:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	// ----- Oid -----
	for {
		if flags.oid {
			break
		}

		change, err := utils.ChangePrompt("identifier", strconv.Itoa(int(old.Tier)), scanner, "position")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			payload.Oid = string(change)
		} else {
			payload.Oid = old.Oid
		}
		break
	}

	// ----- Semester -----
	for {
		if flags.semester {
			break
		}

		change, err := utils.ChangePrompt("semester", old.Semester, scanner, "position")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			payload.Semester = string(change)
		} else {
			payload.Semester = old.Semester
		}
		break
	}

	// ----- Tier -----
	for {
		if flags.tier {
			break
		}

		change, err := utils.ChangePrompt("tier", strconv.Itoa(int(old.Tier)), scanner, "position")
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
			payload.Tier = (int64(newTier))
		} else {
			payload.Tier = old.Tier
		}
		break
	}
	// ----- Confirm -----
	for {
		fmt.Println("Is the position data correct? (y/n)")
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

	// ----- Marshal payload -----
	jsonPayload, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	// ----- Put -----
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
