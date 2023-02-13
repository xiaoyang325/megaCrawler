// Package extractors Provide General Information Extractor for standard websites
package extractors

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

type Config struct {
	Author       bool
	Image        bool
	Language     bool
	PublishDate  bool
	Tags         bool
	Text         bool
	Title        bool
	TextLanguage string
}

func (c *Config) Apply(w *crawlers.WebsiteEngine) {
	w.OnHTML("html", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if c.Author {
			Authors(ctx, element)
		}
		if c.Image {
			Image(ctx, element)
		}
		if c.Language {
			Language(ctx, element)
		}
		if c.PublishDate {
			PublishingDate(ctx, element)
		}
		if c.Tags {
			Tags(ctx, element)
		}
		if c.Text {
			if c.TextLanguage == "" {
				c.TextLanguage = ctx.Language
			}
			Text(ctx, element, c.TextLanguage)
		}
		if c.Title {
			Titles(ctx, element)
		}
	})
}
