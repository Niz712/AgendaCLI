package cmd

import (
	"fmt"

	"AgendaCLI/entity/User"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register an account for the your agenda",
	Long: `you need to input user message to create an account, the arguments include usernam, password, email, telephone.For example:
	AgendaCLI register -u=abcde -p=12345678 -e=1234567890@qq.com -t=15012345678`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("telephone")
		//	fmt.Println("register called by " + username + password + email)
		userInfo := &User.User{username, password, email, phone, make([]string, 0, 5), make([]string, 0, 5)}
		User.Register_user(userInfo)
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "Username")
	registerCmd.Flags().StringP("password", "p", "", "User password")
	registerCmd.Flags().StringP("email", "e", "", "User email")
	registerCmd.Flags().StringP("telephone", "t", "", "User telephone")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
