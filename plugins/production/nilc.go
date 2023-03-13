package production

import (
	"strings"
	"time"

	"megaCrawler/crawlers"

	"github.com/araddon/dateparse"
	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("nilc", "美国国家移民法律中心", "https://nilc.org/")
	w.SetStartingURLs([]string{"https://nilc.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		w.Visit(element.Text, crawlers.News)
	})

	w.OnHTML("div.wpb_column.vc_column_container.vc_col-sm-9 > div > div > div > div > p:nth-child(2) > strong", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML("a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), "pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	w.OnHTML(".wpb_text_column > .wpb_wrapper", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.ChildText("p, h3")
	})

	w.OnHTML("strong", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		t, err := dateparse.ParseAny(element.Text)
		if err == nil {
			ctx.PublicationTime = t.Format(time.RFC3339)
		}
	})
}
