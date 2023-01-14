package usip

import (
	"github.com/gocolly/colly/v2"
	_ "github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("usip", "美国和平研究所战略稳定与安全办公室", "https://www.usip.org/")

	w.SetStartingUrls([]string{"https://www.usip.org/regions/asia/afghanistan",
		"https://www.usip.org/regions/americas/bolivia",
		"https://www.usip.org/regions/asia/burma",
		"https://www.usip.org/regions/africa/central-african-republic",
		"https://www.usip.org/regions/asia/china",
		"https://www.usip.org/regions/americas/colombia",
		"https://www.usip.org/regions/africa/democratic-republic-congo",
		"https://www.usip.org/regions/americas/el-salvador",
		"https://www.usip.org/regions/americas/guatemala",
		"https://www.usip.org/regions/americas/honduras",
		"https://www.usip.org/regions/asia/india",
		"https://www.usip.org/regions/middle-east-and-north-africa/iran",
		"https://www.usip.org/regions/middle-east-and-north-africa/iraq",
		"https://www.usip.org/regions/middle-east-and-north-africa/israel-and-palestinian-territories",
		"https://www.usip.org/regions/middle-east-and-north-africa/libya",
		"https://www.usip.org/regions/americas/nicaragua",
		"https://www.usip.org/regions/africa/nigeria",
		"https://www.usip.org/regions/asia/north-korea",
		"https://www.usip.org/regions/asia/pakistan",
		"https://www.usip.org/regions/asia/papua-new-guinea",
		"https://www.usip.org/regions/asia/philippines",
		"https://www.usip.org/regions/europe/russia",
		"https://www.usip.org/regions/africa/south-sudan",
		"https://www.usip.org/regions/africa/sudan",
		"https://www.usip.org/regions/middle-east-and-north-africa/syria",
		"https://www.usip.org/regions/middle-east-and-north-africa/tunisia",
		"https://www.usip.org/regions/europe/ukraine",
		"https://www.usip.org/regions/americas/venezuela",
		"https://www.usip.org/regions/asia/vietnam",
		"https://www.usip.org/regions/middle-east-and-north-africa/yemen",
		"https://www.usip.org/issue-areas/civilian-military-relations",
		"https://www.usip.org/issue-areas/conflict-analysis-prevention",
		"https://www.usip.org/issue-areas/democracy-governance",
		"https://www.usip.org/issue-areas/economics",
		"https://www.usip.org/issue-areas/electoral-violence",
		"https://www.usip.org/issue-areas/environment",
		"https://www.usip.org/issue-areas/fragility-resilience",
		"https://www.usip.org/issue-areas/gender",
		"https://www.usip.org/issue-areas/global-health",
		"https://www.usip.org/issue-areas/global-policy",
		"https://www.usip.org/issue-areas/human-rights",
		"https://www.usip.org/issue-areas/justice-security-rule-law",
		"https://www.usip.org/issue-areas/mediation-negotiation-dialogue",
		"https://www.usip.org/issue-areas/nonviolent-action",
		"https://www.usip.org/issue-areas/peace-processes",
		"https://www.usip.org/issue-areas/reconciliation",
		"https://www.usip.org/issue-areas/religion",
		"https://www.usip.org/issue-areas/violent-extremism",
		"https://www.usip.org/issue-areas/youth",
		"https://www.usip.org/publications",
		"https://www.usip.org/blog",
		"https://www.usip.org/whyallthecoups"})

	// 从翻页器获取链接并访问
	w.OnHTML("li.pager__item>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})
	w.OnHTML("section>a.btn.-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	// 从index访问新闻
	w.OnHTML(".summary__heading>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	// report.title
	w.OnHTML("main>header>div>div>h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// report.author
	w.OnHTML("body > div.dialog-off-canvas-main-canvas > main > header > div.container.row.-exact > div > p.meta", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("body > div.dialog-off-canvas-main-canvas > main > section.container.row.-exact > article > header > p:nth-child(3)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	//
	//report.publish time
	w.OnHTML("body > div.dialog-off-canvas-main-canvas > main > header > div.container.row.-exact > div > p.meta", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report .content
	w.OnHTML("section.intro-wysiwyg", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = ctx.Content + element.Text
	})
	w.OnHTML("section.wysiwyg", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})
	//内含Expert
	w.OnHTML("body > div.dialog-off-canvas-main-canvas > main > header > div.container.row.-exact > div > p.meta > font > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})
	// expert.Name
	w.OnHTML("header.page__header>h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})
	// expert.title
	w.OnHTML("header.page__header>p.meta", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	// expert.link
	w.OnHTML("div.-with-separator>div", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	// expert.description
	w.OnHTML("section.wysiwyg", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

}
