package cmd

import (
	"agenda/entity"
	"agenda/model/user"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)


var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register a new account with [Username] [Password] [Email] [Telephone]",
	Long: "register a new account with [Username] [Password] [Email] [Telephone]",
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalln("username required")
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalln("password required")
		}
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			log.Fatalln("email required")
		}
		phone, err := cmd.Flags().GetString("telephone")
		if err != nil {
			log.Fatalln("telephone required")
		}
		u := user.User{username, password, email, phone}
		if err := entity.Register(u); err != nil {
			log.Fatal(err)
		}
		fmt.Println("register successfully")
		//_ = entity.Login(username, password)
		//info, err := entity.GetCustom(username)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//view.PrintTable(info)
	},
}

func init() {
	registerCmd.Flags().StringP("username", "u", "", "Your username")
	registerCmd.Flags().StringP("password", "p", "", "Your password")
	registerCmd.Flags().StringP("email", "e", "", "Your email address")
	registerCmd.Flags().StringP("telephone", "t", "", "Your telephone number")
	rootCmd.AddCommand(registerCmd)
}
