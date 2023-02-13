package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("stimson", "史汀森中心东亚项目", "https://www.stimson.org/")

	w.SetStartingURLs([]string{
		`https://www.stimson.org/?ee_search_query=%7B%7D&s=`,
		`https://www.stimson.org/about/people/stimson-staff/`,
	})

	// 访问下一页 Index
	w.OnHTML(`a[class="next page-numbers"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`#main .entry-title > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Expert 从 Index
	w.OnHTML(`[class="ee-post__body ee-post__area"] > a.ee-post__title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 获取 Title
	w.OnHTML(`h1[class="elementor-heading-title elementor-size-default"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Title
	w.OnHTML(`#main > div > section.elementor-section.elementor-top-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div > div > div.elementor-element.elementor-widget.elementor-widget-heading > div > h2`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Name
	w.OnHTML(`#main > div > section.elementor-section.elementor-top-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div > div > div.elementor-element.elementor-widget.elementor-widget-theme-post-title.elementor-page-title.elementor-widget-heading > div > h1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Email & TwitterID
	w.OnHTML(`#main > div > section.elementor-section.elementor-top-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div.elementor-column.elementor-col-50.elementor-top-column.elementor-element.elementor-element-f6cdff5 > div > div.elementor-element.elementor-element-cf76f7c.elementor-widget.elementor-widget-post-info > div > ul > li > a `, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), "mailto:") {
			ctx.Email = strings.TrimSpace(strings.Replace(element.Text, "mailto:", "", 1))
		} else if strings.Contains(element.Attr("href"), "twitter.com") {
			ctx.TwitterID = strings.TrimSpace(strings.Replace(element.Text, "https://twitter.com/", "", 1))
		}
	})

	// 获取 Expert's Phone
	w.OnHTML(`section.elementor-section.elementor-top-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div.elementor-column.elementor-col-50.elementor-top-column.elementor-element > div > div.elementor-element.elementor-widget.elementor-widget-post-info > div > ul > li.elementor-icon-list-item > span.elementor-icon-list-text.elementor-post-info__item.elementor-post-info__item--type-custom`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Phone = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`#main > div > section.elementor-section.elementor-top-section.elementor-element.elementor-section-stretched.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div > div > div > div > div > div > section > div.elementor-container.elementor-column-gap-default > div > div > section.elementor-section.elementor-inner-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default:nth-child(2) > div > div.elementor-column.elementor-col-50.elementor-inner-column.elementor-element  > div> div.elementor-element.elementor-widget.elementor-widget-heading:nth-child(2) > div > div[class="elementor-heading-title elementor-size-default"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Description
	w.OnHTML(`#main > div > section.elementor-section.elementor-top-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div.elementor-column.elementor-col-50.elementor-top-column.elementor-element > div > div.elementor-element.elementor-widget.elementor-widget-theme-post-content > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.ChildText("p"))
	})

	// 获取 Description
	w.OnHTML(`#main > div > section.elementor-section.elementor-top-section.elementor-element.elementor-section-stretched.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div > div > div > div > div > div > section.elementor-section.elementor-top-section.elementor-element.elementor-section-height-min-height.elementor-section-items-stretch.elementor-section-stretched.elementor-section-boxed.elementor-section-height-default > div > div.elementor-column.elementor-col-50.elementor-top-column > div > section > div > div > div > div.elementor-element.elementor-widget.elementor-widget-heading > div > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`#main > div > section.elementor-section.elementor-top-section.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div.elementor-column.elementor-col-50.elementor-top-column.elementor-element.elementor-element-1dfa510a > div > div.elementor-element.elementor-align-left.elementor-widget__width-auto.elementor-widget.elementor-widget-post-info > div > ul > li.elementor-icon-list-item.elementor-repeater-item-6f91969 > span.elementor-icon-list-text.elementor-post-info__item.elementor-post-info__item--type-custom`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`section.elementor-section.elementor-top-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div.elementor-column.elementor-col-50.elementor-top-column.elementor-element:nth-child(2) > div > div > div > ul > li.elementor-icon-list-item.elementor-inline-item:nth-child(1) > span.elementor-icon-list-text.elementor-post-info__item.elementor-post-info__item--type-custom`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`section.elementor-section.elementor-top-section.elementor-element.elementor-section-stretched.elementor-section-content-space-between.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div > div > section.elementor-section.elementor-inner-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div.elementor-column.elementor-col-50.elementor-inner-column.elementor-element > div > div.elementor-element.elementor-widget.elementor-widget-heading > div > div > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`section.elementor-section.elementor-top-section.elementor-element.elementor-section-stretched.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div > div > div > div > div > div > section > div.elementor-container.elementor-column-gap-default > div > div > section.elementor-section.elementor-inner-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div.elementor-column.elementor-col-50.elementor-inner-column.elementor-element > div > div.elementor-element.elementor-widget.elementor-widget-heading > div > div > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`section.elementor-section.elementor-top-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div.elementor-column.elementor-col-50.elementor-top-column.elementor-element > div > div > div > div > div > p:nth-child(1) > strong`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Authors
	w.OnHTML(`#main > div > section.elementor-section.elementor-top-section.elementor-element.elementor-section-stretched.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div > div > div > div > div > div > section > div > div.elementor-column.elementor-col-50.elementor-top-column.elementor-element > div > div.elementor-element.elementor-widget.elementor-widget-template > div > div > div > section > div > div > div > div.elementor-element.elementor-widget.elementor-widget-toolset-view > div > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 PublicationTime
	w.OnHTML(`section.elementor-section.elementor-top-section.elementor-element.elementor-section-stretched.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div > div > div > div > div > div > section > div > div.elementor-column.elementor-col-50.elementor-top-column.elementor-element > div > div.elementor-element.elementor-icon-list--layout-traditional.elementor-list-item-link-full_width.elementor-widget.elementor-widget-icon-list > div > ul > li > span.elementor-icon-list-text`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.elementor-post-info__item.elementor-post-info__item--type-date`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`[class="elementor-widget-wrap elementor-element-populated"] .elementor-widget-container`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p, h1, h2, h3, h4"))
	})

	// 获取 Content
	w.OnHTML(`#scroll-indic-box > div > div.elementor-column.elementor-col-50.elementor-top-column.elementor-element > div > div > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p, ol, ul, h2"))
	})

	// 获取 Content
	w.OnHTML(`div[data-widget_type="theme-post-content.default"] > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p, ol, ul, h2"))
	})

	// 获取 Location
	w.OnHTML(`#main > div > section.elementor-section.elementor-top-section.elementor-element.elementor-section-boxed.elementor-section-height-default.elementor-section-height-default > div > div.elementor-column.elementor-col-50.elementor-top-column.elementor-element > div > div.elementor-element.elementor-align-left.elementor-widget__width-auto.elementor-widget.elementor-widget-post-info > div > ul > li.elementor-icon-list-item:nth-child(3) > span.elementor-icon-list-text.elementor-post-info__item.elementor-post-info__item--type-custom`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Location = strings.TrimSpace(element.Text)
	})

	// 获取 Location
	w.OnHTML(`div.elementor-column.elementor-col-50.elementor-top-column.elementor-element:nth-child(2) > div > div > div > ul > li.elementor-icon-list-item.elementor-inline-item:nth-child(3) > span.elementor-icon-list-text.elementor-post-info__item.elementor-post-info__item--type-custom`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Location = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`.elementor-post-info__terms-list-item`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}
