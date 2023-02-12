package Extractors

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func getTags(dom *colly.HTMLElement) (tags []string) {
	tags = append(tags, dom.ChildTexts("a[rel=\"tag\"], li[itemprop=\"keywords\"] > a")...)
	if len(tags) != 0 {
		tags = append(tags, dom.ChildTexts("a[href*='/tag/'], a[href*='/tags/'], a[href*='/topic/'], a[href*='?keyword='], .entry-category > a")...)
	}
	return Crawler.Unique(tags)
}

func Tags(ctx *Crawler.Context, dom *colly.HTMLElement) {
	ctx.Tags = getTags(dom)
}
