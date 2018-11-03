package cmd

import (
	"fmt"

	"AgendaCLI/entity/Meeting"

	"github.com/spf13/cobra"
	//"time"
)

// createMeetingCmd represents the createMeeting command
var createMeetingCmd =  & cobra.Command {
	Use:"createMeeting", 
	Short:"to create a meeting for user", 
	Long:`the user need to input the title, participants, startTime, endTime of the meeting. For example:
	AgendaCLI createMeeting - t = Work - p = abcde - s = 2016-10-10/10:00 - e = 2017-10-10/10:00`, 
	Run:func(cmd * cobra.Command, args []string) {
		fmt.Println("createMeeting called")
		title, _: = cmd.Flags().GetString("title")
		participants, _: = cmd.Flags().GetStringSlice("participants")
		startTime, _: = cmd.Flags().GetString("startTime")
		endTime, _: = cmd.Flags().GetString("endTime")
		timeS, _: = Meeting.String_to_date(startTime)
		timeE, _: = Meeting.String_to_date(endTime)
		Meeting.Create_meeting( & Meeting.Meeting {title, "", participants, timeS, timeE, ""})
	}, 
}

func init() {
	RootCmd.AddCommand(createMeetingCmd)
	createMeetingCmd.Flags().StringP("title", "t", "", "title")
	createMeetingCmd.Flags().StringSliceP("participants", "p", make([]string, 0), "participants")
	createMeetingCmd.Flags().StringP("startTime", "s", "", "startTime")
	createMeetingCmd.Flags().StringP("endTime", "e", "", "User endTime")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
