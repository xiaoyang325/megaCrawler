package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("reproductiverights", "生殖权利中心",
		"https://reproductiverights.org/")

	w.SetStartingUrls([]string{
		"https://reproductiverights.org/?s=",
		"https://reproductiverights.org/covid-19-resources/",
	})

	// 访问下一页 Index
	w.OnHTML(`div.c-pagination > div.c-pagination__next > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.c-post-card__title > a.c-post-card__link`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`.entry__header > h1.entry__title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`h5.entry__subtitle`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`time.entry__published`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`div.entry__categories > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 covid-19-resources 的所有 Report
	w.OnHTML(`#post-29576 p`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = Crawler.Report

		file_url := element.ChildAttr("em > a", "href")
		if strings.Contains(file_url, ".pdf") {
			subCtx.File = append(subCtx.File, file_url)
		}

		subCtx.Title = element.ChildText("em > a")
		subCtx.Description = element.Text
	})

	// 获取 Content
	w.OnHTML(`div.entry__content`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p, h1, h2, h3, h4"))
	})

	// 获取 Tags
	w.OnHTML(`.c-related-content-list__content > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		tag := strings.TrimSpace(element.Text)
		if tag == "News" {
			ctx.PageType = Crawler.News
		}
		ctx.Tags = append(ctx.Tags, tag)
	})

	// 获取 Tags
	w.OnHTML(`[class="c-post-card__cat tags"] > p > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`.file-attachments > ul > li > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}
