package dev

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
)

func init() {
	engine := Crawler.Register("1079", "孟加拉国企业研究所", "https://bei-bd.org")

	engine.SetStartingUrls([]string{
		"https://bei-bd.org/project-list/bei-outreach-program",
		"https://bei-bd.org/project-list/current-projects",
		"https://bei-bd.org/project-list/completed-projects",
		"https://bei-bd.org/grid/news",
		"https://bei-bd.org/grid/publications",
	})

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

	engine.OnHTML("h2 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		engine.Visit(element.Attr("href"), Crawler.News)
	})

	engine.OnHTML(".page-title-content > h2", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	engine.OnHTML(".blog-tb-details > .blog-social:first_child > p > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.PageType = Crawler.Report
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})
}
