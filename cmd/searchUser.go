package cmd

import (
	"fmt"

	"AgendaCLI/entity/User"

	"github.com/spf13/cobra"
)

// searchUserCmd represents the searchUser command
var searchUserCmd = &cobra.Command{
	Use:   "searchUser",
	Short: "search all users",
	Long: `you don't need to input anything.For example,
	./agenda searchUser`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("searchUser called")
		User.Search_all_user()
	},
}

func init() {
	RootCmd.AddCommand(searchUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
