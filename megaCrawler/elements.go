package megaCrawler

import (
	"github.com/gocolly/colly/v2"
	"time"
)

func SetTitle(h *colly.HTMLElement, title string) {
	h.Request.Ctx.Put("title", title)
}

func SetContent(h *colly.HTMLElement, content string) {
	if h.Request.Ctx.Get("content") == "" {
		h.Request.Ctx.Put("content", content)
	}
}

func AppendContent(h *colly.HTMLElement, content string) {
	h.Request.Ctx.Put("content", h.Request.Ctx.Get("content")+"\n"+content)
}

func SetAuthor(h *colly.HTMLElement, author string) {
	h.Request.Ctx.Put("author", author)
}

func AppendAuthor(h *colly.HTMLElement, author string) {
	if h.Request.Ctx.Get("author") == "" {
		h.Request.Ctx.Put("author", author)
	} else {
		h.Request.Ctx.Put("author", h.Request.Ctx.Get("author")+", "+author)
	}
}

func SetTime(h *colly.HTMLElement, t time.Time) {
	h.Request.Ctx.Put("time", t)
}
