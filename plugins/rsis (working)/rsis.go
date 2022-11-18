package rsis

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("rsis", "拉惹勒南国际研究院", "https://www.rsis.edu.sg/")
	
	w.SetStartingUrls([]string{
		"https://www.rsis.edu.sg/research/cens/",
	})

	// 访问下一页 Index
	w.OnHTML(`.board_pager > a[class="arr next"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 访问 News 从 Index
	w.OnHTML(`div > div.post_text > div > h2 > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.News)
		})

	// 获取 Title
	w.OnHTML(`div[class="blog_single blog_holder"] h2.entry_title`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Description
	w.OnHTML(`div.entry-summary > p`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Description = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`span[class="date entry_date updated"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 CategoryText
	w.OnHTML(`a[rel="category tag"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 Authors
	w.OnHTML(`a.post_author_link`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 CommentCount
	w.OnHTML(`a.post_comments`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			// 裁出数字的字符串并将其转换为 int 类型
			var str = strings.Replace(element.Text, "Comments", "", 1)
			str = strings.TrimSpace(str)
			num, _ := strconv.Atoi(str)
			ctx.CommentCount = num
		})

	// 获取 LikeCount
	w.OnHTML(`.qode-like`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			// 裁出数字的字符串并将其转换为 int 类型
			var str = strings.Replace(element.Text, "Likes", "", 1)
			str = strings.TrimSpace(str)
			num, _ := strconv.Atoi(str)
			ctx.LikeCount = num
		})

	// 获取 ViewCount
	w.OnHTML(`#content > div > div.sub_top_view.con_in > span.look > em`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			str = strings.TrimSpace(element.Text)
			num, _ := strconv.Atoi(str)
			ctx.ViewCount = num
		})

	// 获取 Content
	w.OnHTML(`div.elementor-widget-container`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})

	// 获取 Tags
	w.OnHTML(`a[rel="tag"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
		})

	// 获取 File
	w.OnHTML(`a[type="application/pdf"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			file_url := "https://www.demos.org" + element.Attr("href")
			ctx.File = append(ctx.File, file_url)
		})
}
