package production

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strconv"
	"strings"
)

func init() {
	w := Crawler.Register("csdsafrica", "战略与国防研究中心", "https://csdsafrica.org/")

	w.SetStartingUrls([]string{
		"https://csdsafrica.org/close-protection-africa/",
		"https://csdsafrica.org/our-insights/",
	})

	// 访问 Report 从 Index
	w.OnHTML(`#main > div.fullsize > div > main > div > div > div > div.avia-content-slider.el_after_av_heading.avia-builder-el-last > div > div > article.slide-entry> a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 访问 Report 从 Index
	w.OnHTML(`div > div > div > div > div > div > div > article.slide-entry.flex_column.post-entry.post-format-standard > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		w.Visit(element.Attr("href"), Crawler.Report)
	})

	// 获取 Title
	w.OnHTML(`#main > div.main_color.container_wrap_first.container_wrap.sidebar_right > div > main > div > div > div.flex_column.av_one_full.avia-builder-el-1.el_after_av_image.el_before_av_one_half.first.flex_column_div.av-zero-column-padding > div > h1`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title
	w.OnHTML(`#main > div > div > main > article > div.entry-content-wrapper.clearfix.standard-content > header > h1`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 SubTitle
	w.OnHTML(`#main > div.main_color.container_wrap_first.container_wrap.sidebar_right > div > main > div > div > div.el_after_av_image.el_before_av_one_half.first.flex_column_div.av-zero-column-padding > div > div.av-subheading.av-subheading_below > p`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.SubTitle = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`#main > div.container_wrap.container_wrap_first.main_color.sidebar_right > div > main > article > div.entry-content-wrapper.clearfix.standard-content > div.entry-content > p[style="text-align: center;"] > strong`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		raw_str := strings.Replace(element.Text, "–", "", 1)
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(raw_str))
	})

	// 获取 CommentCount
	w.OnHTML(`#main > div.container_wrap.container_wrap_first.main_color.sidebar_right > div > main > div.comment-entry.post-entry > div.comment_meta_container > div > div > span.comment-count`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		str := strings.TrimSpace(element.Text)
		num, _ := strconv.Atoi(str)
		ctx.CommentCount = num
	})

	// 获取 Content
	w.OnHTML(`#main > div.main_color.container_wrap_first.container_wrap.sidebar_right > div > main > div > div`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("div.flex_column > section > div"))
	})

	// 获取 Content
	w.OnHTML(`#main > div.container_wrap.container_wrap_first.main_color.sidebar_right > div > main > article > div.entry-content-wrapper.clearfix.standard-content > div.entry-content`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 Tags
	w.OnHTML(`#main > div.container_wrap.container_wrap_first.main_color.sidebar_right > div > main > article > div.entry-content-wrapper.clearfix.standard-content > header > span > span > a`, func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})
}
