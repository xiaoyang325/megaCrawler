package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("news_fit", "佛罗里达理工学院", "https://news.fit.edu/")

	w.SetStartingUrls([]string{
		"https://news.fit.edu/sitemap_index.xml",
	})

	// 访问 sitemap
	w.OnXML(`//loc`, func(element *colly.XMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Text, ".xml") {
			w.Visit(element.Text, Crawler.Index)
		} else {
			w.Visit(element.Text, Crawler.News)
		}
	})

	// 获取 Title
	w.OnHTML(`[class="post-title entry-title"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`.entry-sub-title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.entry-header [class="date meta-item fa-before"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`.meta-author .author-name`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`[class="entry-content entry clearfix"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p"))
	})
}
