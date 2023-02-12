package dev

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
)

func init() {
	engine := Crawler.Register("1087", "Associated Press-NORC Center for Public Affairs Research", "https://apnorc.org")

	engine.SetStartingUrls([]string{"https://apnorc.org/post-sitemap.xml"})

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
		engine.Visit(element.Text, Crawler.News)
	})
}
