package crawlers

import (
	"github.com/araddon/dateparse"
	"strings"
	"time"
)

var replacer = strings.NewReplacer(
	"@", "",
	"Reuters/", "",
	"Entertainment Desk", "",
	"REUTERS", "",
	"Posted on", "",
	"Published:", "",
	"am", "AM",
	"pm", "PM",
	"Sept.", "Sep",
	",", " ",
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
	"03:04 PM Jan 2, 2006",
	"15:04 02 01 2006",
	"Monday January 2 2006",
	"03:04 PM EDT Mon January 2 2006",
	"Monday January 2 2006 03:04 PM IST",
	"2 01 2006",
	"02 01 2006 03:04 PM",
	"02 01 2006 15:04",
	"2006 01 02 15:04",
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
