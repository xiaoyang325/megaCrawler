package Extractors

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"regexp"
	"time"
)

func pareDateStr(date string) (time.Time, error) {
	parseAny, err := dateparse.ParseAny(date)
	if err != nil {
		return time.Time{}, fmt.Errorf("cannot parse time: %s", err.Error())
	}
	return parseAny, nil
}

func getPublishingDate(dom *colly.HTMLElement) string {
	var strictDateRegex, err = regexp.Compile("\\d+/\\d+/\\d+")
	if err != nil {
		Crawler.Sugar.Panic("Compile regex failed", err)
		return ""
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
		if datetimeString == "" {
			continue
		}
		if obj, err := pareDateStr(datetimeString); err == nil {
			return obj.Format(time.RFC3339)
		} else {
			Crawler.Sugar.Info(err)
		}
	}

	if dateMatch := strictDateRegex.FindString(dom.Request.URL.String()); dateMatch != "" {
		if obj, err := pareDateStr(dateMatch); err == nil {
			return obj.Format(time.RFC3339)
		}
	}
	return ""
}

func PublishingDate(ctx *Crawler.Context, dom *colly.HTMLElement) {
	ctx.PublicationTime = getPublishingDate(dom)
}
