package dev

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("bbc", "英国广播公司", "https://www.bbc.com/")
	w.SetStartingURLs([]string{"https://www.bbc.com/sitemaps/https-index-com-news.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/sitemaps/") {
			w.Visit(element.Text, crawlers.Index)
		}
		if strings.Contains(element.Text, "/news/") {
			w.Visit(element.Text, crawlers.News)
		}
	})

	// 新闻标题
	w.OnHTML("h1#main-heading", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 新闻正文
	w.OnHTML("article > div > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
}
