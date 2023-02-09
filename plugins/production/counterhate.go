package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("counterhate", "打击数字仇恨中心", "https://counterhate.com/")

	w.SetStartingUrls([]string{
		"https://counterhate.com/topic/anti-muslim-hate/",
		"https://counterhate.com/topic/antisemitism/",
		"https://counterhate.com/topic/inside-the-metaverse/",
		"https://counterhate.com/topic/misogyny/",
		"https://counterhate.com/topic/anti-vaxx-misinformation/",
		"https://counterhate.com/topic/election-and-state-media-misinformation/",
		"https://counterhate.com/topic/climate-change-misinformation/",
		"https://counterhate.com/topic/sexual-reproductive-rights/",
		"https://counterhate.com/topic/stop-funding-misinformation/",
	})

	// 从子频道入口访问 "View all research"
	w.OnHTML(".topic-hub-research__intro>p>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从翻页器访问下一页 Index
	w.OnHTML(".pagination__next>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从 Index 访问报告
	w.OnHTML(".research-post-single__view-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 添加 Title 到ctx
	w.OnHTML(".research-post__title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 添加 SubTitle 到ctx
	w.OnHTML(".research-post__subtitle", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 添加 Tag 到 ctx.Tags
	w.OnHTML(".topic-tag", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	// 添加 Content 到 ctx
	w.OnHTML("div[class=\"research-post__content flow\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 添加 File 到 ctx
	w.OnHTML("a[class=\"button button--green\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	// 添加 PublicationTime 到 ctx
	w.OnHTML(".research-post__about > p > time", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	// 添加 Description 到 ctx
	w.OnHTML(".research-post__intro-text", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
}
