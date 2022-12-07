package cpj

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("cpj", "国际新闻自由组织", "https://cpj.org/")
	w.SetStartingUrls([]string{"https://cpj.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(ctx.Url, "post-sitemap") {
			w.Visit(element.Text, Crawler.News)
		} else if strings.Contains(ctx.Url, "people-sitemap") {
			w.Visit(element.Text, Crawler.Expert)
		} else if strings.Contains(ctx.Url, "sitemap") {
			w.Visit(element.Text, Crawler.Index)
		}
	})

	w.OnHTML("time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Attr("datetime")
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.Title = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	w.OnHTML(".post", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".article--meta a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	w.OnHTML(".people", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	w.OnHTML("#jobs", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML("#coverages", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})
}
