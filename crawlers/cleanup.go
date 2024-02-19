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
)

func TimeCleanup(timeStr string) time.Time {
	timeStr = replacer.Replace(timeStr)
	timeStr = strings.TrimSpace(timeStr)
	parseAny, err := dateparse.ParseAny(timeStr)
	if err == nil {
		return parseAny
	}
	return time.Time{}
}
