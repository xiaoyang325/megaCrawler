package commandImpl

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"time"
)

func List() {
	websites, err := GetWebsites()
	if err != nil {
		println("Service not launched." + err.Error())
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Num", "ID", "Next Iter", "Running?"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	for i, website := range websites {
		if website.IsRunning {
			table.Rich(
				[]string{strconv.Itoa(i), website.Id, website.NextIter.Format(time.RFC3339), strconv.FormatBool(website.IsRunning)},
				[]tablewriter.Colors{
					{tablewriter.Bold, tablewriter.FgGreenColor},
					{tablewriter.Bold, tablewriter.FgGreenColor},
					{tablewriter.Bold, tablewriter.FgGreenColor},
					{tablewriter.Bold, tablewriter.FgGreenColor},
				},
			)
		} else {
			table.Append([]string{strconv.Itoa(i), website.Id, website.NextIter.Format(time.RFC3339), strconv.FormatBool(website.IsRunning)})
		}
	}
	table.Render()
}
