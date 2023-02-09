package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("ips_ubd_edu_bn", "战略与政策研究中心", "https://ips.ubd.edu.bn/")

	w.SetStartingUrls([]string{
		"https://ips.ubd.edu.bn/category/news/",
		"https://ips.ubd.edu.bn/category/seminar-series/",
	})

	// 访问下一页 Index
	w.OnHTML(`.nav-next > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 News 和 Report 从 Index
	w.OnHTML(`div > p.read-more-wrap > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(ctx.Url, "/seminar-series") {
			w.Visit(element.Attr("href"), Crawler.Report)
		} else {
			w.Visit(element.Attr("href"), Crawler.News)
		}
	})

	// 获取 Title
	w.OnHTML(`[class="entry-title page-title h2"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.posted-on > time`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`#main .entry-content`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
