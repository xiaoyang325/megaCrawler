package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

// 由于不知道的原因，Tag 和 Authors 会被重复添加，这个函数用来阻止
func strInSlice(name string, nameList *[]string) bool {
	for _, value := range *nameList {
		if name == value {
			return true
		}
	}
	return false
}

func partGsb(w *crawlers.WebsiteEngine) {
	// 从翻页器获取更多（Show More）并访问
	w.OnHTML(`a[title="Go to next page"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从 Index 访问 Report
	w.OnHTML(`div[class="field__item field--item-view_item"] div[class="views-field views-field-title"] a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title
	w.OnHTML(`h1[class="heading has-icon icon-written-before has-key-taxonomy-above"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 获取 Description
	w.OnHTML(".summary p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})

	// 获取 Publication Time
	w.OnHTML(".date p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// 获取 Authors
	w.OnHTML(".authors .author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// "Adam Hadhazy," -> "Adam Hadhazy"
		name := strings.ReplaceAll(element.Text, ",", "")
		name = strings.TrimSpace(name)
		if !strInSlice(name, &ctx.Authors) {
			ctx.Authors = append(ctx.Authors, name)
		}
	})

	// 获取 Content
	w.OnHTML(`div[class="announcement-stories__idea-story-description as-description"] > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 获取 Tags
	w.OnHTML(".taxonomy-links .align-inline-element .link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if !strInSlice(element.Text, &ctx.Tags) {
			ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
		}
	})
}
