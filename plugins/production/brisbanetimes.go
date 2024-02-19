package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("brisbanetimes", "brisbanetimes", "https://www.brisbanetimes.com.au/")

	w.SetStartingURLs([]string{
		"https://www.brisbanetimes.com.au/",
	})

	w.OnHTML("span[data-testid=\"byline\"]>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("section[data-testid=\"article-body-top\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("span[data-testid=\"byline\"]>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})
	w.OnHTML("nav>div>ul>li>ul>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("nav>div>ul>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a[data-testid=\"article-link\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("time[data-testid=\"datetime\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		var datetime = element.Attr("datetime")
		if len(datetime) > len(ctx.PublicationTime) {
			ctx.PublicationTime = datetime
		}
	})
	w.OnHTML("h1[data-testid=\"headline\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}
