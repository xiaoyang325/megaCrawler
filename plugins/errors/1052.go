package errors

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"strings"
)

func init() {
	engine := Crawler.Register("1052", "环境管理与评估研究所", "https://www.iema.net")

	engine.SetStartingUrls([]string{"https://www.iema.net/sitemap.xml"})

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
		if strings.Contains(element.Text, "sitemap") {
			engine.Visit(element.Text, Crawler.Index)
			println("Index: " + element.Text)
		}
	})
}
