package prri

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

func init() {
	w := Crawler.Register("prri", "公共宗教研究所",
		"https://www.prri.org/")

	w.SetStartingUrls([]string{
		"https://www.prri.org/topic/abortion-reproductive-health/",
		"https://www.prri.org/topic/climate-change-science/",
		"https://www.prri.org/topic/economy/",
		"https://www.prri.org/topic/immigration/",
		"https://www.prri.org/topic/law-criminal-justice/",
		"https://www.prri.org/topic/lgbt/",
		"https://www.prri.org/topic/politics-elections/",
		"https://www.prri.org/topic/race-ethnicity/",
		"https://www.prri.org/topic/religion-culture/",
		"https://www.prri.org/topic/sports/",
	})

	// 从子频道入口中访问Index
	w.OnHTML("div.researchLandingSearch__button>a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 从Index中访问文章。
	w.OnHTML("a.searchResult__title",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 从翻页器获取下一页Index并访问。
	w.OnHTML(".pagination>a[class=\"next page-numbers\"]",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 从文章中获取Title并添加到ctx。
	w.OnHTML("div.researchArticleHeading__title", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 从文章中获取Author并添加到ctx。
	w.OnHTML("div.researchArticleHeading__author>a",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			name := strings.Split(element.Text, ",")[0]
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(name))
		})

	// 从文章中获取Author并添加到ctx。
	w.OnHTML(".researchArticleHeading__author",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			name := strings.Split(element.Text, ",")[0]
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(name))
		})

	w.OnHTML(".researchArticleHeading__date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML(".press__date", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.PublicationTime = element.Text
	})

	w.OnHTML("meta[property=\"og:description\"]", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Attr("content")
	})

	// 从文章中获取Tag并添加到ctx。（Report）（News）
	w.OnHTML(".researchArticleHeading__singleTag",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
		})

	// 从文章中获取Content并添加到ctx。
	w.OnHTML(".has-content-area",
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})
}
