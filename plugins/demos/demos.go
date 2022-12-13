package demos

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("demos", "Demos", "https://www.demos.org/")

	w.SetStartingUrls([]string{
		"https://www.demos.org/our-issues/democratic-reform",
		"https://www.demos.org/our-issues/economic-justice",
		"https://www.demos.org/what-we-do/policy-and-research",
		"https://www.demos.org/what-we-do/litigation",
		"https://www.demos.org/what-we-do/partnerships",
		"https://www.demos.org/what-we-do/policy-and-research",
		"https://www.demos.org/resources",
	})

	// 访问 下一页 Index
	w.OnHTML(`li[class="pager__item pager__item--next"] > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		fmt.Println(element.Attr("href"))
		w.Visit("https://www.demos.org/resources"+element.Attr("href"), Crawler.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.bottom-container h2 > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`div[class="field field--name-node-title field--type-ds field--label-hidden field__item"] > h1`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`div[class="clearfix text-formatted field field--name-field-intro field--type-text-long field--label-hidden field__item"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.article-detail-publication time.datetime`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime (Type 2)
	w.OnHTML(`div[class="field field--name-field-publication-date field--type-datetime field--label-hidden field__item"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`div[class="field field--name-field-article-type field--type-list-string field--label-hidden field__item"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText (Type 2)
	w.OnHTML(`.hero-left div[class="field field--name-field-publication-type field--type-list-string field--label-hidden field__item"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`div[class="field field--name-field-author field--type-entity-reference field--label-hidden field__items"] > div > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`.article-detail-content > div`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Content (Type 2)
	w.OnHTML(`.content-wrapper div.main-content div.field__item`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	w.OnHTML(".main-content > .field > div.field__item:nth-child(-n+2)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	// 获取 Tags
	w.OnHTML(`div[class="field field--name-field-related-issues field--type-entity-reference field--label-hidden field__items"] > .field__item > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 Tags (Type 2)
	w.OnHTML(`ol[class="toc-list "] > li > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`a[type="application/pdf"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		file_url := "https://www.demos.org/" + element.Attr("href")
		ctx.File = append(ctx.File, file_url)
	})
}
