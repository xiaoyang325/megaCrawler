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
)

var template = []string{
	"03:04 PM Jan 2, 2006",
}

func TimeCleanup(timeStr string) time.Time {
	timeStr = replacer.Replace(timeStr)
	timeStr = strings.TrimSpace(timeStr)
	for _, subStr := range strings.Split(timeStr, "\n") {
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
