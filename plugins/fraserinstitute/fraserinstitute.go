package fraserinstitute

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"megaCrawler/Extractors"
	"strings"
)

func init() {
	w := Crawler.Register("fraserinstitute", "菲沙研究所", "https://www.fraserinstitute.org/")

	w.SetStartingUrls([]string{"https://www.fraserinstitute.org/studies/aboriginal-policy",
		"https://www.fraserinstitute.org/studies/pensions-retirement",
		"https://www.fraserinstitute.org/studies/competitiveness",
		"https://www.fraserinstitute.org/studies/covid",
		"https://www.fraserinstitute.org/studies/economic-freedom",
		"https://www.fraserinstitute.org/studies/education",
		"https://www.fraserinstitute.org/categories/energy",
		"https://www.fraserinstitute.org/studies/environment",
		"https://www.fraserinstitute.org/studies/esg",
		"https://www.fraserinstitute.org/studies/government-spending-taxes",
		"https://www.fraserinstitute.org/studies/health-care",
		"https://www.fraserinstitute.org/studies/labour-policy",
		"https://www.fraserinstitute.org/categories/monetary-policy-banking",
		"https://www.fraserinstitute.org/studies/municipal-policy",
		"https://www.fraserinstitute.org/studies/natural-resources",
		"https://www.fraserinstitute.org/studies/poverty-inequality",
		"https://www.fraserinstitute.org/studies/provincial-prosperity",
		"https://www.fraserinstitute.org/studies/school-report-cards",
		"https://www.fraserinstitute.org/studies/other-topics",
		"https://www.fraserinstitute.org/more-from-the-fraser-institute",
		"https://www.fraserinstitute.org/blogs/category/aboriginal-policy",
		"https://www.fraserinstitute.org/blogs/category/atlantic-canada-prosperity",
		"https://www.fraserinstitute.org/blogs/category/covid",
		"https://www.fraserinstitute.org/blogs/category/economic-freedom",
		"https://www.fraserinstitute.org/blogs/category/education",
		"https://www.fraserinstitute.org/blogs/category/environment",
		"https://www.fraserinstitute.org/blogs/category/government-spending-taxes",
		"https://www.fraserinstitute.org/blogs/category/health-care",
		"https://www.fraserinstitute.org/blogs/category/mmt",
		"https://www.fraserinstitute.org/blogs/category/natural-resources",
		"https://www.fraserinstitute.org/blogs/category/other-topics"},
	)

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})

	w.OnHTML("html", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		Extractors.Authors(ctx, element)
		Extractors.PublishingDate(ctx, element)
	})

	// 从翻页器获取链接并访问
	w.OnHTML("div.archive-links>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从翻页器获取链接并访问
	w.OnHTML("div.text-center>ul>li>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML("span.field-content>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	// report.title
	w.OnHTML("h1.page-header", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("div.tab-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	// report .content
	w.OnHTML("div.field-name-body", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	// report.author
	w.OnHTML("div.node-author-fullname>a, div.field-content-author-authors>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//内含Expert
	w.OnHTML("div.node-author-fullname>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	// expert.Name
	w.OnHTML("div.main-author-content > h2 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	// expert.title
	w.OnHTML(" div.main-author-content > h3 > div > div > div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// expert.description
	w.OnHTML(" div.main-author-content > div > div > div > div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	// expert.link
	w.OnHTML(" div.field-name-field-email > div > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
}
