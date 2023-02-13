package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("worldjusticeproject", "World Justice Project", "https://worldjusticeproject.org")
	w.SetStartingUrls([]string{"https://worldjusticeproject.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "/news/") {
			w.Visit(element.Text, crawlers.News)
		} else if strings.Contains(element.Text, "/world-justice-forum-2022/") {
			w.Visit(element.Text, crawlers.Expert)
		} else {
			w.Visit(element.Text, crawlers.Index)
		}
	})

	w.OnHTML(".page__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			ctx.Title = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		}
	})

	w.OnHTML(".field--field-author-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".field--field-conf-speaker-jobtitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML("time", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Attr("datetime")
	})

	w.OnHTML(".region__content__main", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.News {
			ctx.Content = element.Text
		} else if ctx.PageType == crawlers.Expert {
			ctx.Description = element.Text
		}
	})

	w.OnHTML(".img-responsive", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})
}
