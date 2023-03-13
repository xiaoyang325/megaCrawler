package production

import (
	"strings"

	"megaCrawler/crawlers"
	"megaCrawler/extractors"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("cyber_fsi_stanford", "网络政策研究所",
		"https://cyber.fsi.stanford.edu/")

	w.SetStartingURLs([]string{
		"https://cyber.fsi.stanford.edu/io",
		"https://cyber.fsi.stanford.edu/people",
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		extractors.Titles(ctx, element)
	})

	// 访问下一页 Index
	w.OnHTML(`.pager > .pager-next > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.block-publications__title > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Expert 从 Index
	w.OnHTML(`.block-peoples__title > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 获取 Name
	w.OnHTML(`div.block-hero-content__group-right > h2`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = strings.TrimSpace(element.Text)
	})

	// 获取 Email
	w.OnHTML(`[class="field field-name-field-email field-type-email field-label-hidden"] a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Email = strings.TrimSpace(element.Text)
	})

	// 获取 Phone
	w.OnHTML(`[class="field field-name-field-phone field-type-text field-label-hidden"] div[class="field-item even"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Phone = strings.TrimSpace(element.Text)
	})

	// 获取 LinkedInID or TwitterID
	w.OnHTML(`div.block-hero-content__links > div > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), "linkedin.com") {
			ctx.LinkedInID = strings.TrimSpace(element.Text)
		} else if strings.Contains(element.Attr("href"), "twitter.com") {
			ctx.TwitterID = strings.TrimSpace(element.Text)
		}
	})

	// 获取 Expert's Title
	w.OnHTML(`[class="block-hero-content__position list-in-article"] > ul`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`[class="news-header__date text--red"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`[class="news-header__body field-type-text-long"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.news-header__date > .date-display-single`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`.additional-content__name`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`.field-type-text-long, .story-content`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p"))
	})

	// 获取 Tags
	w.OnHTML(`.cat-links > .meta-category > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`div[class="field-item even"] > p > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Text, "Download") {
			ctx.File = append(ctx.File, element.Attr("href"))
		}
	})
}
