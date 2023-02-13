package commands

import (
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

func List() {
	websites, err := GetWebsites()
	if err != nil {
		println("Service not launched." + err.Error())
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	iter := 0.0
	for i, website := range websites {
		if website.IsRunning {
			table.Rich(
				[]string{strconv.Itoa(i), website.ID, website.NextIter.Format(time.RFC3339), strconv.FormatBool(website.IsRunning)},
				[]tablewriter.Colors{
					{},
					{},
					{},
					{tablewriter.Bold, tablewriter.FgHiGreenColor, tablewriter.UnderlineSingle},
				},
			)
			iter += website.IterPerSec
		} else {
			table.Rich(
				[]string{strconv.Itoa(i), website.ID, website.NextIter.Format(time.RFC3339), strconv.FormatBool(website.IsRunning)},
				[]tablewriter.Colors{
					{},
					{},
					{},
					{tablewriter.Bold, tablewriter.FgHiRedColor, tablewriter.UnderlineSingle},
				},
			)
		}
	}
	table.SetHeader([]string{"Num", "ID", "Next Iter", "Running?"})
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetFooter([]string{"", "", "Total Iter/s", strconv.FormatFloat(iter, 'E', -1, 64)})
	table.SetBorder(false)
	table.SetCenterSeparator(" ")
	table.SetColumnSeparator(" ")
	table.SetHeaderColor(tablewriter.Color(tablewriter.FgHiWhiteColor), tablewriter.Color(tablewriter.FgHiWhiteColor), tablewriter.Color(tablewriter.FgHiWhiteColor), tablewriter.Color(tablewriter.FgHiWhiteColor))
	table.SetColumnColor(tablewriter.Color(tablewriter.FgHiYellowColor), tablewriter.Color(tablewriter.FgHiWhiteColor), tablewriter.Color(tablewriter.FgHiWhiteColor), tablewriter.Color(tablewriter.FgHiWhiteColor))
	table.SetFooterColor(tablewriter.Color(), tablewriter.Color(), tablewriter.Color(tablewriter.BgHiWhiteColor, tablewriter.FgHiBlackColor), tablewriter.Color(tablewriter.BgHiWhiteColor, tablewriter.FgHiBlackColor))
	table.Render()
}
