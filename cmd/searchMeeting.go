package cmd

import (
	"fmt"

	"AgendaCLI/entity/Meeting"

	"github.com/spf13/cobra"
)

// searchMeetingCmd represents the searchMeeting command
var searchMeetingCmd =  & cobra.Command {
	Use:"searchMeeting", 
	Short:"search the meeting you create or participate in", 
	Long:`you should input two arguments of the command.For example:
	AgendaCLI searchMeeting - s = 2013-10-10/10:00 - e = 2017-10-10/10:00`, 
	Run:func(cmd * cobra.Command, args []string) {
		fmt.Println("searchMeeting called")
		startTime, _: = cmd.Flags().GetString("startTime")
		endTime, _: = cmd.Flags().GetString("endTime")
		Meeting.Search_meeting(startTime, endTime)
	}, 
}

func init() {
	RootCmd.AddCommand(searchMeetingCmd)
	searchMeetingCmd.Flags().StringP("statTime", "s", "", "start time of meeting")
	searchMeetingCmd.Flags().StringP("endTime", "e", "", "end time of meeting")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
