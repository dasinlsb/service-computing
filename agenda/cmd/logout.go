package cmd

import (
	"agenda/entity"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "log out the current account if you have logged in",
	Long: "log out the current account if you have logged in",
	Run: func(cmd *cobra.Command, args []string) {
		if err := entity.Logout(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("logout successfully")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
