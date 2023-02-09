package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("dav_edu", "外交学院", "https://www.dav.edu.vn/")

	w.SetStartingUrls([]string{
		"https://www.dav.edu.vn/su-kien-hoi-thao-toa-dam/?trang=1",
		"https://www.dav.edu.vn/gioi-thieu-chung-nghien-cu/",
		"https://www.dav.edu.vn/an-pham-nghien-cuu/?trang=1",
	})

	// 访问下一页 Index
	w.OnHTML(`[class="page-item active"] > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		indexURL := Crawler.GetNextIndexURL(ctx.Url, element.Text, "trang")
		w.Visit(indexURL, Crawler.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.row .story__title > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`.detail__title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`.detail__summary`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.detail__meta`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`li[class="breadcrumb-item active"] > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`html .detail__content`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
