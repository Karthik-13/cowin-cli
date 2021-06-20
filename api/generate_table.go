package api

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

const (
	ALIGN_DEFAULT = iota
	ALIGN_CENTER
	ALIGN_RIGHT
	ALIGN_LEFT
)

func GenerateTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetNoWhiteSpace(false)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(ALIGN_LEFT)
	table.SetAlignment(ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(true)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	return table
}
