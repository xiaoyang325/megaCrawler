package production

import (
	"megaCrawler/crawlers"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("csis", "战略与国际研究中心", "https://www.csis.org/")

	w.SetStartingURLs([]string{"https://www.csis.org/experts", "https://www.csis.org/analysis"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// 从翻页器获取链接并访问
	w.OnHTML(".pager__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 获取分类到ctx
	w.OnHTML(".page-type", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = element.Text
	})

	// 尝试访问作者并添加到ctx
	w.OnHTML(".teaser--portrait-image > .teaser__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if element.ChildAttr("a", "href") != "" {
			w.Visit(element.ChildAttr("a", "href"), crawlers.Expert)
		}
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// 从index访问新闻
	w.OnHTML(".teaser__image > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 添加标签到ctx
	w.OnHTML(".field--spaced > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	// 添加标题到ctx
	w.OnHTML("h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if ctx.PageType == crawlers.Expert {
			ctx.Name = element.Text
		} else if ctx.PageType == crawlers.Report || ctx.PageType == crawlers.News {
			ctx.Title = element.Text
		}
	})

	// 添加正文到ctx
	w.OnHTML("article[role=\"article\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 人物头衔到ctx
	w.OnHTML(".layout-constrain > .layout-focus-page__main > .subtitle", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = crawlers.StandardizeSpaces(element.Text)
	})

	// 人物描述到ctx
	w.OnHTML(".layout-constrain > .layout-focus-page__main > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})

	// 人物地区到ctx
	w.OnHTML(".layout-focus-page__main > .field > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Area += element.Text + ", "
	})

	// 正则匹配邮箱和电话号码
	w.OnHTML("div.pane.pane--csis-contributor-contact-expert.pane--details > div.pane__content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		emailRegex := regexp.MustCompile(`Email: ([.@\w]+)`)
		telRegex := regexp.MustCompile(`Tel: ([.\w]+)`)
		emailMatch := emailRegex.FindStringSubmatch(element.Text)
		telMatch := telRegex.FindStringSubmatch(element.Text)
		if len(emailMatch) == 2 {
			ctx.Email = emailMatch[1]
		}
		if len(telMatch) == 2 {
			ctx.Phone = telMatch[1]
		}
	})

	w.OnHTML(".nav__link--linkedin", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.LinkedInID = element.Attr("href")
	})

	w.OnHTML(".nav__link--twitter", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.TwitterID = element.Attr("href")
	})
}
