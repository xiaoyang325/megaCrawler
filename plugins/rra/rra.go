package rra

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"strings"
)

func init() {
	w := Crawler.Register("rra", "激进派右翼分析中心", "https://www.radicalrightanalysis.com/")

	w.SetStartingUrls([]string{"https://www.radicalrightanalysis.com/wp-sitemap-posts-post-1.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		w.Visit(element.Text, Crawler.News)
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		Extractors.Titles(ctx, element)
		Extractors.Authors(ctx, element)
		Extractors.PublishingDate(ctx, element)
		Extractors.Tags(ctx, element)
	})

	w.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
