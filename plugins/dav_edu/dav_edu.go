package dav_edu

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"net/url"
	"strconv"
	"strings"
)

// 这个函数修改当前 Index 页面的查询参数，以获取下一页 Index，并返回相应的 URL
func getNextIndexURL(current_url string, current_page_num string,
	param_name string) string {
	this_url, _ := url.Parse(current_url)
	param_list := this_url.Query()

	current_num, _ := strconv.Atoi(current_page_num)
	current_num++

	param_list.Set(param_name, strconv.Itoa(current_num))
	this_url.RawQuery = param_list.Encode()

	return this_url.String()
}

func init() {
	w := Crawler.Register("dav_edu", "外交学院", "https://www.dav.edu.vn/")

	w.SetStartingUrls([]string{
		"https://www.dav.edu.vn/su-kien-hoi-thao-toa-dam/?trang=1",
		"https://www.dav.edu.vn/gioi-thieu-chung-nghien-cu/",
		"https://www.dav.edu.vn/an-pham-nghien-cuu/?trang=1",
	})

	// 访问下一页 Index
	w.OnHTML(`[class="page-item active"] > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			url := getNextIndexURL(ctx.Url, element.Text, "trang")
			w.Visit(url, Crawler.Index)
		})

	// 访问 Report 从 Index
	w.OnHTML(`.row .story__title > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 获取 Title
	w.OnHTML(`.detail__title`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Description
	w.OnHTML(`.detail__summary`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Description = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`.detail__meta`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 CategoryText
	w.OnHTML(`li[class="breadcrumb-item active"] > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 Content
	w.OnHTML(`html .detail__content`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})
}
