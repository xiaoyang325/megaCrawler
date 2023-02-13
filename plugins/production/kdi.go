package production

import (
	"megaCrawler/crawlers"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("kdi", "发展研究会", "https://www.kdi.re.kr/")

	w.SetStartingURLs([]string{
		"https://www.kdi.re.kr/kdi_eng/topics/dep_strategy.jsp",
		"https://www.kdi.re.kr/kdi_eng/topics/dep_policy.jsp",
		"https://www.kdi.re.kr/kdi_eng/topics/office_studies.jsp",
		"https://www.kdi.re.kr/kdi_eng/issues/policy_information.jsp",
		"https://www.kdi.re.kr/kdi_eng/topics/office_global.jsp",
	})

	// 访问下一页 Index
	w.OnHTML(`.list_pagination > span > a.on`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		pageNum := strings.TrimSpace(element.Text)
		num, _ := strconv.Atoi(pageNum)
		nextURL := "https://www.kdi.re.kr/kdi_eng/issues/policy_information.jsp?pg="
		nextURL += strconv.Itoa(num+1) + "&pp="
		w.Visit(nextURL, crawlers.Index)
	})

	// 访问 Expert's Index
	w.OnHTML(`#ui_contents > div.page_contents > div > ul > li:nth-child(2) > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 Report 从 Index 到 subCtx
	w.OnHTML(`.board_list_wrap > ul > li`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = crawlers.Report
		subCtx.Title = strings.TrimSpace(element.ChildText(".list_tit > a > strong"))
		subCtx.PublicationTime = strings.TrimSpace(element.ChildText(".list_tit > a > .data > span:nth-child(2)"))
		subCtx.Content = strings.TrimSpace(element.ChildText(".more_txt"))
	})

	// 访问 Expert 从 Index 到 subCtx
	w.OnHTML(`#ui_contents > div.page_contents > div > div[class="topic_issues expert"] > div > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = crawlers.Expert
		subCtx.Name = strings.TrimSpace(element.ChildText(".info > strong > a"))
		subCtx.Title = strings.TrimSpace(element.ChildText(".info > p"))
		subCtx.Description = strings.TrimSpace(element.ChildText(".more_contents"))
	})

	// 访问 "More information" 到 Report 从 Index
	w.OnHTML(`.btn > .btn_more`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Report)
	})

	// 获取 Title
	w.OnHTML(`dl > dd > div > strong`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title (Type 2)
	w.OnHTML(`ul > li.title > strong`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Title (Type 3)
	w.OnHTML(`dd > strong`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Content
	w.OnHTML(`.cnts_detail`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime
	w.OnHTML(`#ui_contents > div > div.board_view_wrap > div.top_title > dl > dd > div > p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime (Type 2)
	w.OnHTML(`#ui_contents > div > div.board_view_wrap > div.top_title > div > ul > li:nth-child(2) > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime (Type 3)
	w.OnHTML(`#ui_contents > div > div.board_view_wrap.seminar_view > div.top_title > dl > dd > p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = strings.TrimSpace(element.Text)
	})

	// 获取 Location
	w.OnHTML(`#ui_contents > div > div.board_view_wrap.seminar_view > div.repoart_contents > div:nth-child(2) > div.cnts_left > ul > li:nth-child(1)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Location = strings.TrimSpace(element.Text)
	})

	// 获取 Authors
	w.OnHTML(`#ui_contents > div > div.board_view_wrap > div.top_title > dl > dd > ul > li.author > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 获取 Language
	w.OnHTML(`#ui_contents > div > div.board_view_wrap > div.top_title > dl > dd > ul > li.float_left > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Language = strings.TrimSpace(element.Text)
	})

	// 获取 Language (Type 2)
	w.OnHTML(`#ui_contents > div > div.board_view_wrap > div.top_title > div > ul > li:nth-child(3) > span`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Language = strings.TrimSpace(element.Text)
	})

	// 获取 File
	w.OnHTML(`#ui_contents > div > div.board_view_wrap > div.top_title > div > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		rawStr := element.Attr("onclick")
		rawStr = strings.Replace(rawStr, "downloadvar(", "", 1)
		rawStr = strings.Replace(rawStr, "download2(", "", 1)
		rawStr = strings.Replace(rawStr, ");return false;", "", 1)
		rawStr = strings.ReplaceAll(rawStr, "'", "")
		paramList := strings.Split(rawStr, ",")
		for index, value := range paramList {
			paramList[index] = strings.TrimSpace(value)
		}
		param1, param2, param3 := paramList[0], paramList[1], paramList[2]

		fileURL := "https://www.kdi.re.kr/kdi_eng/common/report_download.jsp?list_no="
		fileURL += param1 + "&member_pub=" + param2 + "&type=" + param3
		fileURL += "&cacheclear=81"

		ctx.File = append(ctx.File, fileURL)
	})
}
