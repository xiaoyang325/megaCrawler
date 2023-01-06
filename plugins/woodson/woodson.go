package woodson

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("woodson", "Woodson center", "http://woodsoncenter.org/")
	w.SetStartingUrls([]string{"http://woodsoncenter.org/post-sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		parsed, err := element.Request.URL.Parse("url Here")
		if err != nil {
			return
		}
		w.Visit(parsed.String(), Crawler.News)
		w.Visit(element.Text, Crawler.News)
	})

	w.OnHTML(".p-date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".p-author", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimPrefix(element.Text, "Posted By "))
	})

	w.OnHTML(".et_pb_post_content ", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".et_pb_post_content a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), "pdf") {
			ctx.Link = append(ctx.Link, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})
}
