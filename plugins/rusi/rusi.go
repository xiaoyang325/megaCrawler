package rusi

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
)

func init() {
	w := Crawler.Register("rusi", "皇家联合服务研究所", "https://rusi.org/")
	w.SetStartingUrls([]string{"https://rusi.org/explore-our-research/topics",
		"https://rusi.org/explore-our-research/region-and-country-groups"})

	//访问索引页
	w.OnHTML("a.TagLink-module--component--Yq0DF", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//访问人物
	w.OnHTML("#our-experts > div > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//获取人物姓名
	w.OnHTML(" div.ProfileTitleBlock-module--text--qfxrH > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//获取人物头衔
	w.OnHTML(" div.ProfileTitleBlock-module--text--qfxrH > samll", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取人物领域
	w.OnHTML("aside > ul > li > a > span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})

	//获取人物描述
	w.OnHTML("div.Section-module--content--Of\\+As > div > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	//访问报告
	w.OnHTML("#latest-publications > div.SearchComponent-module--component--LLifa > div > ul > li > div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	//获取报告,新闻标题
	w.OnHTML(" div:nth-child(1) > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		if ctx.PageType == Crawler.News {
			ctx.Title = element.Text
		} else if ctx.PageType == Crawler.Report {
			ctx.Title = element.Text
		}
	})

	//获取作者
	w.OnHTML(" div.Layout-module--component--JbLd9 > div > main > article > div.Article-module--contentArea--7IPcP > div.Article-module--sidebar--7hr7W > aside:nth-child(2) > section > a > div > div.PersonCard-module--text--UQlnj > div > h3",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, element.Text)
		})

	//获取标签
	w.OnHTML("#gatsby-focus-wrapper > div:nth-child(2) > div.Layout-module--component--JbLd9 > div > main > article > div.Article-module--contentArea--7IPcP > div:nth-child(1) > section.KeywordTags-module--component--s-ZCe.Article-module--section--vy2LJ.hideOnPrint > div:nth-child(2) > ul > li > a > span",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Tags = append(ctx.Tags, element.Text)
		})

	//pdf
	w.OnHTML("#gatsby-focus-wrapper > div:nth-child(2) > div.Layout-module--component--JbLd9 > div > main > article > a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.File = append(ctx.File, element.Attr("href"))
		})

	//访问新闻
	w.OnHTML("#news-\\&-comment > ul > li> div > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})
	w.OnHTML("#latest-news-and-comment > ul > li > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//获取正文
	w.OnHTML("#gatsby-focus-wrapper > div:nth-child(2) > div.Layout-module--component--JbLd9 > div > main > article > div.Article-module--contentArea--7IPcP > div:nth-child(1) > div",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = element.Text
		})
}
