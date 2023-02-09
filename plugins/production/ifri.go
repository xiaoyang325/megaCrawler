package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("ifri", "Institut français des relations internationales",
		"https://www.ifri.org/")

	w.SetStartingUrls([]string{
		"https://www.ifri.org/fr/espace-media/communiques",
		"https://www.ifri.org/fr/espace-media/dossiers-dactualite",
		"https://www.ifri.org/fr/espace-media/lifri-medias",
		"https://www.ifri.org/fr/publications",
		"https://www.ifri.org/fr/equipe",
	})

	// 访问下一页 Index
	w.OnHTML(`.pagination > li[class="next last"] > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`article > .vignette-result > div > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`.main-title > span`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		raw := element.Text
		raw = strings.Replace(raw, element.ChildText(`.sub-title`), "", 1)
		raw = strings.Replace(raw, element.ChildText(`.border-bottom`), "", 1)
		ctx.Title = strings.TrimSpace(raw)
	})

	// 获取 Name
	w.OnHTML(`.main-content > .bandeau-titre-page > .bloc-titre > .titre-page`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PageType = Crawler.Expert
		ctx.Name = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Title
	w.OnHTML(`div > div.col-xs-12.col-md-11 > div > div.row.head-expert > div.col-xs-12.col-sm-12.col-md-6.col-lg-6.right-part-expert > div > div > div > div > div > p:nth-child(1) > strong`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`.main-title > span > .sub-title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 Description
	w.OnHTML(`.chapo [class="field-item even"]  > p`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Description
	w.OnHTML(`.tab-content [class="field-item even"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.main-content .contenu-date-publish`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = Crawler.StandardizeSpaces(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`.thematique-contenu .main-thematique`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		str := strings.Replace(element.Text, "mailto:", "", 1)
		ctx.CategoryText = strings.TrimSpace(str)
	})

	// 获取 Email
	w.OnHTML(`.links-expert > a[title="Envoyer un mail"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Email = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`.content-text-riche [class="field-item even"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`.list-motcle > span`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`.document-dl-link > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.File = append(ctx.File, element.Attr("href"))
	})
}
