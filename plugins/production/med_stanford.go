package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func partMed(w *crawlers.WebsiteEngine) {
	// 从翻页器获取下一页 Index 并访问
	w.OnHTML(`div[class="col-md-8 panel-builder-66-col panel-builder-left"] .next > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从 Index 访问 News
	w.OnHTML(`ul[class="news news  "] > li div[class="newsfeed-item-container newsfeed-item-image-container"] a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取 Title
	w.OnHTML(`div[class="heading parbase"] .black`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML("p.news-excerpt", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})

	// 获取 Publication Time
	w.OnHTML(".datePublished", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	// 获取 Authors
	w.OnHTML(`div[itemprop="articleBody"]> p:nth-child(2)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// "October 31, 2022 - By Emily Moskal" -> "Emily Moskal"

		results := strings.Split(element.Text, "By")
		if len(results) == 1 {
			// 没有注明作者，什么也不做
		} else {
			name := results[1]
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(name))
		}
	})

	// 获取 Content
	w.OnHTML(`div[class="main parsys"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 获取 Tags
	w.OnHTML("a.nav-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}
