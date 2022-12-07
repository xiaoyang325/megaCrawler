package ned

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("ned", "美国国家民主基金会", "https://www.ned.org/")
	w.SetStartingUrls([]string{"https://www.ned.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		w.Visit(element.Text, Crawler.News)
	})

	w.OnHTML("time.published", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Attr("datetime")
	})

	w.OnHTML(".entry-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
