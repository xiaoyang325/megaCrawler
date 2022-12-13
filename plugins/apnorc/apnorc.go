package apnorc

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("apnorc", "Associated Press-NORC Center for Public Affairs Research",
		"https://apnorc.org/")

	w.SetStartingUrls([]string{
		"https://apnorc.org/experts/",
		"https://apnorc.org/topics/economics/",
		"https://apnorc.org/topics/politics/",
		"https://apnorc.org/topics/healthcare/",
		"https://apnorc.org/topics/environment-energy-and-resilience/",
		"https://apnorc.org/topics/news-and-media/",
		"https://apnorc.org/topics/race-ethnicity-and-social-inequality/",
	})

	// 访问下一页 Index
	w.OnHTML(`a[class="next page-numbers"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`header > h2 > a[rel="bookmark"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`header > div > div > h1[class="entry-title heading--single-project"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`div.entry-summary > p`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`div[class="row align-center"] > div > div > p:nth-child(2) > em`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`header > div > div > div.term-name > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`div[class="entry-content clearfix"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	w.OnHTML(".experts-list", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = Crawler.Expert
		subCtx.Title = element.ChildText(".designation")
		subCtx.Area = element.ChildText(".organisation")
		subCtx.Email = strings.TrimSpace(element.ChildText(".email-address"))
		subCtx.Phone = element.ChildText(".contact-no")
		subCtx.Name = element.ChildText(".entry-title")
		subCtx.Image = []string{element.ChildAttr(".has-thumbnail > img", "src")}
		subCtx.Description = element.ChildText(".content")
	})
}
