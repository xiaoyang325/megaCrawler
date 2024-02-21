package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("india", "india", "https://zeenews.india.com/")

	w.SetStartingURLs([]string{
		"https://zeenews.india.com/",
	})

	w.OnHTML("div.articleauthor_details", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("div#fullArticle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("div.category-slider>div>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("div.news_description>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.articleauthor_details > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "IST") {
			ctx.PublicationTime = element.Text
		}
	})
	w.OnHTML("body > section.main-container.articledetails-page > div.container > div > div > div:nth-child(3)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}
