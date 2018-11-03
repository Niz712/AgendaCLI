package cmd

import (
	"fmt"

	"AgendaCLI/entity/Meeting"

	"github.com/spf13/cobra"
)

// operateParticipantCmd represents the operateParticipant command
var operateParticipantCmd =  & cobra.Command {
	Use:"operateParticipant", 
	Short:"the sponsor of the meeting can add or delete the participant in the meeting", 
	Long:`you need to input three arguments(title(t), operation(o), participants(p)).For example:
	addCMD:AgendaCLI operateParticipant - t = Work - o = add - p = bbb; 
	deleteCMD:AgendaCLI operateParticipant - t = Work - o = del - p = aaa`, 
	Run:func(cmd * cobra.Command, args []string) {
		fmt.Println("operateParticipant called")
		title, _: = cmd.Flags().GetString("title")
		operation, _: = cmd.Flags().GetString("op")
		participants, _: = cmd.Flags().GetStringSlice("participants")
		Meeting.Operate_participants(title, operation, participants)
	}, 
}

func init() {
	RootCmd.AddCommand(operateParticipantCmd)
	operateParticipantCmd.Flags().StringP("title", "t", "", "title")
	operateParticipantCmd.Flags().StringP("op", "o", "", "Operation To Participant")
	operateParticipantCmd.Flags().StringSliceP("participants", "p", make([]string, 0), "Names of Participants")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// operateParticipantCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// operateParticipantCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
