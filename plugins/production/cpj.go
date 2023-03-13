package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("cpj", "国际新闻自由组织", "https://cpj.org/")
	w.SetStartingURLs([]string{"https://cpj.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		switch {
		case strings.Contains(ctx.URL, "post-sitemap"):
			w.Visit(element.Text, crawlers.News)
		case strings.Contains(ctx.URL, "people-sitemap"):
			w.Visit(element.Text, crawlers.Expert)
		case strings.Contains(ctx.URL, "sitemap"):
			w.Visit(element.Text, crawlers.Index)
		}
	})

	w.OnHTML("time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Attr("datetime")
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			ctx.Title = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		}
	})

	w.OnHTML(".post", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".article--meta a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	w.OnHTML(".people", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})

	w.OnHTML("#jobs", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML("#coverages", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Area = element.Text
	})
}
