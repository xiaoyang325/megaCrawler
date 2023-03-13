package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("ndi", "国际民主研究所", "https://www.ndi.org/")
	w.SetStartingURLs([]string{"https://www.ndi.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		w.Visit(element.Text, crawlers.News)
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		extractors.Titles(ctx, element)
	})

	w.OnHTML(".file > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
		if strings.Contains(element.Attr("href"), "pdf") {
			ctx.PageType = crawlers.Report
		}
	})

	w.OnHTML(".body", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML("a[typeof=\"skos:Concept\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	w.OnHTML("span[property=\"dc:date\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Attr("content")
	})
}
