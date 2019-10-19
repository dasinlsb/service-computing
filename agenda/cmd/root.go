package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "os"
)

var rootCmd = &cobra.Command{
  Use:   "Agenda",
  Short: "a program manages users' information",
  Long: "a program manages users' information",
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
}
