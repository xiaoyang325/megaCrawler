package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("businessmirror", "businessmirror", "https://businessmirror.com.ph/")

	w.SetStartingURLs([]string{
		"https://businessmirror.com.ph/",
	})

	w.OnHTML("#primary > div.cs-entry__header.cs-entry__header-simple.cs-video-wrap > div > div.cs-entry__header-info > div.cs-entry__post-meta > div.cs-meta-author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("#primary > div.cs-entry__wrap > div > div.cs-entry__content-wrap > div.entry-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content += element.Text
	})
	w.OnHTML("div.cs-page__archive-description", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description += element.Text
	})
	w.OnHTML("a.cs-meta-author-inner", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})
	w.OnHTML("li.menu-item>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("div.cs-page__author-social > div > div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Email = element.Text
	})
	w.OnHTML("h1.cs-page__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name += element.Text
	})
	w.OnHTML(".cs-entry__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML(" div.cs-entry__header > div > div.cs-entry__header-info > div.cs-entry__post-meta > div.cs-meta-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime += element.Text
	})
	w.OnHTML(".cs-entry__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title += element.Text
	})
}
