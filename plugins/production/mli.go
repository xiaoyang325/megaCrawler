package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("mli", "麦克唐纳-劳里埃研究所", "https://macdonaldlaurier.ca/")
	w.SetStartingURLs([]string{"https://www.macdonaldlaurier.ca/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.HasSuffix(element.Text, ".xml") {
			w.Visit(element.Text, crawlers.Index)
		} else if strings.Contains(ctx.URL, "post-sitemap") {
			w.Visit(element.Text, crawlers.News)
		}
	})

	w.OnHTML(".jeg_post_title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".jeg_meta_date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML(".content-inner", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}
