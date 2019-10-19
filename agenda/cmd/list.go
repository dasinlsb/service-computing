package cmd

import (
	"agenda/entity"
	"agenda/view"
	"github.com/spf13/cobra"
	"log"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all users if you have logged in",
	Long: "list all users if you have logged in",
	Run: func(cmd *cobra.Command, args []string) {
		info, err := entity.List()
		if err != nil {
			log.Fatal(err)
		}
		view.PrintTable(info)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
