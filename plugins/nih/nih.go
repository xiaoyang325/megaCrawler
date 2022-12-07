package nih

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("nih", "美国国家卫生研究院", "https://www.nih.gov/")
	w.SetStartingUrls([]string{"https://www.nih.gov/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "/news-releases/") {
			w.Visit(element.Text, Crawler.News)
		}
		if strings.Contains(element.Text, "sitemap") {
			w.Visit(element.Text, Crawler.Index)
		}
	})

	w.OnHTML("*[property=\"dc:date\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Attr("content")
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".subtitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.SubTitle = element.Text
	})

	w.OnHTML(".l-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.ChildText(":not([class])")
	})
}
