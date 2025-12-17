package events

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

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/requests"
)

var PutEvents = &cobra.Command{
	Use:   "put",
	Short: "Used to update an event",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")

		payload := models.CreateEventParams{}
		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		payload.Location, _ = cmd.Flags().GetString("location")
		startAtString, _ := cmd.Flags().GetString("startat")
		durationString, _ := cmd.Flags().GetString("duration")
		payload.IsAllDay, _ = cmd.Flags().GetBool("allday")
		payload.Host, _ = cmd.Flags().GetString("host")

		if startAtString != "" {
			var err error
			payload.StartAt, err = utils.ByteSlicetoUnix([]byte(startAtString))
			if err != nil {
				fmt.Println(err)
				return
			}
			if durationString != "" {
				var err error
				payload.EndAt, err = utils.TimeAfterDuration(payload.StartAt, durationString)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}

		changedFlags := eventFlags{
			uuid:     cmd.Flags().Lookup("uuid").Changed,
			location: cmd.Flags().Lookup("location").Changed,
			startat:  cmd.Flags().Lookup("startat").Changed,
			duration: cmd.Flags().Lookup("duration").Changed,
			isallday: cmd.Flags().Lookup("isallday").Changed,
			host:     cmd.Flags().Lookup("host").Changed,
		}

		updateEvent(id, &payload, changedFlags, config.Cfg)
	},
}

func init() {
	// URL Flags
	PutEvents.Flags().String("id", "", "Event to update")
	PutEvents.Flags().String("urlhost", "127.0.0.1", "Custom host")
	PutEvents.Flags().String("port", "8080", "Custom port")

	// Payload flags
	PutEvents.Flags().StringP("uuid", "u", "", "Set uuid of new event")
	PutEvents.Flags().StringP("location", "l", "", "Set location of new event")
	PutEvents.Flags().StringP("startat", "s", "", "Set the start time of new event (Format: 03:04:05PM 01/02/06)")
	PutEvents.Flags().StringP("duration", "d", "", "Set the end time of new event (Format: 03:04:05)")
	PutEvents.Flags().StringP("host", "H", "", "Set host of new event")
	PutEvents.Flags().BoolP("isallday", "a", false, "Set if new event is all day")

	// This flag is neccessary
	PutEvents.MarkFlagRequired("id")
}

