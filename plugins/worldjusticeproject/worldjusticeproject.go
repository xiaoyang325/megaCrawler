package worldjusticeproject

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("worldjusticeproject", "World Justice Project", "https://worldjusticeproject.org")
	w.SetStartingUrls([]string{"https://worldjusticeproject.org/sitemap.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, "/news/") {
			w.Visit(element.Text, Crawler.News)
		} else if strings.Contains(element.Text, "/world-justice-forum-2022/") {
			w.Visit(element.Text, Crawler.Expert)
		} else {
			w.Visit(element.Text, Crawler.Index)
		}
	})

	w.OnHTML(".page__title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.Title = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		}
	})

	w.OnHTML(".field--field-author-name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".field--field-conf-speaker-jobtitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML("time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Attr("datetime")
	})

	w.OnHTML(".region__content__main", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.Content = element.Text
		} else if ctx.PageType == Crawler.Expert {
			ctx.Description = element.Text
		}
	})

	w.OnHTML(".img-responsive", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Image = append(ctx.Image, element.Attr("src"))
	})
}
