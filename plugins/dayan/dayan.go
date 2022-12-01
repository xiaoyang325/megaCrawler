package dayan

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("dayan", "摩西·达扬中东和非洲研究中心", "https://dayan.org/")
	
	w.SetStartingUrls([]string{
		"https://dayan.org/subject/archives",
		"https://dayan.org/journal/tel-aviv-notes-contemporary-middle-east-analysis",
		"https://dayan.org/he/journal/328",
		"https://dayan.org/journal/bayan-arabs-israel",
		"https://dayan.org/journal/beehive-middle-east-social-media",
		"https://dayan.org/journal/ifriqiya-africa-research-and-analysis",
		"https://dayan.org/journal/iqtisadi-middle-east-economy",
		"https://dayan.org/journal/turkeyscope-insights-turkish-affairs",
		"https://dayan.org/journal/surveys-konrad-adenauer-program",
		"https://dayan.org/journal/jihadiscope",
		"https://dayan.org/journal/scholarly-journals",
		"https://dayan.org/journal/external-publications",
		"https://dayan.org/subject/mdc-videos-and-media",
	})

	// 访问下一页 Index 
	w.OnHTML(`li[class="pager__item pager__item--next"] > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(ctx.Url + element.Attr("href"), Crawler.Index)
		})

	// 访问 Report 从 Index 
	w.OnHTML(`.views-row a.title`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 获取 Title 
	w.OnHTML(`h1.page-title`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Description 
	w.OnHTML(`.top-area > [class="field field--name-field-introduction field--type-string-long field--label-hidden field__item"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Description = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime 
	w.OnHTML(`.field__item > time`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 Authors 
	w.OnHTML(`.top-area .field__item > a[tuafontsizes="14"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 Content 
	w.OnHTML(`[class="clearfix text-formatted field field--name-body field--type-text-with-summary field--label-hidden field__item"] `,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.ChildText("p"))
		})

	// 获取 Tags 
	w.OnHTML(`[class="field field--name-field-subject field--type-entity-reference field--label-hidden field__items"] a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
		})

	// 获取 File 
	w.OnHTML(`[class="file file--mime-application-pdf file--application-pdf"] > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			file_url := "https://dayan.org" + element.Attr("href")
			ctx.File = append(ctx.File, file_url)
		})
}
