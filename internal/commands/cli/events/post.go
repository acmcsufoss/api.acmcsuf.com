package events

import (
	"fmt"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/commands/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
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

		PostEvent(&payload, changedFlags, config.Cfg)
	},
}
