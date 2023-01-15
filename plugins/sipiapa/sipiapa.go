package sipiapa

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("sipiapa", "美洲新闻协会", "https://www.sipiapa.org/")

	w.SetStartingUrls([]string{"https://www.sipiapa.org/contenidos/noticias-impunidad.html",
		"https://www.sipiapa.org/contenidos/leyes-libertad-de-prensa.html",
		"https://www.sipiapa.org/contenidos/jurisprudencia-chapultepec.html",
		"https://www.sipiapa.org/contenidos/informes.html",
		"https://www.sipiapa.org/contenidos/documentos-relevantes.html",
		"https://www.sipiapa.org/contenidos/recursos-y-publicaciones.html",
		"https://www.sipiapa.org/contenidos/covid-19.html"})

	// 从index访问新闻
	w.OnHTML(".title>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML(".title2>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("h1.title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report.description
	w.OnHTML("blockquote.copete", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})
	// report.author
	w.OnHTML("div.article-content>p:nth-child(1)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// report .content
	w.OnHTML("div.article-content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
}
