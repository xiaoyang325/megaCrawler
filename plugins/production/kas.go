package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("kas", "康拉德·阿登纳基金会", "https://www.kas.de/")

	w.SetStartingUrls([]string{
		"https://www.kas.de/de/publikationen",
	})

	// 访问下一页 Index
	w.OnHTML(`[class="lfr-pagination-buttons pager"] > li[class=""] > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`#column-2 a.c-publication-tiles__tile`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`div > div > div > div.c-page-intro > div.o-page-module.o-page-module--bare-bottom > h1`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`h2.c-page-intro__subheadline`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`div.c-page-intro__copy`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`span[class="o-metadata o-metadata--date"] > svg`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`div.o-page-module.o-page-module--bare > h4 > span > em`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`div.c-page-intro > div.o-page-module.o-page-module--bare-bottom > div.c-page-intro__author`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		name := strings.TrimSpace(strings.Replace(element.Text, "von", "", 1))
		ctx.Authors = append(ctx.Authors, name)
	})

	// 获取 Content
	w.OnHTML(`div[class="c-page-main__text o-richtext"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`span.c-tag__name`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`div.c-aside-teaser.c-aside-teaser--links.c-aside-teaser--covermedia > div.c-aside-teaser__entry > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		fileUrl := "https://www.kas.de" + element.Attr("href")
		ctx.File = append(ctx.File, fileUrl)
	})
}
