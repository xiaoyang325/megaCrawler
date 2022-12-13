package quincyinst

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("quincyinst", "昆西负责任治国方略研究所",
		"https://quincyinst.org/")

	w.SetStartingUrls([]string{
		"https://quincyinst.org/category/middle-east/",
		"https://quincyinst.org/category/east-asia/",
		"https://quincyinst.org/category/grand-strategy/",
		"https://quincyinst.org/category/democratizing-foreign-policy/",
		"https://quincyinst.org/category/eurasia/",
		"https://quincyinst.org/category/global-partners/",
		"https://quincyinst.org/category/reports/",
		"https://quincyinst.org/press/",
		"https://quincyinst.org/events/",
	})

	// 从Index中访问文章。
	w.OnHTML(".link>.post__title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), "/report/") {
			w.Visit(element.Attr("href"), Crawler.Report)
		} else {
			w.Visit(element.Attr("href"), Crawler.News)
		}

	})

	// 从Index中访问文章。(/press/)
	w.OnHTML("article>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), "/press/") {
			w.Visit(element.Attr("href"), Crawler.Report)
		}
	})

	// 从Index中访问文章。(/events/)
	w.OnHTML("section[class=\"events events__past\"]>.events-list>article>.card-body>.entry-title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 从翻页器获取下一页Index并访问。
	w.OnHTML("a[class=\"next page-numbers\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	w.OnHTML("meta[property=\"og:title\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.ReplaceAll(element.Attr("content"), " - Quincy Institute for Responsible Statecraft", "")
	})

	w.OnHTML("meta[property=\"article:published_time\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Attr("content")
	})

	// 从文章中获取Author并添加到ctx。（Report）
	w.OnHTML("header.post__header>div.post__header--container>div.post__author--container>div.post__author--info>div.post__author>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 从文章中获取Author并添加到ctx。（News）
	w.OnHTML("header.post__header>div.post__header--container>div>div>div>div>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 从文章中获取Tag并添加到ctx。（Report）（News）
	w.OnHTML("div.post__tag>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 从文章中获取Content并添加到ctx。（Report）
	w.OnHTML(".report-body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 从文章中获取Content并添加到ctx。（News）
	w.OnHTML("article>section.post__content>div.post__content--article", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 从文章中获取Content并添加到ctx。（/press/）
	w.OnHTML("div.press-release-content__container", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 从文章中获取Content并添加到ctx。（/event/）
	w.OnHTML(".event-body>.event-description", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 从文章中获取Location并添加到ctx。（/event/）
	w.OnHTML(".event__location--copy", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Location = strings.TrimSpace(element.Text)
	})
}
