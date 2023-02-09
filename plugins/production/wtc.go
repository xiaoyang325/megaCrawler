package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("wtc", "World Trade Center", "https://wtc.com")
	w.SetStartingUrls([]string{"https://www.wtc.com/wtc_sitemap.xml"})

	w.OnResponse(func(r *colly.Response, ctx *Crawler.Context) {
		Crawler.Sugar.Infow("Response received", "url", r.Request.URL, "status", r.StatusCode, "DOM", string(r.Body))
	})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "/news/") {
			w.Visit(element.Text, Crawler.News)
		}
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML("div.news-article-content.dark-grey > p.med-grey > span:nth-child(2)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".news-article-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	w.OnHTML(".news-article-date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.ChildText("span")
	})
}
