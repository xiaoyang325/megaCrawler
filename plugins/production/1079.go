package production

import (
	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	engine := crawlers.Register("1079", "孟加拉国企业研究所", "https://bei-bd.org")

	engine.SetStartingUrls([]string{
		"https://bei-bd.org/project-list/bei-outreach-program",
		"https://bei-bd.org/project-list/current-projects",
		"https://bei-bd.org/project-list/completed-projects",
		"https://bei-bd.org/grid/news",
		"https://bei-bd.org/grid/publications",
	})

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

	engine.OnHTML("h2 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		engine.Visit(element.Attr("href"), crawlers.News)
	})

	engine.OnHTML(".page-title-content > h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	engine.OnHTML(".blog-tb-details > .blog-social:first_child > p > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			ctx.PageType = crawlers.Report
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})
}
