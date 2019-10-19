package view

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func PrintTable(data [][]string) {
	head := data[0]
	body := data[1:]
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(head)
	for _, r := range body {
		table.Append(r)
	}
	table.Render()
}