package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strconv"
	"strings"
)

func init() {
	w := Crawler.Register("yerepouni_news", "Yerepouni Daily News",
		"https://www.yerepouni-news.com/")

	w.SetStartingUrls([]string{
		`https://www.yerepouni-news.com/category/01-armenian-news/`,
		`https://www.yerepouni-news.com/category/arm-articles/`,
		`https://www.yerepouni-news.com/category/arm-interviews/`,
		`https://www.yerepouni-news.com/category/arm-sports/`,
		`https://www.yerepouni-news.com/category/%d5%a1%d5%b5%d5%ac%d5%a1%d5%a6%d5%a1%d5%b6/arm-misc/`,
		`https://www.yerepouni-news.com/category/%d5%a1%d5%b5%d5%ac%d5%a1%d5%a6%d5%a1%d5%b6/arm-culture/`,
		`https://www.yerepouni-news.com/category/%d5%a1%d5%b5%d5%ac%d5%a1%d5%a6%d5%a1%d5%b6/arm-social/`,
		`https://www.yerepouni-news.com/category/%d5%ac%d5%b8%d6%82%d6%80%d5%a5%d6%80-%d5%a1%d6%80%d5%a5%d6%82%d5%a5%d5%ac%d5%a1%d5%b0%d5%a1%d5%b5%d5%a5%d6%80%d5%a7%d5%b6/`,
		`https://www.yerepouni-news.com/category/english/3-english-news/africa/`,
		`https://www.yerepouni-news.com/category/english/3-english-news/america/`,
		`https://www.yerepouni-news.com/category/english/3-english-news/asia/`,
		`https://www.yerepouni-news.com/category/english/3-english-news/australia/`,
		`https://www.yerepouni-news.com/category/english/3-english-news/europe/`,
		`https://www.yerepouni-news.com/category/english/3-english-news/middle-east/`,
		`https://www.yerepouni-news.com/category/english/markets-and-economy/`,
		`https://www.yerepouni-news.com/category/english/en-inter-press/`,
		`https://www.yerepouni-news.com/category/english/health/`,
		`https://www.yerepouni-news.com/category/english/social/`,
		`https://www.yerepouni-news.com/category/english/entertainment/`,
		`https://www.yerepouni-news.com/category/english/technology/`,
		`https://www.yerepouni-news.com/category/2-arabic/`,
	})

	// 访问下一页 Index
	w.OnHTML(`.jeg_block_navigation a[class="page_nav next"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 News & Report 从 Index
	w.OnHTML(`.jeg_block_container article > .jeg_thumb > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(ctx.Url, "news/") {
			w.Visit(element.Attr("href"), Crawler.News)
		} else {
			w.Visit(element.Attr("href"), Crawler.Report)
		}
	})

	// 获取 Title
	w.OnHTML(`h1.jeg_post_title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.entry-header .jeg_meta_date > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CommentCount
	w.OnHTML(`i[class="fa fa-comment-o"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		str := strings.TrimSpace(element.Text)
		num, _ := strconv.Atoi(str)
		ctx.CommentCount = num
	})

	// 获取 Content
	w.OnHTML(`[class="content-inner "]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p, h2, h3"))
	})

	// 获取 Tags
	w.OnHTML(`.entry-header .jeg_meta_category > span > a[rel="category tag"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}
