package crawlers

import (
	"github.com/araddon/dateparse"
	"strconv"
	"strings"
	"time"
)

var replacer = strings.NewReplacer(
	// Italian
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

	// Spanish
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

	// German
	"januar", "january",
	"februar", "february",
	"märz", "march",
	"april", "april",
	"mai", "may",
	"juni", "june",
	"juli", "july",
	"august", "august",
	"september", "september",
	"oktober", "october",
	"november", "november",
	"dezember", "december",

	"@", "",
	"Reuters/", "",
	"Entertainment Desk", "",
	"REUTERS", "",
	"Last Updated:", "",
	"Posted on", "",
	"Posted at:", "",
	"Published:", "",
	"Updated:", "",
	"From", "",
	"am", "AM",
	"pm", "PM",
	"Sept.", "Sep",
	",", " ",
	"ET", "",
	"IST", "",
	" at", " ",
	"at ", " ",
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
	"15:04 2 Jan 2006",
	"Monday January 2 2006",
	"Monday 2 Jan 2006",
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

func ParseRelativeTime(element string) (time.Time, bool) {
	split := strings.Split(strings.TrimSpace(element), " ")
	if len(split) < 2 {
		return time.Time{}, true
	}
	now := time.Now()
	unit := split[1]
	number, err := strconv.Atoi(split[0])
	if err != nil {
		return time.Time{}, true
	}
	switch unit {
	case "second", "seconds":
		now = now.Add(time.Duration(-number) * time.Second)
	case "minute", "minutes":
		now = now.Add(time.Duration(-number) * time.Minute)
	case "hour", "hours":
		now = now.Add(time.Duration(-number) * time.Hour)
	case "day", "days":
		now = now.AddDate(0, 0, -number)
	case "week", "weeks":
		now = now.AddDate(0, 0, -number*7)
	case "month", "months":
		now = now.AddDate(0, -number, 0)
	case "year", "years":
		now = now.AddDate(-number, 0, 0)
	default:
		return time.Time{}, true
	}
	return now, false
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
