package asaninst

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

// 这个函数用于分隔使用 "," 和 "and" 的字符串
// 并返回分割开的 []string
func cutToList(input_str string) ([]string)  {
	name_str := strings.Replace(input_str, "and", ",", 1)
	name_list := strings.Split(name_str, ",")
	for index, value := range name_list {
		name_list[index] = strings.TrimSpace(value)
	}

	return name_list
}

func init() {
	w := Crawler.Register("asaninst", "峨山政策研究院", "https://en.asaninst.org/")
	
	w.SetStartingUrls([]string{
		"https://en.asaninst.org/contents/issues/security/",
		"https://en.asaninst.org/contents/issues/international-law/",
		"https://en.asaninst.org/contents/issues/culture-and-society/",
		"https://en.asaninst.org/contents/issues/economy/",
		"https://en.asaninst.org/contents/issues/foreign-relations/",
		"https://en.asaninst.org/contents/issues/energy/",
		"https://en.asaninst.org/contents/issues/global-governance/",
		"https://en.asaninst.org/contents/issues/democracy-2/",
		"https://en.asaninst.org/contents/issues/science-and-technology/",
		"https://en.asaninst.org/contents/issues/nuclear-issues/",
		"https://en.asaninst.org/regions/",
	})

	// 访问下一页 Index
	w.OnHTML(`#content > div > div.paging > ul > li > a[class="next page-numbers"]`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Index)
		})

	// 访问 Report 从 Index
	w.OnHTML(`#content > div > div:nth-child(3) > article > div.post_desc.right > h3 > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 访问 Expert 从 Report
	w.OnHTML(`#content div.post_export_wrap.bg_gray > div > div.author_desc.right > h5 > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Expert)
		})

	// 获取 Title
	w.OnHTML(`#content header > div.single_post_info > h3`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`#content header > div.post_date_big`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 Authors
	w.OnHTML(`#content header > div.single_post_info > ul > li:nth-child(2) > div > span > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 CategoryText
	w.OnHTML(`#content > header > h2.archive-title`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 Content
	w.OnHTML(`#content > article > div.entry-content`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.ChildText("p"))
		})

	// 获取 Tags
	w.OnHTML(`#content header > div.single_post_info > ul > li > div`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			// 将字符串切成 []string 并填入
			ctx.Tags = append(ctx.Tags, cutToList(element.Text)...)
		})

	// 获取 Tags
	w.OnHTML(`#tertiary > div > div > aside:nth-child(4) > div > div.tag_wrap > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Tags = append(ctx.Tags, element.Text)
		})

	// 获取 Tags
	w.OnHTML(`#content div.tag_wrap > div > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Tags = append(ctx.Tags, element.Text)
		})

	// 获取 File
	w.OnHTML(`#content div.entry-meta > div.post_download > a.pdf`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.File = append(ctx.File, element.Attr("href"))
		})

	// 获取 Expert 的相关信息
	w.OnHTML(`#content > div > div.list_experts > article > div.member_desc.right`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Name = strings.TrimSpace(element.ChildText("h3"))
			// 去除其他信息以得到 Expert 的 Title
			title_raw := strings.Replace(element.Text, ctx.Name, "", 1)
			title_raw = strings.Replace(element.Text, element.ChildText("p"), "", 1)
			ctx.Title = strings.TrimSpace(title_raw)
		})
}
