package test2

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/megaCrawler"
	"time"
)

func init() {
	s := megaCrawler.Register("go2", "https://go.dev")
	s.UrlGetter.OnHTML(".Hero-blurb", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("url", "https://go.dev")
		element.Request.Ctx.Put("lastMod", time.Now())
	})
	s.UrlProcessor.OnHTML(".Hero-gopherLadder", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("title", element.Attr("alt"))
	})
	s.UrlProcessor.OnHTML(".Hero-blurbList", func(element *colly.HTMLElement) {
		element.Request.Ctx.Put("content", element.Text)
	})
}
