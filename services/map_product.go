package services

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// please add based on defect-dojo scan
var scan_type_list = map[int64]string{
	1: "Semgrep JSON Report",
	2: "Nuclei Scan",
	3: "Trufflehog3 Scan",
	4: "Trufflehog Scan",
}

func scan_type(num int64) (string, bool) {
	value, ok := scan_type_list[num]
	return value, ok
}

func GetList() {
	data := [][]string{}
	for num, value := range scan_type_list {
		data = append(data, []string{
			fmt.Sprintf("%v -scan_type=%v", num, num),
			value,
		})

	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Scan type number", "Scan Type"})
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiYellowColor},
	)
	table.SetBorder(true)  // Set Border to false
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
