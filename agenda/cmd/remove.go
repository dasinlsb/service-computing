package cmd

import (
	"agenda/entity"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove current account if you have logged in",
	Long: "remove current account if you have logged in",
	Run: func(cmd *cobra.Command, args []string) {
		if err := entity.Remove(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("remove successfully")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
