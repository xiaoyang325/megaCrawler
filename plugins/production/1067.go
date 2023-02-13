package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1067", "布伦南司法中心", "https://www.brennancenter.org")

	engine.SetStartingURLs([]string{"https://www.brennancenter.org/sitemap.xml"})

	extractorConfig := extractors.Config{
		Author:       true,
		Image:        true,
		Language:     true,
		PublishDate:  true,
		Tags:         true,
		Text:         true,
		Title:        true,
		TextLanguage: "",
	}

	extractorConfig.Apply(engine)

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "sitemap.xml") {
			engine.Visit(element.Text, crawlers.Index)
		}
		if strings.Contains(element.Text, "/about/staff/") {
			engine.Visit(element.Text, crawlers.Expert)
		}
		if strings.Contains(element.Text, "/our-work/") {
			engine.Visit(element.Text, crawlers.News)
		}
	})

	engine.OnHTML(".page-info-header__author-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML(".page-bio-header__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		}
	})

	engine.OnHTML(".page-bio-header__role", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Title = element.Text
		}
	})
}
