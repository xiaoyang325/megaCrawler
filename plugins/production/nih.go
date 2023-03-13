package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("nih", "美国国家卫生研究院", "https://www.nih.gov/")
	w.SetStartingURLs([]string{"https://www.nih.gov/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/news-releases/") {
			w.Visit(element.Text, crawlers.News)
		}
		if strings.Contains(element.Text, "sitemap") {
			w.Visit(element.Text, crawlers.Index)
		}
	})

	w.OnHTML("*[property=\"dc:date\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Attr("content")
	})

	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".subtitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = element.Text
	})

	w.OnHTML(".l-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.ChildText(":not([class])")
	})
}
