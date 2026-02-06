package events

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	// TODO: db params shouldn't be exposed here
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/charmbracelet/huh"

	"github.com/acmcsufoss/api.acmcsuf.com/utils/requests"
	"github.com/spf13/cobra"
)

var PostEvent = &cobra.Command{
	Use:   "post",
	Short: "Post a new event.",

	Run: func(cmd *cobra.Command, args []string) {
		payload := models.CreateEventParams{}
		// err := huh.NewForm().Run()
		// if err != nil {
		// 	if err == huh.ErrUserAborted {
		// 		fmt.Println("User canceled the form — exiting.")
		// 	}
		// 	fmt.Println("Uh oh:", err)
		// 	os.Exit(1)
		//}

		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		payload.Location, _ = cmd.Flags().GetString("location")
		startAtString, _ := cmd.Flags().GetString("startat")
		duration, _ := cmd.Flags().GetString("duration")
		payload.IsAllDay, _ = cmd.Flags().GetBool("isallday")
		payload.Host, _ = cmd.Flags().GetString("host")

		if startAtString != "" {
			var err error
			payload.StartAt, err = utils.ByteSlicetoUnix([]byte(startAtString))
			if err != nil {
				fmt.Println(err)
				return
			}
			if duration != "" {
				var err error
				payload.EndAt, err = utils.TimeAfterDuration(payload.StartAt, duration)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}

		if duration != "" && startAtString == "" {
			fmt.Printf("--startat is required in order to use --duration")
		}

		changedFlags := eventFlags{
			uuid:     cmd.Flags().Lookup("uuid").Changed,
			location: cmd.Flags().Lookup("location").Changed,
			startat:  cmd.Flags().Lookup("startat").Changed,
			duration: cmd.Flags().Lookup("duration").Changed,
			isallday: cmd.Flags().Lookup("isallday").Changed,
			host:     cmd.Flags().Lookup("host").Changed,
		}

		postEvent(&payload, changedFlags, config.Cfg)
	},
}

func init() {
	PostEvent.Flags().StringP("uuid", "u", "", "Set uuid of new event")
	PostEvent.Flags().StringP("location", "l", "", "Set location of new event")
	PostEvent.Flags().StringP("startat", "s", "", "Set the start time of new event (Format: 03:04:05PM 01/02/06)")
	PostEvent.Flags().StringP("duration", "d", "", "Set the duration of new event (Format: 03:04:05)")
	PostEvent.Flags().StringP("host", "H", "", "Set host of new event")
	PostEvent.Flags().BoolP("isallday", "a", false, "Set if new event is all day")
}

func postEvent(payload *models.CreateEventParams, changedFlag eventFlags, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	// ----- Uuid -----
	for {
		if changedFlag.uuid {
			break
		}

		var uuid string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Event Post:").
			Description("Please enter event's uuid:").
			Prompt("> ").
			Value(&uuid).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(uuid))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		uuidBuffer := scanner.Bytes()
		payload.Uuid = string(uuidBuffer)
		break
	}

	// ----- Location -----
	for {
		if changedFlag.location {
			break
		}

		var location string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Event Post:").
			Description("Please enter the event's location:").
			Prompt("> ").
			Value(&location).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(location))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		locationBuffer := scanner.Bytes()
		payload.Location = string(locationBuffer)
		break
	}

	// ----- Start Time -----
	for {

		if changedFlag.startat {
			break
		}

		var timeStart string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Event Post:").
			Description("Please enter the start time of the event in the following format:\n [Month]/[Day]/[Year] [Hour]:[Minute][PM | AM]\nFor example: \x1b[93m01/02/06 03:04PM\x1b[0m").
			Prompt("> ").
			Value(&timeStart).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(timeStart))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading start time:", err)
			continue
		}
		startTimeBuffer := scanner.Bytes()
		startTime, err := utils.ByteSlicetoUnix(startTimeBuffer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		payload.StartAt = startTime
		break
	}

	// ----- End Time (Duration) -----
	for {

		if changedFlag.duration {
			break
		}

		var duration string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Event Post:").
			Description("Please enter the duration of the event in the following format:\n [Hour]:[Minute]\nFor example: \x1b[93m03:04\x1b[0m").
			Prompt("> ").
			Value(&duration).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(duration))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading end time:", err)
			continue
		}

		endTimeBuffer := scanner.Bytes()
		endTime, err := utils.TimeAfterDuration(payload.StartAt, string(endTimeBuffer))
		if err != nil {
			fmt.Println(err)
			continue
		}

		payload.EndAt = endTime
		break
	}

	// ----- Is all day -----

	for {
		if changedFlag.isallday {
			break
		}

		var allDayYes string
		err := huh.NewSelect[string]().
			Title("ACMCSUF-CLI Event Post:").
			Description("Is your event all day?").
			Options(
				huh.NewOption("Yes", "yes"),
				huh.NewOption("No", "n"),
			).
			Value(&allDayYes).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(allDayYes))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		isAllDayBuffer := scanner.Bytes()

		isAllDay, err := utils.YesOrNo(isAllDayBuffer, scanner)
		if err != nil {
			fmt.Println(err)
		}
		payload.IsAllDay = isAllDay
		break
	}

	// ----- Host -----
	for {
		if changedFlag.host {
			break
		}

		var host string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Event Post:").
			Description("Please enter the event host").
			Prompt("> ").
			Value(&host).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(host))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		hostBuffer := scanner.Bytes()
		payload.Host = string(hostBuffer)
		break
	}

	// ----- Confirmation -----
	for {
		var option string
		description := "Is your event data correct?\n" + utils.PrintStruct(payload)
		err := huh.NewSelect[string]().
			Title("ACMCSUF-CLI Event Post:").
			Description(description).
			Options(
				huh.NewOption("Yes", "yes"),
				huh.NewOption("No", "n"),
			).
			Value(&option).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(option))
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

	// ----- Convert to Json -----
	jsonEvent, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	postURL := baseURL.JoinPath("v1/events")

	// ----- Post -----
	request, err := requests.NewRequestWithAuth(http.MethodPost, postURL.String(),
		strings.NewReader(string(jsonEvent)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating post request: %v", err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Failed to post event:", err)
		return
	}
	defer response.Body.Close()

	// ----- Read Response Info -----
	if response.StatusCode != http.StatusOK {
		fmt.Println("Response status", response.Status)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}
	utils.PrettyPrintJSON(body)
}
