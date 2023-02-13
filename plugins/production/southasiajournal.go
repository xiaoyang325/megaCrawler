package production

import (
	"fmt"
	"megaCrawler/crawlers"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("southasiajournal", "南亚分析集团", "http://southasiajournal.net")

	w.SetStartingURLs([]string{
		"http://southasiajournal.net/category/events/",
		"http://southasiajournal.net/category/e-saj-features/",
		"http://southasiajournal.net/category/blog/",
		"http://southasiajournal.net/category/reviews/",
		"http://southasiajournal.net/category/environment/",
	})

	// 从翻页器中获取(/e-saj-features, /blog, /reviews中的）所有页的链接并访问，并将其标注为Index。
	w.OnHTML("div[class=\"page-nav td-pb-padding-side\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 仅在/e-saj-features中执行。
		if strings.Contains(ctx.URL, "/e-saj-features") {
			// 仅在第一页执行。
			if element.ChildText("span[class=\"current\"]") == "1" {
				// 获取最大页数的字符串。
				strMaxI := element.ChildText("a[class=\"last\"]")
				// 将最大页数转换为正数类型。
				maxI, _ := strconv.ParseInt(strMaxI, 10, 64)
				// 通过for循环获取并访问从2到最大页数的所有页面。
				for i := int64(2); i <= maxI; i++ {
					w.Visit(fmt.Sprintf("http://southasiajournal.net/category/e-saj-features/page/%d", i),
						crawlers.Index)
				}
			}
		}

		// 仅在/blog中执行。
		if strings.Contains(ctx.URL, "/blog") {
			// 仅在第一页执行。
			if element.ChildText("span[class=\"current\"]") == "1" {
				// 获取最大页数的字符串。
				strMaxI := element.ChildText("a[class=\"last\"]")
				// 将最大页数转换为正数类型。
				maxI, _ := strconv.ParseInt(strMaxI, 10, 64)
				// 通过for循环获取并访问从2到最大页数的所有页面。
				for i := int64(2); i <= maxI; i++ {
					w.Visit(fmt.Sprintf("http://southasiajournal.net/category/blog/page/%d", i),
						crawlers.Index)
				}
			}
		}

		// 仅在/reviews中执行。
		if strings.Contains(ctx.URL, "/reviews") {
			// 仅在第一页执行。
			if element.ChildText("span[class=\"current\"]") == "1" {
				// 获取最大页数的字符串。
				strMaxI := element.ChildText("a[class=\"last\"]")
				// 将最大页数转换为正数类型。
				maxI, _ := strconv.ParseInt(strMaxI, 10, 64)
				// 通过for循环获取并访问从2到最大页数的所有页面。
				for i := int64(2); i <= maxI; i++ {
					w.Visit(fmt.Sprintf("http://southasiajournal.net/category/reviews/page/%d", i),
						crawlers.Index)
				}
			}
		}
	})

	// 从页面获取链接并访问文章(/events, /e-saj-features, /blog, /reviews中的)，并将其标注，
	// 将（/events, /e-saj-features, /blog）标记为News，
	// 将（/reviews, /environment）标记为Report。
	w.OnHTML("div[class=\"td-block-span6\"] h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 若URL中包含/reviews或是/environment，则是从此访问，标记为Report。
		if strings.Contains(ctx.URL, "/reviews") || strings.Contains(ctx.URL, "/environment") {
			w.Visit(element.Attr("href"), crawlers.Report)
		} else { // 否则标记为News。
			w.Visit(element.Attr("href"), crawlers.News)
		}
	})

	// 从文章中（/events, /e-saj-features, /blog, /reviews中的）添加标题(title)到ctx。
	w.OnHTML("h1[class=\"entry-title\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".td-post-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML("meta[property=\"og:description\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Attr("content")
	})

	// 从文章中（/events, /e-saj-features, /blog, /reviews中的）添加作者(author)到ctx。
	w.OnHTML("div[class=\"td-post-author-name\"] > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 不在Index中添加。
		if ctx.PageType != crawlers.Index {
			ctx.Authors = append(ctx.Authors, element.Text)
		}
	})

	// 从网页（/events, /e-saj-features, /blog, /reviews中的）获取分类（category）到ctx中。
	w.OnHTML("li[class=\"entry-category\"] > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 若分类为空则直接添加。
		if ctx.CategoryText == "" {
			ctx.CategoryText = element.Text
		} else { // 若已有分类则合并。
			ctx.CategoryText += " " + element.Text
		}
	})

	// 从文章中（/events, /e-saj-features, /blog, /reviews中的）获取文章正文到ctx。
	w.OnHTML("div[class=\"td-ss-main-content\"] > article > div[class=\"td-post-content tagdiv-type\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}
