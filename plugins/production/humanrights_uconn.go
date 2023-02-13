package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("humanrights_uconn", "国防部人权中心",
		"https://uconn.edu/")

	w.SetStartingUrls([]string{
		"https://today.uconn.edu/archives/",
		"https://humanrights.uconn.edu/leadership-staff/",
		"https://humanrights.uconn.edu/about/our-people/faculty/",
		"https://humanrights.uconn.edu/the-gladstein-committee/",
		"https://teachingdatabase.humanrights.uconn.edu/",
	})

	// 访问下一页 Index
	w.OnHTML(`a[class="next page-numbers"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 News 从 Index
	w.OnHTML(`[class="archive-list multiple-archives"] a.small-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.entry-header > .entry-title > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Expert 从 Index
	w.OnHTML(`.person .person-name > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 获取 Title
	w.OnHTML(`.entry-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Title
	w.OnHTML(`#main div > p:nth-child(2)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Description
	w.OnHTML(`#main  div > p:nth-child(5)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Email, Phone, Location
	w.OnHTML(`.table-responsive > table.table > tbody`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Email = element.ChildText("td.person-email > a")
		ctx.Phone = strings.TrimSpace(element.ChildText("td.person-phone"))
		ctx.Location = element.ChildText("td.person-email > a")
		ctx.ExpertWebsite = element.ChildAttr("td.person-email > a", "href")
	})

	// 获取 Name
	w.OnHTML(`#main h1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`.post-excerpt-contain`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.contain .byline-date`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`.contain .byline-author`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = crawlers.Unique(append(ctx.Authors, strings.TrimSpace(element.Text)))
	})

	// 获取 Content
	w.OnHTML(`.entry-content`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p, h1, h2, h3, h4"))
	})

	// 获取 Tags
	w.OnHTML(`[class="category-tag "] > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}
