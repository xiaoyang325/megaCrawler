package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("mli", "麦克唐纳-劳里埃研究所", "https://macdonaldlaurier.ca/")
	w.SetStartingUrls([]string{"https://www.macdonaldlaurier.ca/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.HasSuffix(element.Text, ".xml") {
			w.Visit(element.Text, Crawler.Index)
		} else if strings.Contains(ctx.Url, "post-sitemap") {
			w.Visit(element.Text, Crawler.News)
		}
	})

	w.OnHTML(".jeg_post_title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".jeg_meta_date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML(".content-inner", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
