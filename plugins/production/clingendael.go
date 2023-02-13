package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("clingendael", "国际关系研究所", "https://www.clingendael.org/")

	w.SetStartingUrls([]string{"https://www.clingendael.org/publications",
		"https://www.clingendael.org/topic/security-and-justice",
		"https://www.clingendael.org/topic/business-and-fragility",
		"https://www.clingendael.org/topic/aid-architecture",
		"https://www.clingendael.org/topic/defence",
		"https://www.clingendael.org/topic/terrorism",
		"https://www.clingendael.org/topic/cyber-security",
		"https://www.clingendael.org/topic/cbrn",
		"https://www.clingendael.org/topic/transatlantic",
		"https://www.clingendael.org/topic/europes-east",
		"https://www.clingendael.org/topic/asia",
		"https://www.clingendael.org/research-program/geopolitics-technology-and-digitalisation",
		"https://www.clingendael.org/topic/eu-migration",
		"https://www.clingendael.org/nl/topic/eu-integration",
		"https://www.clingendael.org/nl/topic/social-europe",
		"https://www.clingendael.org/topic/support-eu",
		"https://www.clingendael.org/topic/brexit",
		"https://www.clingendael.org/topic/new-silk-road",
		"https://www.clingendael.org/topic/geostrategic-risk-management",
		"https://www.clingendael.org/topic/strategic-monitor",
		"https://www.clingendael.org/topic/global-security-pulse",
		"https://www.clingendael.org/topic/scenarios",
		"https://www.clingendael.org/topic/sustainability",
		"https://www.clingendael.org/research-program/coronavirus",
		"https://www.clingendael.org/research-program/china-centre",
		"https://www.clingendael.org/research-program/russia-eastern-europe-centre",
		"https://www.clingendael.org/research-program/foreign-affairs-barometer",
		"https://www.clingendael.org/publications?regions=37",
		"https://www.clingendael.org/publications?regions=20",
		"https://www.clingendael.org/publications?regions=19",
		"https://www.clingendael.org/publications?regions=40",
		"https://www.clingendael.org/publications?regions=39",
		"https://www.clingendael.org/publications?regions=38",
		"https://www.clingendael.org/publications?regions=36",
		"https://www.clingendael.org/events"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// 从翻页器获取链接并访问
	w.OnHTML("a.show-more-results", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a.show-more-results", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	w.OnHTML("a.link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	// 从index访问新闻
	w.OnHTML("h3.title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.highlight-info>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.publication>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("div.node-info>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	// report.title
	w.OnHTML(".block-title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// report.publish time
	w.OnHTML("div.block-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("div.staff-person-name>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// 内含Expert
	w.OnHTML("div.staff-person-name>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})
	// report .content
	w.OnHTML("div.publication-content, .event-content-wrapper, .training-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	// report.category
	w.OnHTML("div.type-and-topic", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = element.Text
	})

	// expert.Name
	w.OnHTML("div.expert-details>h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = element.Text
	})
	// expert.title
	w.OnHTML(".function", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	// expert.description
	w.OnHTML("div.expert-info > div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	// expert.area
	w.OnHTML("div.expertise>ul>li", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Area = ctx.Area + "," + element.Text
	})
	// expert.link
	w.OnHTML("div.contact-icons>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
}
