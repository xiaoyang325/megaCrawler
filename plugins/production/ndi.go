package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"strings"
)

func init() {
	w := Crawler.Register("ndi", "国际民主研究所", "https://www.ndi.org/")
	w.SetStartingUrls([]string{"https://www.ndi.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		w.Visit(element.Text, Crawler.News)
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		Extractors.Titles(ctx, element)
	})

	w.OnHTML(".file > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
		if strings.Contains(element.Attr("href"), "pdf") {
			ctx.PageType = Crawler.Report
		}
	})

	w.OnHTML(".body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML("a[typeof=\"skos:Concept\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	w.OnHTML("span[property=\"dc:date\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Attr("content")
	})
}
