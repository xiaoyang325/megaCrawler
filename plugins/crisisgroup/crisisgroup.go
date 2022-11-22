package crisisgroup

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"net/url"
	"strconv"
	"strings"
)

// 这个函数修改当前 Index 页面的查询参数，以获取下一页 Index，并返回相应的 URL
func getNextIndexURL(currentUrl string, currentPageNum string,
	paramName string) string {
	thisUrl, _ := url.Parse(currentUrl)
	paramList := thisUrl.Query()

	currentNum, _ := strconv.Atoi(currentPageNum)
	currentNum++

	paramList.Set(paramName, strconv.Itoa(currentNum))
	thisUrl.RawQuery = paramList.Encode()

	return thisUrl.String()
}

func init() {
	w := Crawler.Register("crisisgroup", "国际危机组织", "https://www.crisisgroup.org/")

	w.SetStartingUrls([]string{
		"https://www.crisisgroup.org/latest-updates?page=0",
		"https://www.crisisgroup.org/who-we-are/our-people",
	})

	// 访问下一页 Index //
	w.OnHTML(`body > div.dialog-off-canvas-main-canvas > main > div > div.s-component.c-news-list.o-container.o-container--cols1 > div.o-pagination-container.\[.u-df.u-jcc.\] > ul > li.u-mar-l15.u-mar-r15.u-mar-l25\@m.u-mar-r25\@m > span`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			numList := strings.Split(element.Text, "of")
			currentNum, _ := strconv.Atoi(strings.TrimSpace(numList[0]))
			maxNum, _ := strconv.Atoi(strings.TrimSpace(numList[1]))
			if currentNum <= maxNum {
				w.Visit(getNextIndexURL(ctx.Url, strings.TrimSpace(numList[0]), "page"), Crawler.Index)
			}
		})

	// 访问 Report 从 Index //
	w.OnHTML(`div.s-component.c-news-list.o-container.o-container--cols1 div.c-news-listing__content.\[.u-df.u-flexdc.\] > h4 > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 访问 Expert 从 Index //
	w.OnHTML(`body > div.dialog-off-canvas-main-canvas > main > div > div > div > div.c-our-people.\[.u-df\@m.u-flexdc.u-flexdr\@m.u-flexww.u-jcfs.u-pad-b40\@m.u-mar-t50\@m.\] > div > div.c-media__img.o-ar-1x1.o-image.o-image--cover.\[.u-ofh.u-pr.\].\[.u-mar-la\@l.u-mar-ra\@l.\].u-z-1 > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Expert)
		})

	// 获取 Title //
	w.OnHTML(`body > div.dialog-off-canvas-main-canvas > main > div > div.s-wrapper.u-display-flex.u-bg-white.u-z-1 > article > div.o-container > div.c-page-hero.c-page-hero--vd.js-page-hero > div.c-page-hero__details.\[.u-pr.\] > h1`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Description //
	w.OnHTML(`div[class="c-page-hero__details [ u-pr ]"] div[class="c-page-hero__teaser [ u-ptserif u-fs18 u-fs15@m u-fs18@l u-fsi ]"] > div > p[style="font-weight: 400;"]:nth-child(1)`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Description = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime //
	w.OnHTML(`div[class="c-page-hero__details [ u-pr ]"] time`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 CategoryText //
	w.OnHTML(`div[class="c-page-hero__details [ u-pr ]"] > div > a >strong`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 Authors //
	w.OnHTML(`div[class="c-media--contributors  u-df u-flexdr u-flexdc@m u-mar-b50"] div[class="c-media__title [ u-fwn u-fwb@m u-fwn@l ]"] > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 Content //
	w.OnHTML(`div[class="s-article__body u-pos-relative"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = element.Text
		})

	// 获取 Tags //
	w.OnHTML(`.s-article__sidebar div.s-list > .u-ttu > small`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
		})

	// 获取 File //
	w.OnHTML(`body > div.dialog-off-canvas-main-canvas > main > div > div.s-wrapper.u-display-flex.u-bg-white.u-z-1 > article > div.o-container > div.c-toolbar.c-toolbar--lang.js-toolbar > div > ul > li.u-ofh > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.File = append(ctx.File, element.Attr("href"))
		})

	// 获取 Name //
	w.OnHTML(`body > div.dialog-off-canvas-main-canvas > main > div > article > div.o-container.o-container--m > header > div > div.\[.u-tac.u-tal\@m.\].\[.u-mar-t15.u-mar-t0\@m.\] > h1`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Name = strings.TrimSpace(element.Text)
		})

	// 获取 Title //
	w.OnHTML(`body > div.dialog-off-canvas-main-canvas > main > div > article > div.o-container.o-container--m > header > div > div.\[.u-tac.u-tal\@m.\].\[.u-mar-t15.u-mar-t0\@m.\] > div > div:nth-child(1)`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 TwitterId //
	w.OnHTML(`ul > li > a[title="Twitter"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			id := strings.Replace(element.Attr("href"), "https://twitter.com/", "", 1)
			ctx.TwitterId = strings.TrimSpace(id)
		})

	// 获取 Description //
	w.OnHTML(`div[class="s-article__body s-copy u-pos-relative"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Description = strings.TrimSpace(element.Text)
		})

	// 获取 Email //
	w.OnHTML(`ul > li > a[title="Email"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			address := strings.Replace(element.Attr("href"), "mailto:", "", 1)
			ctx.Email = strings.TrimSpace(address)
		})

	// 获取 Location //
	w.OnHTML(`body > div.dialog-off-canvas-main-canvas > main > div > div.s-wrapper.u-display-flex.u-bg-white.u-z-1 > article > div.o-container > div.s-article__main.s-copy.u-display-flex.u-flexdc.u-flexdr\@m.u-flexww\@m > div > div`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Location = strings.TrimSpace(element.Text)
		})

	// 获取 LinkedInId //
	w.OnHTML(`ul > li > a[title="Linkedin"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			address := strings.Replace(element.Attr("href"), "https://www.linkedin.com/in/", "", 1)
			ctx.LinkedInId = strings.TrimSpace(address)
		})
}
