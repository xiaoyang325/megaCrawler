package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("pids_gov_ph", "发展研究院", "https://pids.gov.ph/")

	w.SetStartingUrls([]string{
		"https://pids.gov.ph/research-agenda",
		"https://pids.gov.ph/research",
		"https://pids.gov.ph/publications/category/books",
		"https://pids.gov.ph/publications/category/research-paper-series",
		"https://pids.gov.ph/content/publication/index-pjd",
		"https://pids.gov.ph/publications/category/economic-policy-monitor",
		"https://pids.gov.ph/publications/category/discussion-papers",
		"https://pids.gov.ph/publications/category/development-research-news",
		"https://pids.gov.ph/publications/category/policy-notes",
		"https://pids.gov.ph/publications/category/economic-issue-of-the-day",
		"https://pids.gov.ph/publications/category/annual-reports",
		"https://pids.gov.ph/publications/category/working-papers",
		"https://pids.gov.ph/publications/category/monograph-series",
		"https://pids.gov.ph/publications/category/staff-papers",
		"https://pids.gov.ph/publications/category/economic-outlook-series",
		"https://pids.gov.ph/content/publication/index-archive",
		"https://pids.gov.ph/events",
		"https://pids.gov.ph/content/public/search-contenttype?contenttype=news&category=press-releases",
		"https://pids.gov.ph/content/public/search-contenttype?contenttype=news&category=in-the-news",
		`https://pids.gov.ph/content/public/index-custom?view=news%2Findex-pids-updates&category=news`,
		"https://pids.gov.ph/legislative-inputs",
	})

	// 访问下一页 Index
	w.OnHTML(`[class="page-item next"] > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.card-content > .card-title > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title
	w.OnHTML(`.page-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.content-max-600 > div:nth-child(1) > .meta-value`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.sidebar-links > div:nth-child(2)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`div.content > div > div:nth-child(2) > div > a > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`div.content > div > div:nth-child(4) > div.meta-value.list > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`.content p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`div.content > div > div:nth-child(5) > div.meta-value.list-boxed > ul > li> a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 Tags
	w.OnHTML(`.sidebar-links > div:nth-child(8) > [class="meta-value list-boxed"] li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`#modalDownload > #w2`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		file_url := "https://pids.gov.ph" + element.Attr("action")
		ctx.File = append(ctx.File, file_url)
	})

	// 获取 File
	w.OnHTML(`a[class="btn btn-primary"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}
