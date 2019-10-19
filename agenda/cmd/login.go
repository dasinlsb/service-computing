package cmd

import (
	"agenda/entity"
	"agenda/view"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login an account with [Username] & [Password]",
	Long: "login an account with [Username] & [Password]",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalln("username required")
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalln("password required")
		}
		if err := entity.Login(username, password); err != nil {
			log.Fatal(err)
		}
		fmt.Println("login successfully")
		info, err := entity.GetCustom(username)
		if err != nil {
			log.Fatal(err)
		}
		view.PrintTable(info)
	},
}

func init() {
	loginCmd.Flags().StringP("username", "u", "", "Your username")
	loginCmd.Flags().StringP("password", "p", "", "Your password")
	rootCmd.AddCommand(loginCmd)
}
