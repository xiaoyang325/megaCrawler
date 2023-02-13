package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("wtc", "World Trade Center", "https://wtc.com")
	w.SetStartingURLs([]string{"https://www.wtc.com/wtc_sitemap.xml"})

	w.OnResponse(func(r *colly.Response, ctx *crawlers.Context) {
		crawlers.Sugar.Infow("Response received", "url", r.Request.URL, "status", r.StatusCode, "DOM", string(r.Body))
	})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/news/") {
			w.Visit(element.Text, crawlers.News)
		}
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML("div.news-article-content.dark-grey > p.med-grey > span:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".news-article-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".news-article-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.ChildText("span")
	})
}
