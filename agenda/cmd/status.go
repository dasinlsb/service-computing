package cmd

import (
	"agenda/entity"
	"agenda/view"
	"github.com/spf13/cobra"
	"log"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "list status of current account if you have logged in",
	Long: "list status of current account if you have logged in",
	Run: func(cmd *cobra.Command, args []string) {
		info, err := entity.Status()
		if err != nil {
			log.Fatal(err)
		}
		view.PrintTable(info)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
