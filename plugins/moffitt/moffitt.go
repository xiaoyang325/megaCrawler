package moffitt

import (
	"github.com/araddon/dateparse"
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
	"time"
)

func init() {
	w := Crawler.Register("moffitt", "莫菲特癌症中心", "https://moffitt.org/")

	w.SetStartingUrls([]string{"https://moffitt.org/XMLsitemap"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		w.Visit(element.Text, Crawler.Index)
	})

	w.OnHTML(".article-head", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".m-article__content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PageType = Crawler.News
		ctx.Content = Crawler.HTML2Text(strings.TrimSpace(element.Text))
	})

	w.OnHTML("h3", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = element.Text
	})

	w.OnHTML(".fa-tag", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	w.OnHTML(".article > .text-sm", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		t, _ := dateparse.ParseAny(element.Text)
		ctx.PublicationTime = t.Format(time.RFC3339)
	})
}
