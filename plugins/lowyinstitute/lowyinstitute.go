package lowyinstitute

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("lowyinstitute", "洛伊国际政策研究所", "https://www.lowyinstitute.org/")

	w.SetStartingUrls([]string{
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=195&related_issues=All",
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=199&related_issues=All",
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=195&related_issues=310",
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=207&related_issues=All",
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=204&related_issues=All",
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=198&related_issues=All",
		"https://www.lowyinstitute.org/publications?issues=200",
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=201&related_issues=All",
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=205&related_issues=All",
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=203&related_issues=All",
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=4&related_issues=All",
		"https://www.lowyinstitute.org/publications?keys=&author=All&type=All&issues=202&related_issues=All",
	})

	// 访问下一页 Index
	w.OnHTML(`#block-lowy-content > div > div > div.margin-horizontal > div > nav > ul > li.pager__item.pager__item--next.has-previous-item > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 访问 Report 从 Index
	w.OnHTML(`#block-lowy-content > div > div > div.view-content a.card__title`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 获取 Title
	w.OnHTML(`body > div > section > div.flex.relative.above > div > h1`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Title
	w.OnHTML(`#block-lowy-content > article > div > div.flexible-content-page__top.margin-horizontal.margin-t-4 > div > h1`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Description
	w.OnHTML(`body > div > section > div.flex.relative.above > div > p.caption`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Description = strings.TrimSpace(element.Text)
		})

	// 获取 Description
	w.OnHTML(`#block-lowy-content > article > div > div.flexible-content-page__top.margin-horizontal.margin-t-4 > div > div.article-intro.intro-text`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Description = strings.TrimSpace(element.Text)
		})

	// 获取 Content
	w.OnHTML(".flexible-content__main-content",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`body > div > section > div.flex.relative.above > div > span`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`#block-lowy-content > article > div > div.flexible-content-page__top.margin-horizontal.margin-t-4 > div > div.article-published-date`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 CategoryText
	w.OnHTML(`#block-lowy-content > article > div > div.flexible-content-page__top.margin-horizontal.margin-t-4 > div > div.article-category`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 Authors
	w.OnHTML(`body > main > article > p.caption.txt-muted > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 Authors
	w.OnHTML(`#block-lowy-content > article > div > div.flexible-content-page__top.margin-horizontal.margin-t-4 > div > div.article-authors > span > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 Authors
	w.OnHTML(`body > div > aside > div > div[class="post flex-col"] > a > div`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 Content
	w.OnHTML(`div.elementor-widget-container`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})

	// 获取 File
	w.OnHTML(`div[class="flex-col footer-actions"] > #download`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			file_url := "https://interactives.lowyinstitute.org" + element.Attr("href")
			ctx.File = append(ctx.File, file_url)
		})
}
