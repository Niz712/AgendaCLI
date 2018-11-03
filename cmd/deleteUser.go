package cmd

import (
	"fmt"

	"AgendaCLI/entity/Meeting"

	"github.com/spf13/cobra"
)

// deleteUserCmd represents the deleteUser command
var deleteUserCmd = &cobra.Command{
	Use:   "deleteUser",
	Short: "to delete the user own account",
	Long: `you don't need to input anything.For example:
	AgendaCLI deleteUser`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteUser called")
		Meeting.Delete_user()
	},
}

func init() {
	RootCmd.AddCommand(deleteUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
