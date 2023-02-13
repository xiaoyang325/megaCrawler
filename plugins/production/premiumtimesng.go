package production

import (
	"megaCrawler/crawlers"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("premiumtimesng", "Premium Official News",
		"https://www.premiumtimesng.com/")

	w.SetStartingURLs([]string{
		"https://www.premiumtimesng.com/category/gender",
		"https://www.premiumtimesng.com/category/news",
		"https://www.premiumtimesng.com/category/foreign",
		"https://www.premiumtimesng.com/category/investigationspecial-reports",
		"https://www.premiumtimesng.com/category/business/business-news",
		"https://www.premiumtimesng.com/category/business/financial-inclusion",
		"https://www.premiumtimesng.com/category/business/business-data",
		"https://www.premiumtimesng.com/category/business/business-interviews",
		"https://www.premiumtimesng.com/category/oilgas-reports/faac-reports",
		"https://www.premiumtimesng.com/category/oilgas-reports/revenue-oilgas-reports",
		"https://www.premiumtimesng.com/category/health/health-news",
		"https://www.premiumtimesng.com/category/health/health-investigations",
		"https://www.premiumtimesng.com/category/health/health-interviews",
		"https://www.premiumtimesng.com/category/health/health-features",
		"https://www.premiumtimesng.com/category/health#",
		"https://www.premiumtimesng.com/category/agriculture/agric-news",
		"https://www.premiumtimesng.com/category/agriculture/agric-special-reports-and-investigations",
		"https://www.premiumtimesng.com/category/agriculture/agric-interviews",
		"https://www.premiumtimesng.com/category/agriculture/agric-multimedia",
		"https://www.premiumtimesng.com/aun-premium-times-data-hub-2#",
		"https://www.premiumtimesng.com/endsars-dashboard-2",
		"https://www.premiumtimesng.com/category/parliament-watch/",
		"https://www.premiumtimesng.com/pandora-papers/",
		"https://www.premiumtimesng.com/category/panama-papers",
		"https://www.premiumtimesng.com/tag/paradise-papers",
		"https://www.premiumtimesng.com/category/agahrin-project",
	})

	// 访问下一页 Index
	w.OnHTML(`.jeg_block_navigation a[class="page_nav next"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 Report & News 从 Index
	w.OnHTML(`[class="jeg_cat_content row"] .jeg_thumb > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), "news/") {
			w.Visit(element.Attr("href"), crawlers.News)
		} else {
			w.Visit(element.Attr("href"), crawlers.Report)
		}
	})

	// 获取 Title
	w.OnHTML(`h1.jeg_post_title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`h2.jeg_post_subtitle`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.entry-header .jeg_meta_date > a:nth-child(1)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`[class="jeg_meta_author coauthor"] > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 CommentCount
	w.OnHTML(`.comment-count`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		var str = strings.Replace(element.Text, "comments", "", 1)
		str = strings.TrimSpace(str)
		num, _ := strconv.Atoi(str)
		ctx.CommentCount = num
	})

	// 获取 Content
	w.OnHTML(`[class="content-inner "]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p, h2, h3"))
	})
}
