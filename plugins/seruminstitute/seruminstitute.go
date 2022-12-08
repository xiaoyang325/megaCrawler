package seruminstitute

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("seruminstitute", "印度血清研究所", "https://www.seruminstitute.com/")

	w.SetStartingUrls([]string{
		"https://www.seruminstitute.com/news.php",
	})

	// 访问 News 从 Index
	w.OnHTML(`.listarea > .list-text > h4 > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// 获取 Title
	w.OnHTML(`.news-heading`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`.career-content > h2`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime & Authors 或是 PublicationTime
	w.OnHTML(`span[class="date entry_date updated"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		// 检查是否有 Authors 信息
		if strings.Contains(element.Text, "by") {
			infos := strings.Split(element.Text, "by")
			ctx.PublicationTime = strings.TrimSpace(infos[0])
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(infos[1]))
		} else {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		}
	})

	// 获取 Content
	w.OnHTML(`.newsdetails-content`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p"))
	})
}
