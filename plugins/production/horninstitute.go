package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("horninstitute", "Horn Institute", "https://horninstitute.org")
	w.SetStartingURLs([]string{"https://horninstitute.org/blogs/", "https://horninstitute.org/reports/", "https://horninstitute.org/papers/"})

	w.OnHTML(".title_link > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	w.OnHTML("a.page-numbers", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	w.OnHTML(".cz_post_content .vc_col-sm-4 .vc_column-inner", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.File = append(subCtx.File, element.ChildAttr(".cz_btn", "href"))
		subCtx.Title = element.ChildText("strong")
		subCtx.Image = append(subCtx.Image, element.ChildAttrs("img", "src")...)
		subCtx.Content = element.ChildText(".cz_post_content")
		subCtx.PageType = crawlers.News
	})

	w.OnHTML(".section_title ", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	w.OnHTML(".byline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		authorAndDate := strings.TrimSpace(element.Text)
		date := strings.TrimSpace(element.ChildText(".postdate"))
		author := strings.TrimSuffix(authorAndDate, date)
		ctx.Authors = append(ctx.Authors, author)
		ctx.PublicationTime = date
	})

	w.OnHTML(".cz_post_content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}
