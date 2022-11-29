package arinus

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"regexp"
)

func init() {
	w := Crawler.Register("arinus", "亚洲研究所", "https://ari.nus.edu.sg/")
	w.SetStartingUrls([]string{"https://ari.nus.edu.sg/about-ari/people/academic/",
		"https://ari.nus.edu.sg/about-ari/people/administrative/",
		"https://ari.nus.edu.sg/media/news/"})

	//index
	w.OnHTML("a.page-link", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问人物
	w.OnHTML("people-info", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//访问新闻
	w.OnHTML("div.publication-info > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//人物姓名,新闻标题
	w.OnHTML("people-title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Name = element.Text
		} else if ctx.PageType == Crawler.News {
			ctx.Title = element.Text
		}
	})

	//头衔
	w.OnHTML("table > tbody > tr:nth-child(1) > td:nth-child(3)", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//介绍
	w.OnHTML("div.fl-col.col-md-12.col-sm-12.col-xs-12", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			ctx.Description = element.Text
		}
	})

	//邮箱与电话,新闻正文
	w.OnHTML("table.table-borderless.table-people-info", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.Expert {
			emailRegex, _ := regexp.Compile("Email: ([.@\\w]+)")
			telRegex, _ := regexp.Compile("Tel: ([.\\w]+)")
			emailMatch := emailRegex.FindStringSubmatch(element.Text)
			telMatch := telRegex.FindStringSubmatch(element.Text)
			if len(emailMatch) == 2 {
				ctx.Email = emailMatch[1]
			}
			if len(telMatch) == 2 {
				ctx.Phone = telMatch[1]
			}
		} else if ctx.PageType == Crawler.News {
			ctx.Content = element.Text
		}
	})

}
