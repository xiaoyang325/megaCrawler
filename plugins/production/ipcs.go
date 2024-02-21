package production

import (
	"errors"
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("ipcs", "和平与冲突研究所", "http://www.ipcs.org/")

	w.SetStartingURLs([]string{"http://www.ipcs.org/commentaries.php", "http://www.ipcs.org/issue_briefs.php",
		"http://www.ipcs.org/special_reports.php", "http://www.ipcs.org/discussion_reports.php", "http://www.ipcs.org/expert_media.php", "http://www.ipcs.org/research_paper.php"})

	w.OnResponse(func(response *colly.Response, ctx *crawlers.Context) {
		if strings.Contains(string(response.Body), "connection unsuccessful") {
			crawlers.RetryRequest(response.Request, errors.New("connection unsuccessful"), w)
		}
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// COMMENTARIES  -> new
	// visit commentaries

	w.OnHTML(".clearfix>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 从翻页器获取链接并访问
	w.OnHTML("#pagination>ul>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	// 访问new
	w.OnHTML(".clearfix>a:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// new .publish time [time中含有评论，是否删除？]
	w.OnHTML("#main_wrapper > section > div > div:nth-child(2) > div.col-md-9 > p:nth-child(3)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.Split(element.Text, "·")[0]
	})
	// new description
	w.OnHTML("#main_wrapper > section > div > div:nth-child(2) > div.col-md-9 > p:nth-child(6)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})
	//#main_wrapper > section > div > div:nth-child(2) > div.col-md-9 > p:nth-child(4)
	// new . author_name
	w.OnHTML("#main_wrapper > section > div > div:nth-child(3) > div.col-md-9 > div > div.col-md-3 > div > a > h6", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// new ,author . information
	w.OnHTML("#main_wrapper > section > div > div:nth-child(3) > div.col-md-9 > div > div.col-md-3 > div > span:nth-child(3)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// new.content
	w.OnHTML("#main_wrapper > section > div > div:nth-child(3) > div.col-md-9 > div > div.col-md-9", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	//https://www.csis.org/analysis -> report  http://www.ipcs.org/special_reports.php->report http://www.ipcs.org/discussion_reports.php->report http://www.ipcs.org/discussion_reports.php->report http://www.ipcs.org/research_paper.php->report
	// 从翻页器获取链接并访问
	w.OnHTML("#pagination > ul > li:nth-child > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问report
	w.OnHTML("#main_wrapper > section > div > div > div.col-md-9 > div:nth-child(2) > div > ul > li:nth-child > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})
	// 访问report -http://www.ipcs.org/research_paper.php
	w.OnHTML("#main_wrapper > section > div > div > div.col-md-9 > div:nth-child(2) > div > ul > li:nth-child > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// reort .author .publish time . catagory
	w.OnHTML("#main_wrapper > section > div > div.col-md-9 > div > span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		texts := strings.Split(element.Text, "·")
		authorName := texts[0]
		authorName = strings.Replace(authorName, "&nbsp;", "", 100)
		authorName = strings.Replace(authorName, "&amp;", "", 100)
		authorName = strings.TrimSpace(authorName)
		ctx.Authors = append(ctx.Authors, authorName)

		publishTime := texts[1]
		publishTime = strings.Replace(publishTime, "&nbsp;", "", 100)
		publishTime = strings.TrimSpace(publishTime)
		ctx.PublicationTime = publishTime

		categoryText := texts[2]
		categoryText = strings.Replace(categoryText, "&nbsp;", "", 100)
		categoryText = strings.TrimSpace(categoryText)
		ctx.CategoryText = categoryText
	})
	// reort.content
	w.OnHTML("#main_wrapper > section > div > div.col-md-9", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("#main_wrapper > section > div > div.col-md-9 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})
	w.OnHTML("#main_wrapper > section > div > div.col-md-3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	//http://www.ipcs.org/ipcs_books_reviews.php ->new
	// 访问new
	w.OnHTML("#main_wrapper > section > div > div.col-md-9 > div.rows_container > div > ul > li:nth-child > div.col-md-10 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	w.OnHTML(".main_title > h5", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	w.OnHTML("div.rows_container > div.col-md-9 > p:nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// new.anthor
	w.OnHTML("#main_wrapper > section > div:nth-child(2) > div > span > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// new .content
	w.OnHTML("#main_wrapper > section > div:nth-child(2) > div > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	//http://www.ipcs.org/expert_media.php
	// 访问new
	w.OnHTML("##main_wrapper > section > div > div.col-md-9 > div.rows_container > div > ul > li:nth-child > a:nth-child(1)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
		ctx.CategoryText = "IPCS EXPERTS IN THE MEDIA"
	})
}
