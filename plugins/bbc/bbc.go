package bbc

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("bbc", "英国广播公司", "https://www.bbc.com/")
	w.SetStartingUrls([]string{"https://www.bbc.com/sitemaps/https-index-com-news.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "/sitemaps/") {
			w.Visit(element.Text, Crawler.Index)
		}
		if strings.Contains(element.Text, "/news/") {
			w.Visit(element.Text, Crawler.News)
		}
	})

	//新闻标题
	w.OnHTML("h1#main-heading", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//新闻正文
	w.OnHTML("article > div > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content += element.Text
	})
}
