package common

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func TableShow(keys []string, values [][]string, options *TLOptions) {
	if !options.IsApiMode {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAlignment(tablewriter.ALIGN_CENTER)
		table.SetHeader(keys)
		table.AppendBulk(values)
		table.Render()
	}
}
