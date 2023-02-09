package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"strings"
)

func init() {
	w := Crawler.Register("investigatemidwest", "中西部调查报道中心", "https://www.investigatemidwest.org/")

	w.SetStartingUrls([]string{"https://investigatemidwest.org/post.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		w.Visit(element.Text, Crawler.News)
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		Extractors.Titles(ctx, element)
		Extractors.PublishingDate(ctx, element)
	})

	w.OnHTML(".subtitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.SubTitle = element.Text
	})

	w.OnHTML(".entry-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = Crawler.HTML2Text(strings.TrimSpace(element.Text))
	})

	w.OnHTML(".post-category-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})
}
