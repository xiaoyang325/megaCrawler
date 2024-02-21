package crawlers

import (
	"github.com/araddon/dateparse"
	"strings"
	"time"
)

var replacer = strings.NewReplacer(
	"gennaio", "january",
	"febbraio", "february",
	"marzo", "march",
	"aprile", "april",
	"maggio", "may",
	"giugno", "june",
	"luglio", "july",
	"agosto", "august",
	"settembre", "september",
	"ottobre", "october",
	"novembre", "november",
	"dicembre", "december",

	"enero", "january",
	"febrero", "february",
	"marzo", "march",
	"abril", "april",
	"mayo", "may",
	"junio", "june",
	"julio", "july",
	"agosto", "august",
	"septiembre", "september",
	"octubre", "october",
	"noviembre", "november",
	"diciembre", "december",

	"@", "",
	"Reuters/", "",
	"Entertainment Desk", "",
	"REUTERS", "",
	"Last Updated:", "",
	"Posted on", "",
	"Posted at:", "",
	"Published:", "",
	"Updated:", "",
	"am", "AM",
	"pm", "PM",
	"Sept.", "Sep",
	",", " ",
	"ET", "",
	"IST", "",
	"at", " ",
	"— updated on", "\n",
	"|", " ",
	"—", " ",
	".", " ",
	" de ", " ",
	"/", " ",
	"(", "",
	")", "",
)

var template = []string{
	"03:04 PM Jan 2 2006",
	"15:04 02 01 2006",
	"Monday January 2 2006",
	"Monday 2 January 2006",
	"03:04 PM EDT Mon January 2 2006",
	"Monday January 2 2006 03:04 PM",
	"Monday January 2 2006",
	"1 2 2006 3:4 PM",
	"1 2 2006",
	"2 1 2006",
	"2006 1 2",
	"2 1 2006 3:4 PM",
	"2 1 2006 15:4",
	"2006 1 2 15:4",
	"2 January 2006 15:4",
}

func TimeCleanup(timeStr string) time.Time {
	timeStr = replacer.Replace(timeStr)
	timeStr = strings.TrimSpace(timeStr)
	for _, subStr := range strings.Split(timeStr, "\n") {
		subStr = StandardizeSpaces(strings.TrimSpace(subStr))
		parse, err := dateparse.ParseAny(subStr)
		if err == nil {
			return parse
		}
		for _, temp := range template {
			parse, err = time.Parse(temp, subStr)
			if err == nil {
				return parse
			}
		}
	}
	return time.Time{}
}
