package aflcio

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("aflcio", "美国劳工联合会", "https://www.aflcio.org/")
	w.SetStartingUrls([]string{"https://www.aflcio.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "sitemap") {
			w.Visit(element.Text, Crawler.Index)
		} else {
			w.Visit(element.Text, Crawler.News)
		}
	})

	w.OnHTML("*[property=\"schema:name\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".byline", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	w.OnHTML("time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Attr("datetime")
	})

	w.OnHTML(".section-article-body img", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})

	w.OnHTML(".section-article-body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".content-col h5", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = element.Text
	})
}
