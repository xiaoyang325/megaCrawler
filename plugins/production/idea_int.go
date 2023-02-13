package production

import (
	"megaCrawler/crawlers"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("idea_int", "国际民主与选举援助研究所", "https://www.idea.int/")

	w.SetStartingUrls([]string{
		"https://www.idea.int/publications/catalogue",
		"https://www.idea.int/news-media/news",
		"https://www.idea.int/blog",
		"https://www.idea.int/news-media/event-archive",
		"https://www.idea.int/news-media/media",
		"https://www.idea.int/news-media/multimedia-reports",
	})

	// 访问下一页 Index
	w.OnHTML(`.view-header .next > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问下一页 Index
	w.OnHTML(`.text-center .next > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.view-content [class="views-field views-field-php"] a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Report 从 Index
	w.OnHTML(`[class="views-field views-field-title"]  .field-content > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Report 从 Index
	w.OnHTML(`[class="views-field views-field-title-1"]  .field-content > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.col-12 > div .description-holder > .blog-link > a `, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title
	w.OnHTML(`.page-header div[class="field-item even"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title
	w.OnHTML(`h3.blog-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title
	w.OnHTML(`h3.news-title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`.subtitle div[class="field-item even"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`[class="field field-name-field-news-date field-type-datetime field-label-inline clearfix"] .date-display-single`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`[class="field field-name-field-blog-date field-type-datetime field-label-hidden"] .date-display-single`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`div.pubdata > div.field.field-name-field-news-date.field-type-datetime.field-label-hidden > div > div > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`div.row > div.col-sm-6.eventdata > div.field.field-name-event-calendar-date.field-type-datetime.field-label-hidden > div > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`.easy-breadcrumb > span:nth-child(5) > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`h1.page-header`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
		if ctx.CategoryText == "News" {
			ctx.PageType = crawlers.News
		}
	})

	// 获取 Authors
	w.OnHTML(`[class="field field-name-field-author field-type-text field-label-inline clearfix"] [class="field-item even"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		authorList := crawlers.SplitDelimiters(element.Text, []string{"&", "and"})
		ctx.Authors = append(ctx.Authors, authorList...)
	})

	// 获取 Authors
	w.OnHTML(`.autor-full-name`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		authorList := crawlers.SplitDelimiters(element.Text, []string{"&", "and"})
		ctx.Authors = append(ctx.Authors, authorList...)
	})

	// 获取 Authors
	w.OnHTML(`#authorbox > .info > .name`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		authorList := crawlers.SplitDelimiters(element.Text, []string{"&", "and"})
		ctx.Authors = append(ctx.Authors, authorList...)
	})

	// 获取 ViewCount
	w.OnHTML(`.row .views [class="field-item even"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		var str = strings.TrimSpace(element.Text)
		str = strings.Replace(str, ",", "", -1)
		num, _ := strconv.Atoi(str)
		ctx.ViewCount = num
	})

	// 获取 Description
	w.OnHTML(`[class="field field-name-field-description field-type-text-long field-label-hidden"] [class="field-item even"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`.content > .blog-main-content`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`[class="field field-name-body field-type-text-with-summary field-label-hidden"] > div > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Language
	w.OnHTML(`[class="field field-name-field-language field-type-list-text field-label-inline clearfix"] [class="field-item even"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Language = strings.TrimSpace(element.Text)
	})

	// 获取 Location
	w.OnHTML(`[class="field field-name-field-location field-type-text field-label-hidden"] [class="field-item even"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Location = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`div.pubdata > div.field.field-type-entityreference.field-label-hidden > div > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`.field-button > a[target="_blank"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}
