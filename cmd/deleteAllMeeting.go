package cmd

import (
	"fmt"

	"AgendaCLI/entity/Meeting"

	"github.com/spf13/cobra"
)

// deleteAllMeetingsCmd represents the deleteAllMeetings command
var deleteAllMeetingsCmd =  & cobra.Command {
	Use:"deleteAllMeeting", 
	Short:"delete all the meetings that user create", 
	Long:`you don't need to input anything.For example:
	AgendaCLI deleteAllMeeting`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteAllMeetings called")
		Meeting.Delete_all_meetings()
	},
}

func init() {
	RootCmd.AddCommand(deleteAllMeetingsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteAllMeetingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteAllMeetingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
