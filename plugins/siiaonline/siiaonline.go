package siiaonline

import (
	"github.com/araddon/dateparse"
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"regexp"
	"strings"
	"time"
)

func init() {
	w := Crawler.Register("siiaonline", "新加坡国际事务学院", "http://www.siiaonline.org/")
	w.SetStartingUrls([]string{"http://www.siiaonline.org/our-people/",
		"http://www.siiaonline.org/reports-index/"})

	//人物信息
	w.OnHTML("article.mix", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = Crawler.Expert
		subCtx.Name = element.ChildText(".portfolio_title")
		subCtx.Title = element.ChildText("span.project_category")

		match := regexp.MustCompile("paoc-popup-cust-(\\d+)").FindStringSubmatch(
			element.ChildAttr(".paoc-popup-click", "class"),
		)
		if len(match) > 1 {
			subCtx.Description = strings.TrimSpace(element.DOM.Find("paoc-popup-" + match[1]).Text())
		}
	})

	//访问报告
	w.OnHTML("div.vc_btn3-container.vc_btn3-left > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//报告标题
	w.OnHTML(".title_subtitle_holder > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//报告标签
	w.OnHTML(".tags_text>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, element.Text)
	})

	//pdf
	w.OnHTML("div.post_text > div > p > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})

	//正文
	w.OnHTML("div.post_text_inner", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	w.OnHTML(".post_author_link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	w.OnHTML(".post_text > div > p:nth-child(4)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		matches := regexp.MustCompile("Date :(.+)").FindStringSubmatch(element.Text)
		if len(matches) > 1 {
			parseAny, err := dateparse.ParseAny(matches[1])
			if err != nil {
				Crawler.Sugar.Error(err)
				return
			}
			ctx.PublicationTime = parseAny.Format(time.RFC3339)
		}
	})
}
