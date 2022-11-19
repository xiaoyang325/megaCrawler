package ifans

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// 这个函数用于从 onclick 函数调用中获取信息，拼接成 Report 的 URL，并返回
func getURLFromFunctionCall(functionCall string, channelType string) string {
	reg := regexp.MustCompile("fnCmdView\\('(\\d+)','(\\w+)'\\)")
	paramList := reg.FindStringSubmatch(functionCall)

	for index, value := range paramList {
		paramList[index] = strings.TrimSpace(value)
	}

	param1, param2 := paramList[1], paramList[2]
	fileUrl := "https://www.ifans.go.kr/knda/ifans/eng/act/"
	fileUrl += channelType + ".do?sn=" + param1 + "&boardSe=" + param2

	return fileUrl
}

// 这个函数修改当前 Index 页面的查询参数，以获取下一页 Index，并返回相应的 URL
func getNextIndexURL(current_url string, current_page_num string, param_name string) string {
	this_url, _ := url.Parse(current_url)
	param_list := this_url.Query()

	current_num, _ := strconv.Atoi(current_page_num)
	current_num++

	param_list.Set(param_name, strconv.Itoa(current_num))
	this_url.RawQuery = param_list.Encode()

	return this_url.String()
}

func init() {
	w := Crawler.Register("ifans", "外交与国家安全研究所", "https://www.ifans.go.kr/")

	w.SetStartingUrls([]string{
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityList.do?ctgrySe=02&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityList.do?ctgrySe=15&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityList.do?ctgrySe=03&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityList.do?ctgrySe=04&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityList.do?ctgrySe=17&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityList.do?ctgrySe=01&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityList.do?ctgrySe=18&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityList.do?ctgrySe=19&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityList.do?ctgrySe=20&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityAreaList.do?ctgrySe=06&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityAreaList.do?ctgrySe=11&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityAreaList.do?ctgrySe=12&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityAreaList.do?ctgrySe=07&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityAreaList.do?ctgrySe=08&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityAreaList.do?ctgrySe=13&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityAreaList.do?ctgrySe=09&pageIndex=1",
		"https://www.ifans.go.kr/knda/ifans/eng/act/ActivityAreaList.do?ctgrySe=10&pageIndex=1",
	})

	// 访问下一页 Index
	w.OnHTML(`#listForm > div.pagination > span.on`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			// 从当前 Index 的 URL 获取下一页 Index 的 URL
			next_url := getNextIndexURL(ctx.Url, strings.TrimSpace(element.Text), "pageIndex")
			w.Visit(next_url, Crawler.Index)
		})

	// 访问 Report 从 Index
	w.OnHTML(`#listForm > ul.board_list > li > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			var report_url string

			// 从 a[onclick] 中的函数调用获取 Report 的 URL
			if strings.Contains(ctx.Url, "ActivityList") {
				report_url = getURLFromFunctionCall(element.Attr("onclick"), "ActivityView")
			} else if strings.Contains(ctx.Url, "ActivityAreaList") {
				report_url = getURLFromFunctionCall(element.Attr("onclick"), "ActivityAreaView")
			}

			w.Visit(report_url, Crawler.Report)
		})

	// 获取 Title
	w.OnHTML(`#content > div > div.sub_top_view.con_in > strong.tit`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Title
	w.OnHTML(`#detailForm > div.editor.board_con_top.con_in > strong.tit`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Description
	w.OnHTML(`#detailForm > div.editor.board_con.con_in > span > span`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Description = strings.TrimSpace(element.Text)
		})

	// 获取 Description
	w.OnHTML(`#detailForm > div.board_con.con_in > span:nth-child(1)`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Description = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`#content > div > div.sub_top_view.con_in > span.date > em`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 PublicationTime
	w.OnHTML(`#detailForm > div.editor.board_con_top.con_in > span.date > em`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.PublicationTime = strings.TrimSpace(element.Text)
		})

	// 获取 CategoryText
	w.OnHTML(`#content > div > div.sub_top_view.con_in > span.subj`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.CategoryText = strings.TrimSpace(element.Text)
		})

	// 获取 Authors
	w.OnHTML(`#content > div > div.sub_top_view.con_in > strong.write`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
		})

	// 获取 ViewCount
	w.OnHTML(`#content > div > div.sub_top_view.con_in > span.look > em`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			str := strings.TrimSpace(element.Text)
			num, _ := strconv.Atoi(str)
			ctx.ViewCount = num
		})

	// 获取 ViewCount
	w.OnHTML(`#detailForm > div.editor.board_con_top.con_in > span.look > em`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			str := strings.TrimSpace(element.Text)
			num, _ := strconv.Atoi(str)
			ctx.ViewCount = num
		})

	// 获取 File
	w.OnHTML(`#detailForm > div.editor.board_con.con_in > dl > dd > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			file_url := "https://www.ifans.go.kr" + element.Attr("href")
			ctx.File = append(ctx.File, file_url)
		})

	// 获取 Tags
	w.OnHTML(`#detailForm > div.editor.board_con.con_in > span.tag > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			tag_str := strings.TrimSpace(element.Text)
			// 删除 Tag 中的 "#"
			tag_str = strings.Replace(tag_str, "#", "", 1)
			ctx.Tags = append(ctx.Tags, tag_str)
		})

	// 获取 Tags
	w.OnHTML(`#detailForm > div.editor.board_con_top.con_in > span.subj`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
		})
}
