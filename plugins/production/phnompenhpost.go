package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("phnompenhpost", "金边邮报", "https://www.phnompenhpost.com/")

	w.SetStartingUrls([]string{
		"https://www.phnompenhpost.com/post-depth",
		"https://www.phnompenhpost.com/politics-0",
		"https://www.phnompenhpost.com/kr-tribunal",
		"https://www.phnompenhpost.com/education",
		"https://www.phnompenhpost.com/financial",
		"https://phnompenhpost.com/post-property/supp-post-property",
		"https://www.phnompenhpost.com/socialite",
		"https://www.phnompenhpost.com/around-ngos",
		"https://www.phnompenhpost.com/supplements",
		"https://www.phnompenhpost.com/pdf-supplement",
		"https://www.phnompenhpost.com/opinion",
		"https://www.phnompenhpost.com/international",
	})

	// 访问下一页 Index
	w.OnHTML(`ul[class="pager pager-load-more"] > li[class="pager-next first last"] > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 News 从 Index
	w.OnHTML(`.view-content > .item-list li .article-image > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// 获取 Title
	w.OnHTML(`body > div.container > div > div.section.section-width-sidebar.single-article-header > h2`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`body > div.container > div > div.section.section-width-sidebar.single-article-header > div > p`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		raw := element.Text
		raw = strings.Replace(raw, element.ChildText("a"), "", 1)
		raw = strings.Replace(raw, "|", "", 1)
		ctx.PublicationTime = strings.TrimSpace(raw)
	})

	// 获取 Authors
	w.OnHTML(`body > div.container > div > div.section.section-width-sidebar.single-article-header > div > p > a  > span`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`div[itemprop="articleBody"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
