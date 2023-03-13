package production

import (
	"time"

	"megaCrawler/crawlers"

	"github.com/araddon/dateparse"
	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("nih", "美国国家卫生研究院", "https://www.niaid.nih.gov/")
	w.SetStartingURLs([]string{"https://www.niaid.nih.gov/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		w.Visit(element.Text, crawlers.News)
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML("h2", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = element.Text
	})

	w.OnHTML("article > div > p:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		t, err := dateparse.ParseAny(element.Text)
		if err != nil {
			return
		}
		ctx.PublicationTime = t.Format(time.RFC3339)
	})

	w.OnHTML("article", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.ChildText("p")
	})
}
