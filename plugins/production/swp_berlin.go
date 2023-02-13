package production

import (
	"megaCrawler/crawlers"
	"net/url"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

// 将 URL 的路劲加上 "/en" 以进入英文页面
func switchToEnglish(thisURL *string) string {
	urlStruct, _ := url.Parse(*thisURL)
	path := urlStruct.Path
	path = "/en" + path
	urlStruct.Path = path
	return urlStruct.String()
}

func init() {
	w := crawlers.Register("swp_berlin", "Stiftung Wissenschaft und Politik",
		"https://www.swp-berlin.org/")

	w.SetStartingURLs([]string{
		"https://www.swp-berlin.org/sitemap.xml",
	})

	w.OnXML(`//loc`, func(element *colly.XMLElement, ctx *crawlers.Context) {
		switch {
		case strings.Contains(element.Text, "sitemap.xml"):
			w.Visit(element.Text, crawlers.Index)
		case strings.Contains(element.Text, "/wissenschaftler-in"):
			w.Visit(switchToEnglish(&element.Text), crawlers.Expert)
		case strings.Contains(element.Text, "/presse"):
			w.Visit(switchToEnglish(&element.Text), crawlers.News)
		case strings.Contains(element.Text, "/publikation"):
			w.Visit(switchToEnglish(&element.Text), crawlers.Report)
		}
	})

	// 获取 Title //
	w.OnHTML(`body > header.publication-page > div > h1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title //
	w.OnHTML(`body > div.news-header > div > h1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Language //
	w.OnHTML(`body > header.publication-page > div > ul.publication-list__languages > li > a > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Language = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle //
	w.OnHTML(`body > header.publication-page > div > p.subtitle`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle //
	w.OnHTML(`body > div.news-header > div > p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime //
	w.OnHTML(`.news-header > div > time, .publication-page > div > span.small-text`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		reg := regexp.MustCompile(`[\d./]+`)
		ctx.PublicationTime = reg.FindString(element.Text)
	})

	// 获取 Authors //
	w.OnHTML(`body > header.publication-page > div > ul.authors > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content //
	w.OnHTML(`.ce-bodytext, .publication-page__fulltext > div > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags //
	w.OnHTML(`body > header.publication-page > div > ul.publication-page__links > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File //
	w.OnHTML(`body > header.publication-page > section > div > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			fileURL := "https://www.swp-berlin.org" + element.Attr("href")
			ctx.File = append(ctx.File, fileURL)
		}
	})

	// 获取 File //
	w.OnHTML(`div > ul.downloads__list > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			fileURL := "https://www.swp-berlin.org" + element.Attr("href")
			ctx.File = append(ctx.File, fileURL)
		}
	})

	// 获取 Name //
	w.OnHTML(`body > section.webprofile > div > div.webprofile__text > header > h1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Title //
	w.OnHTML(`body > section.webprofile > div > div.webprofile__text > div.webprofile__profile > div > p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Email //
	w.OnHTML(`body > section.webprofile > div > div.webprofile__text > div.webprofile__profile > div > a.link--mail`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Email = strings.TrimSpace(element.Text)
	})

	// 获取 Phone //
	w.OnHTML(`body > section.webprofile > div > div.webprofile__text > div.webprofile__profile > div > .phone`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Phone = strings.TrimSpace(element.Text)
	})
}
