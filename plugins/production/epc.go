package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("epc", "欧洲国家政治经济中心", "https://epc.eu/")

	w.SetStartingUrls([]string{"https://epc.eu/en/search?tag=6",
		"https://epc.eu/en/search?tag=368",
		"https://epc.eu/en/search?tag=509",
		"https://epc.eu/en/search?tag=559",
		"https://epc.eu/en/search?tag=599",
		"https://epc.eu/en/search?tag=740",
		"https://epc.eu/en/search?tag=60",
		"https://epc.eu/en/search?tag=471",
		"https://epc.eu/en/search?tag=517",
		"https://epc.eu/en/search?tag=564",
		"https://epc.eu/en/search?tag=617",
		"https://epc.eu/en/search?tag=764",
		"https://epc.eu/en/search?tag=65",
		"https://epc.eu/en/search?tag=498",
		"https://epc.eu/en/search?tag=531",
		"https://epc.eu/en/search?tag=596",
		"https://epc.eu/en/search?tag=702",
		"https://epc.eu/en/search?tag=886"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = Crawler.Report
		}
	})

	// 从index访问新闻
	w.OnHTML("div.container.searchpage > div:nth-child(5) > div:nth-child(2) > div > div > div> div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//内含Expert
	w.OnHTML("div.expertname>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})
	// expert.Name
	w.OnHTML("div.expertname", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	// expert.title
	w.OnHTML("div.expertexpertise>span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	// expert.link
	w.OnHTML("div.experticons>a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})

	// expert.description
	w.OnHTML("div.expertdescription", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	// expert.area
	w.OnHTML(" div.row.whitebg > div.borderleft.col-xs-12.col-sm-4.col-md-4.col-lg-3.expertdetail.f-content-right > font:nth-child(6) > font", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})
	//expert language
	w.OnHTML("div.row.whitebg > div.borderleft.col-xs-12.col-sm-4.col-md-4.col-lg-3.expertdetail.f-content-right > font:nth-child(19) > font", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Language = element.Text
	})

	// report.title
	w.OnHTML("div.projecttitle>h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})
	//report.publish time
	w.OnHTML("div.publidetailpage", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.author
	w.OnHTML("div.pubiauthorname", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// report .content
	w.OnHTML("div.eventtext", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

}
