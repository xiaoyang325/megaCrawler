package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"strings"
)

func init() {
	engine := Crawler.Register("1067", "布伦南司法中心", "https://www.brennancenter.org")

	engine.SetStartingUrls([]string{"https://www.brennancenter.org/sitemap.xml"})

	extractorConfig := Extractors.Config{
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

	engine.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "sitemap.xml") {
			engine.Visit(element.Text, Crawler.Index)
		}
		if strings.Contains(element.Text, "/about/staff/") {
			engine.Visit(element.Text, Crawler.Expert)
		}
		if strings.Contains(element.Text, "/our-work/") {
			engine.Visit(element.Text, Crawler.News)
		}
	})

	engine.OnHTML(".page-info-header__author-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	engine.OnHTML(".page-bio-header__title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	engine.OnHTML(".page-bio-header__role", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Title = element.Text
		}
	})
}
