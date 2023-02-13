package extractors

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func getTags(dom *colly.HTMLElement) (tags []string) {
	tags = append(tags, dom.ChildTexts("a[rel=\"tag\"], li[itemprop=\"keywords\"] > a")...)
	if len(tags) != 0 {
		tags = append(tags, dom.ChildTexts("a[href*='/tag/'], a[href*='/tags/'], a[href*='/topic/'], a[href*='?keyword='], .entry-category > a")...)
	}
	return crawlers.Unique(tags)
}

func Tags(ctx *crawlers.Context, dom *colly.HTMLElement) {
	ctx.Tags = getTags(dom)
}
