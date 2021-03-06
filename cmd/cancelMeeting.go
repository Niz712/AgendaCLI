package cmd

import (
	"fmt"

	"AgendaCLI/entity/Meeting"

	"github.com/spf13/cobra"
)

// cancelMeetingCmd represents the cancelMeeting command
var cancelMeetingCmd =  & cobra.Command {
	Use:"cancelMeeting", 
	Short:"to delete the meeting that user creates", 
	Long:`user should input the title of the meeting. For example:
	AgendaCLI cancelMeeting - t = Work`, 
	Run:func(cmd * cobra.Command, args []string) {
		fmt.Println("cancelMeeting called")
		title, _: = cmd.Flags().GetString("title")
		Meeting.Cancel_meeting(title)
	}, 
}

func init() {
	RootCmd.AddCommand(cancelMeetingCmd)
	cancelMeetingCmd.Flags().StringP("title", "t", "", "title")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
