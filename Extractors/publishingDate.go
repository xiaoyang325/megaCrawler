package Extractors

import (
	"errors"
	"github.com/araddon/dateparse"
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"regexp"
	"time"
)

var strictDateRegex, _ = regexp.Compile("(?<=\\W)([./\\-_]?(19|20)\\d{2})[./\\-_]?(([0-3]?\\d[./\\-_])|(\\w{3,5}[./\\-_]))([0-3]?\\d[./\\-]?)?")

func pareDateStr(date string) (time.Time, error) {
	parseAny, err := dateparse.ParseAny(date)
	if err != nil {
		return time.Time{}, errors.New("cannot parse time")
	}
	return parseAny, nil
}

func getPublishingDate(dom *colly.HTMLElement) string {
	if dateMatch := strictDateRegex.FindStringSubmatch(dom.Request.URL.String()); len(dateMatch) > 0 {
		if obj, err := pareDateStr(dateMatch[0]); err != nil {
			return obj.Format(time.RFC3339)
		}
	}

	publishDateTags := []selectorContentPair{
		{
			selector: "*[property=\"rnews:datePublished\"]",
			content:  "content",
		},
		{
			selector: "*[property=\"article:published_time\"]",
			content:  "content",
		},
		{
			selector: "*[name=\"OriginalPublicationDate\"]",
			content:  "content",
		},
		{
			selector: "*[itemprop=\"datePublished\"]",
			content:  "datetime",
		},
		{
			selector: "*[property=\"og:published_time\"]",
			content:  "content",
		},
		{
			selector: "*[name=\"article_date_original\"]",
			content:  "content",
		},
		{
			selector: "*[name=\"publication_date\"]",
			content:  "content",
		},
		{
			selector: "*[name=\"sailthru.date\"]",
			content:  "content",
		},
		{
			selector: "*[name=\"PublishDate\"]",
			content:  "content",
		},
		{
			selector: "*[pubdate=\"pubdate\"]",
			content:  "datetime",
		},
		{
			selector: "*[name=\"publish_date\"]",
			content:  "content",
		},
	}
	for _, tag := range publishDateTags {
		datetimeString := dom.ChildAttr(tag.selector, tag.content)
		if obj, err := pareDateStr(datetimeString); err != nil {
			return obj.Format(time.RFC3339)
		}
	}
	return ""
}

func PublishingDate(ctx *Crawler.Context, dom *colly.HTMLElement) {
	ctx.PublicationTime = getPublishingDate(dom)
}