func updateEvent(id string, payload *models.CreateEventParams, changedFlags eventFlags, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	// ----- Retrieve payload -----
	retrievalURL := baseURL.JoinPath(fmt.Sprintf("v1/events/", id))
	getResponse, err := http.Get(retrievalURL.String())
	if err != nil {
		fmt.Printf("Error retrieving %s: %s", id, err)
		return
	}

	defer getResponse.Body.Close()
	body, err := io.ReadAll(getResponse.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	if getResponse.StatusCode != http.StatusOK {
		fmt.Println("Response status:", getResponse.Status)
		return
	}

	var oldpayload models.CreateEventParams
	err = json.Unmarshal(body, &oldpayload)
	if err != nil {
		fmt.Println("error unmarshalling previous event data:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	// ----- Change the event -----
	// Note: We want to PUT the payload, not old payload
	// payload values are empty if user did not input a value in the command line

	// ----- uuid -----
	for {
		if payload.Uuid == "" {
			changeTheEventUuid, err := utils.ChangePrompt("uuid", oldpayload.Uuid, scanner, "event")
			if err != nil {
				fmt.Println(err) // Custom errors in changePrompt()
				continue
			}

			if changeTheEventUuid != nil {
				payload.Uuid = string(changeTheEventUuid)
			} else {
				payload.Uuid = oldpayload.Uuid
			}
		}
		break
	}

	// ----- Location -----
	for {
		if changedFlags.location {
			break
		}
		changeTheEventLocation, err := utils.ChangePrompt("location", oldpayload.Location, scanner, "event")
		if err != nil {
			fmt.Println(err)
			continue
		}

		if changeTheEventLocation != nil {
			payload.Location = string(changeTheEventLocation)
		} else {
			payload.Location = oldpayload.Location
		}
		break
	}

	// ----- Start time -----
	for {
		if changedFlags.startat {
			break
		}
		changeTheEventStartAt, err := utils.ChangePrompt("start time (format: 01/02/06 03:04PM)", utils.FormatUnix(oldpayload.StartAt), scanner, "event")
		if err != nil {
			fmt.Println(err)
			continue
		}

		if changeTheEventStartAt != nil {
			payload.StartAt, err = utils.ByteSlicetoUnix(changeTheEventStartAt)
			if err != nil {
				fmt.Println("Error with reading start integer:", err)
				continue
			}
		} else {
			payload.StartAt = oldpayload.StartAt
		}
		break
	}

	// ----- End time (Duration) -----
	for {
		if changedFlags.duration {
			break
		}
		changeTheEventEndAt, err := utils.ChangePrompt("end time (format: 01/02/06 03:04 )", utils.FormatUnix(oldpayload.EndAt), scanner, "event")
		if err != nil {
			fmt.Println(err)
			continue
		}

		if changeTheEventEndAt != nil {
			payload.EndAt, err = utils.ByteSlicetoUnix(changeTheEventEndAt)
			if err != nil {
				fmt.Println("Error with reading end integer:", err)
				continue
			}
		} else {
			payload.EndAt = oldpayload.EndAt
		}
		break
	}

	// ----- All day -----
	// This is kind of awkward but I don't know have a workaround at the moment
	for {
		if changedFlags.isallday {
			break
		}

		changeTheEventAllDay, err := utils.ChangePrompt("all day status", strconv.FormatBool(oldpayload.IsAllDay), scanner, "event")
		if err != nil {
			fmt.Println(err)
			continue
		}

		if changeTheEventAllDay != nil {
			newAllDayBuffer := scanner.Bytes()
			payload.IsAllDay, err = utils.YesOrNo(newAllDayBuffer, scanner)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			payload.IsAllDay = oldpayload.IsAllDay
		}
		break
	}

	// ----- Host -----
	for {
		if changedFlags.host {
			break
		}
		changeTheEventHost, err := utils.ChangePrompt("host", oldpayload.Host, scanner, "event")
		if err != nil {
			fmt.Println(err)
			continue
		}

		if changeTheEventHost != nil {
			payload.Host = string(changeTheEventHost)
		} else {
			payload.Host = oldpayload.Host
		}
		break
	}

	// ----- PUT the payload -----

	updatePayload := models.UpdateEventParams{
		Uuid:     payload.Uuid,
		Location: utils.StringtoNullString(payload.Location),
		StartAt:  utils.Int64toNullInt64(payload.StartAt),
		EndAt:    utils.Int64toNullInt64(payload.EndAt),
		IsAllDay: utils.BooltoNullBool(payload.IsAllDay),
		Host:     utils.StringtoNullString(payload.Host),
	}

	// Confirmation
	// TODO: Fix put
	for {
		fmt.Println("Are these changes okay?[y/n]")
		utils.PrintStruct(updatePayload)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error scanning confirmation:", err)
			continue
		}

		confirmationBuffer := scanner.Bytes()
		confirmation, err := utils.YesOrNo(confirmationBuffer, scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if !confirmation {
			return
		}

		break
	}

	// ----- Put the Payload -----
	newPayload, err := json.Marshal(updatePayload)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	client := &http.Client{}

	request, err := requests.NewRequestWithAuth(http.MethodPut, retrievalURL.String(), bytes.NewBuffer(newPayload))
	if err != nil {
		fmt.Println("Problem with PUT:", err)
		return
	}

	putResponse, err := client.Do(request)
	if err != nil {
		fmt.Println("Error with response:", err)
		return
	}

	if putResponse.StatusCode != http.StatusOK {
		fmt.Println("Response status:", putResponse.Status)
	}
	defer putResponse.Body.Close()
	body, err = io.ReadAll(putResponse.Body)
	if err != nil {
		fmt.Println("Error with body:", err)
		return
	}
	utils.PrettyPrintJSON(body)
}
