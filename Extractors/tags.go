package Extractors

import "github.com/gocolly/colly/v2"

func getTags(dom *colly.HTMLElement) (tags []string) {
	tags = append(tags, dom.ChildTexts("a[rel=\"tag\"]")...)
	if len(tags) != 0 {
		tags = append(tags, dom.ChildTexts("a[href*='/tag/'], a[href*='/tags/'], a[href*='/topic/'], a[href*='?keyword=']")...)
	}
	return
}
