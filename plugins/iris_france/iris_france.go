package iris_france

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("iris_france", "Institut de Relations Internationales et Stratégiques", "https://www.iris-france.org/")

	w.SetStartingUrls([]string{
		"https://www.iris-france.org/programmes/industrie-de-defense-et-de-securite/",
		"https://www.iris-france.org/programmes/europe-et-strategie/",
		"https://www.iris-france.org/programmes/humanitaire-et-developpement/",
		"https://www.iris-france.org/programmes/sport-et-geopolitique/",
		"https://www.iris-france.org/programmes/sport-et-geopolitique/",
		"https://www.iris-france.org/programmes/amerique-latine-caraibe/",
		"https://www.iris-france.org/programmes/asie-pacifique/",
		"https://www.iris-france.org/programmes/afrique-s/",
		"https://www.iris-france.org/programmes/moyen-orient-afrique-du-nord/",
		"https://www.iris-france.org/securite-defense-et-nouveaux-risques/",
		"https://www.iris-france.org/equilibres-internationaux-et-mondialisation/",
		"https://www.iris-france.org/energie-et-environnement/",
		"https://www.iris-france.org/humanitaire-et-developpement/",
		"https://www.iris-france.org/enjeux-de-societe/",
		"https://www.iris-france.org/sport-et-relations-internationales/",
		"https://www.iris-france.org/aires-regionales/",
	})

	// 访问下一页 Index
	w.OnHTML(`ul[class="page-pagination-list clearfix"] > li:nth-last-child(1) > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.block-content > div .title > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`[class="block-content single-video-content"] > .single-title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title
	w.OnHTML(`.content > .single-title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`[class="block-content single-video-content"] > .article-date`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.content .date`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`[class="content row"] .intro-title`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`[class="block-content single-video-content"] > .article-author > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Authors
	w.OnHTML(`.block-content .content .author > span`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Content
	w.OnHTML(`[class="block-content single-video-content"]`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`.content`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})
}
