package mei

import (
	"github.com/araddon/dateparse"
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"time"
)

func init() {
	w := Crawler.Register("mei", "中东研究所", "https://www.mei.edu/")
	w.SetStartingUrls([]string{"https://www.mei.edu/policy-analysis", "https://www.mei.edu/experts"})

	//index
	w.OnHTML(".pager__item > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Index)
	})

	//news
	w.OnHTML("article.feature.feature-1 > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.News)
	})

	//标题
	w.OnHTML(".title-wrapper.container > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//时间
	w.OnHTML("date.post__date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	//作者
	w.OnHTML(".post__author", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	//正文
	w.OnHTML("div.content", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = element.Text
	})

	//专家
	w.OnHTML("figure.photo > a", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Expert)
	})

	//姓名
	w.OnHTML("h1.profile__name", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//头衔
	w.OnHTML("h3.profile__title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//描述
	w.OnHTML(" div.col-md-8 > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description += element.Text
	})

	w.OnHTML(".post__date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		date, err := dateparse.ParseAny(element.Text)
		if err != nil {
			return
		}
		ctx.PublicationTime = date.Format(time.RFC3339)
	})
}
