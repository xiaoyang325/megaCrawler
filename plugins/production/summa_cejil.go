package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("summa_cejil", "国际法与司法中心 SUMMA频道",
		"https://summa.cejil.org/")

	w.SetStartingUrls([]string{
		`https://summa.cejil.org/es/library/?q=(order:desc,sort:metadata.fecha,types:!(%2758b2f3a35d59f31e1345b4ac%27,%2758b2f3a35d59f31e1345b471%27,%2758b2f3a35d59f31e1345b482%27,%2758b2f3a35d59f31e1345b479%27,%275a4d294c79f3f44b101e2816%27,%2759ee427f40f4a54920bd9b67%27,%2759ee461f40f4a54920bd9b87%27))`,
	})

	// 访问下一页 Index
	w.OnHTML(`#app > div.content > div > div > div > main > div > div > div.row > div > a:nth-child(1)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 Report 从 Index
	w.OnHTML(`.item-group.item-group-zoom-0 > div > div.item-actions > div > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title
	w.OnHTML(`#tabpanel-metadata > div > div.view > div > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`.metadata-type-date.metadata-name-fecha > dd`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Location
	w.OnHTML(`.metadata-type-relationship.metadata-name-pa_s > dd > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Location = strings.TrimSpace(element.Text)
	})

	// 获取 Language
	w.OnHTML(`.filelist  div.file .badge > span.translation`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Language = strings.TrimSpace(element.Text)
	})

	// 获取 CategoryText
	w.OnHTML(`.metadata-type-multiselect.metadata-name-tipo > dd > ul > li:nth-child(1)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`#tabpanel-metadata > div > div.view > div > span > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 File
	w.OnHTML(`a.file-download.btn.btn-outline-secondary`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		file_url := "https://summa.cejil.org" + element.Attr("href")
		ctx.File = append(ctx.File, file_url)
	})
}
