package cmd

import (
	"fmt"

	"AgendaCLI/entity/Meeting"

	"github.com/spf13/cobra"
)

// exitMeetingCmd represents the exitMeeting command
var exitMeetingCmd =  & cobra.Command {
	Use:"exitMeeting", 
	Short:"to exit the meeting as a participant", 
	Long:`you need to input the title of the meeting you want to exit.For example:
	AgendaCLI exitMeeting - t = Work`, 
	Run:func(cmd * cobra.Command, args []string) {
		fmt.Println("exitMeeting called")
		title, _: = cmd.Flags().GetString("title")
		Meeting.Exit_meeting(title)
	}, 
}

func init() {
	RootCmd.AddCommand(exitMeetingCmd)
	exitMeetingCmd.Flags().StringP("title", "t", "", "title")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exitMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exitMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
