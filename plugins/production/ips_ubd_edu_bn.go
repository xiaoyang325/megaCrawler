package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("ips_ubd_edu_bn", "战略与政策研究中心", "https://ips.ubd.edu.bn/")

	w.SetStartingURLs([]string{
		"https://ips.ubd.edu.bn/category/news/",
		"https://ips.ubd.edu.bn/category/seminar-series/",
	})

	// 访问下一页 Index
	w.OnHTML(`.nav-next > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 News 和 Report 从 Index
	w.OnHTML(`div > p.read-more-wrap > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(ctx.URL, "/seminar-series") {
			w.Visit(element.Attr("href"), crawlers.Report)
		} else {
			w.Visit(element.Attr("href"), crawlers.News)
		}
	})

	// 获取 Title
	w.OnHTML(`[class="entry-title page-title h2"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.posted-on > time`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`#main .entry-content`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
