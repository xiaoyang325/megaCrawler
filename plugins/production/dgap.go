package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("dgap", "外交关系协会", "https://dgap.org/")

	w.SetStartingURLs([]string{
		"https://dgap.org/en/research/expertise/geo-economics/g7",
		"https://dgap.org/en/research/expertise/geo-economics/resources-and-energy",
		"https://dgap.org/en/research/expertise/geo-economics/trade",
		"https://dgap.org/en/research/expertise/dossier-russias-war-against-ukraine",
		"https://dgap.org/en/dossier/german-elections-2021",
		"https://dgap.org/en/research/expertise/european-union",
		"https://dgap.org/en/research/expertise/international-order-democracy",
		"https://dgap.org/en/research/expertise/security",
		"https://dgap.org/en/research/expertise/technology-digitization",
		"https://dgap.org/en/research/expertise/migration",
		"https://dgap.org/en/research/expertise/zeitenwende",
		"https://dgap.org/en/research/expertise/climate",
		"https://dgap.org/en/research/expertise/africa",
		"https://dgap.org/en/research/expertise/americas",
		"https://dgap.org/en/research/expertise/asia",
		"https://dgap.org/en/research/expertise/europe",
		"https://dgap.org/en/research/expertise/europe/eastern-europe",
		"https://dgap.org/en/research/expertise/europe/france",
		"https://dgap.org/en/research/expertise/middle-east-north-africa",
		"https://dgap.org/en/research/expertise/russia-central-asia",
		"https://dgap.org/en/publications",
		"https://dgap.org/en/events",
		"https://dgap.org/en/research/programs/international-order-and-democracy-program",
		"https://dgap.org/en/research/programs/center-climate-and-foreign-policy",
		"https://dgap.org/en/research/programs/alfred-von-oppenheim-center-european-policy-studies",
		"https://dgap.org/en/research/programs/asia-program",
		"https://dgap.org/en/research/programs/geo-economics-program",
		"https://dgap.org/en/research/programs/impact-innovation-lab",
		"https://dgap.org/en/research/programs/migration-program",
		"https://dgap.org/en/research/programs/security-and-defense-program",
		"https://dgap.org/en/research/programs/technology-and-global-affairs-program",
	})

	// 访问下一页 Index
	w.OnHTML(`#Recent-publications > div > div > div > div:nth-child(2) > div > div > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问下一页 Index
	w.OnHTML(`#Events > div > div > div > div > div > div > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问下一页 Index
	w.OnHTML(`#Archive > div > div > div > div > div > div > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问下一页 Index
	w.OnHTML(`#block-dgap-content > article > div > div > div > div > div > div > div > div > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`#Recent-publications > div > div > div > div:nth-child(2) > div > div > div > div > div > article > h3 > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Report 从 Index
	w.OnHTML(`#block-dgap-content > article > div > div > div > div > div > div > div > div > div.views-infinite-scroll-content-wrapper.clearfix > div > div > article > h3 > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Expert 从 Index
	w.OnHTML(`#Experts article div.user--content div.field.field--name-realname.field--type-text.field--label-inline > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 访问 Report 从 Index
	w.OnHTML(`#Events > div > div > div > div > div > div > div.views-infinite-scroll-content-wrapper.clearfix > div > div > article > h3 > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 访问 Report 从 Index
	w.OnHTML(`#Archive > div > div > div > div > div > div > div.views-infinite-scroll-content-wrapper.clearfix > div > div > article > h3 > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title
	w.OnHTML(`#block-dgap-content > article div.block.block-layout-builder.block-field-blocknodedgap-articletitle.block-field-block > h1 > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title
	w.OnHTML(`#block-dgappagetitle > div > div > div > div > div > div > div > h1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`#block-dgap-content > article div.block.block-layout-builder.block-field-blocknodedgap-articlefield-subheadline.block-field-block > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`#block-dgap-content > article > div > div:nth-child(1) > div > div > div > div > div > p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`#block-dgap-content article div.layout--twocol-section.layout div.no-offset-widescreen.is-9-widescreen.is-9-tablet.column > div.block.block-layout-builder> div > p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`#block-dgappagetitle > div > div > div > div > footer > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`#block-dgap-content > article > div > div:nth-child(2) > div > div > div > div.fields-inline.show-icon-calendar-clock.block.block-layout-builder.block-field-blocknodedgap-eventfield-date-range.block-field-block > div > div.field__item > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`#block-dgappagetitle > div > div > div > div > footer > h4`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`#block-dgappagetitle > div > div > div > div > div > div > div > nav > ol > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`#block-dgap-content article div.block.block-layout-builder.block-field-blocknodedgap-articlefield-authors.block-field-block div.field__item > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Authors
	w.OnHTML(`#block-dgap-content > article > div > div.layout--twocol-section.layout--twocol-section--is-6\:is-3.layout > div > div > div.is-offset-1-widescreen.is-3-widescreen.is-4-desktop.column > div.block.block-layout-builder.block-field-blocknodedgap-eventfield-contact.block-field-block > div > div.field__items > div > article > div.user--content > div.field.field--name-realname.field--type-text.field--label-inline > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Authors
	w.OnHTML(`#block-dgap-content > article > div > div.layout--twocol-section.layout--twocol-section--is-6\:is-3.layout > div > div > div.is-offset-1-widescreen.is-6-widescreen.is-8-desktop.column > div.views-element-container.block.block-views.block-views-blockdgap-referenced-experts-experts-on-current-node article div.user--content > div.field.field--name-realname.field--type-text.field--label-inline > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`article div.layout--twocol-section.layout div.is-offset-1-widescreen.is-6-widescreen.is-8-desktop.column div.block.block-layout-builder.block-field-blocknodedgap-articlebody > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`#block-dgap-content > article > div > div.layout--twocol-section.layout--twocol-section--is-6\:is-3.layout > div > div > div.is-offset-1-widescreen.is-6-widescreen.is-8-desktop.column > div.block.block-layout-builder.block-field-blocknodedgap-eventbody.block-field-block > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`#block-dgap-content article div.layout--twocol-section.layout--twocol-section--is-6\:is-3.layout div.is-offset-1-widescreen.is-3-widescreen.is-4-desktop.column > div.arrow-list.block.block-dgap-base.block-dgap-core-topics-and-regions > div > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`article div.layout--twocol-section div.no-offset-widescreen.is-2-widescreen.is-3-tablet.is-2-widescreen.column div.block.block-layout-builder.block-field-blocknodedgap-articlefield-pdf-download.block-field-block > div > div.field__item > article div.field.field--name-field-media-file.field--type-file.field--label-hidden.field__item > span > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := "https://dgap.org/" + element.Attr("href")
		ctx.File = append(ctx.File, fileURL)
	})

	// 获取 Location
	w.OnHTML(`#block-dgap-content > article div.fields-inline.show-icon-location.block.block-layout-builder.block-field-blocknodedgap-eventfield-event-location.block-field-block > div > div.field__item > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Location = element.Attr("href")
	})

	// 获取 Name
	w.OnHTML(`#block-dgap-content > article > div.layout--twocol-section.layout--twocol-section--is-6\:is-3.layout > div > div > div.is-offset-1-widescreen.is-3-widescreen.is-4-desktop.column > div.block.block-layout-builder.block-field-blockuserusername.block-field-block > h1 > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = strings.TrimSpace(element.Text)
	})

	// 获取 Title
	w.OnHTML(`#block-dgap-content > article > div.layout--twocol-section.layout--twocol-section--is-6\:is-3.layout > div > div > div.is-offset-1-widescreen.is-3-widescreen.is-4-desktop.column  div.field.field--name-field-position-en.field--type-link.field--label-hidden.field__item > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Phone
	w.OnHTML(`#block-dgap-content > article > div.layout--twocol-section.layout--twocol-section--is-6\:is-3.layout > div > div > div.is-offset-1-widescreen.is-3-widescreen.is-4-desktop.column > div:nth-child(3) > div > div.field.field--name-field-phone-business.field--type-string.field--label-inline > div.field__item`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Phone = strings.TrimSpace(element.Text)
	})

	// 获取 Email
	w.OnHTML(`#block-dgap-content > article > div.layout--twocol-section.layout--twocol-section--is-6\:is-3.layout div.is-offset-1-widescreen.is-3-widescreen.is-4-desktop.column div.field.field--name-dgap-mail.field--type-text.field--label-inline > div.field__item > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Email = strings.TrimSpace(element.Text)
	})

	// 获取 TwitterID
	w.OnHTML(`#block-dgap-content > article > div.layout--twocol-section.layout--twocol-section--is-6\:is-3.layout div.is-offset-1-widescreen.is-3-widescreen.is-4-desktop.column div.field.field--name-field-twitter.field--type-link.field--label-inline > div.field__item > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.TwitterID = strings.TrimSpace(strings.Replace(element.Text, "@", "", 1))
	})

	// 获取 Description
	w.OnHTML(`#block-dgap-content > article > div.layout--twocol-section.layout--twocol-section--is-6\:is-3.layout div.is-offset-1-widescreen.is-6-widescreen.is-8-desktop.column > div > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})
}
