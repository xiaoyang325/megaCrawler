package china_org_cn

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strconv"
	"strings"
)

func init() {
	w := Crawler.Register("china_org_cn", "中国新闻办公室", "http://www.china.org.cn/")

	w.SetStartingUrls([]string{
		"http://www.china.org.cn/china/node_7075073.htm",
		"http://www.china.org.cn/china/node_7075074.htm",
		"http://www.china.org.cn/china/node_7075075.htm",
		"http://www.china.org.cn/china/node_7075076.htm",
		"http://www.china.org.cn/china/node_7075077.htm",
		"http://www.china.org.cn/world/node_7075229.htm",
		"http://www.china.org.cn/world/node_7075230.htm",
		"http://www.china.org.cn/world/node_7075231.htm",
		"http://www.china.org.cn/world/node_7075232.htm",
		"http://www.china.org.cn/world/node_7075233.htm",
		"http://www.china.org.cn/business/node_7074857.htm",
		"http://www.china.org.cn/business/node_7074861.htm",
		"http://www.china.org.cn/business/node_7074862.htm",
		"http://www.china.org.cn/business/node_7164397.htm",
		"http://www.china.org.cn/business/node_7164398.htm",
		"http://www.china.org.cn/business/node_7074864.htm",
		"http://www.china.org.cn/business/node_7074865.htm",
		"http://www.china.org.cn/business/node_7074866.htm",
		"http://www.china.org.cn/business/node_7074867.htm",
		"http://www.china.org.cn/business/node_7074868.htm",
		"http://www.china.org.cn/business/node_7074869.htm",
		"http://www.china.org.cn/business/node_7074870.htm",
		"http://www.china.org.cn/business/node_7074871.htm",
		"http://www.china.org.cn/business/node_7074858.htm",
		"http://www.china.org.cn/opinion/node_7074949.htm",
		"http://www.china.org.cn/opinion/node_7074948.htm",
		"http://www.china.org.cn/opinion/node_7164234.htm",
	})

	// 访问下一页 Index
	w.OnHTML(`.columns1 > #autopage > center`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			// 仅在第一页访问所有其他 Index
			if element.ChildText(`span`) == "1" {
				url_list := element.ChildAttrs("a", "href")
				for _, url := range url_list {
					w.Visit(element.Attr(url), Crawler.Index)
				}
			}
		})

	// 访问 News 从 Index
	w.OnHTML(`.columns1 > ul > li > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.News)
		})

	// 获取 Title
	w.OnHTML(`.wrapper #title`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Authors & PublicationTime
	w.OnHTML(`.wrapper #guild > dd`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			raw := strings.Replace(element.Text, ",", "*", 1)
			out_list := strings.Split(raw, "*")
			var date string
			if len(out_list) >= 2 {
				ctx.Authors = append(ctx.Authors, strings.TrimSpace(out_list[0]))
				date = out_list[1]
			} else {
				date = out_list[0]
			}
			ctx.PublicationTime = strings.TrimSpace(date)
		})

	// 获取 CommentCount
	w.OnHTML(`.wrapper #guild > dd > span font#pinglun`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			str := strings.TrimSpace(element.Text)
			num, _ := strconv.Atoi(str)
			ctx.CommentCount = num
		})

	// 获取 Content
	w.OnHTML(`#container_txt`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.ChildText("p"))
		})

	// 获取 Expert 通过 SubContext 在 Index
	w.OnHTML(`div.apDiv1 > div>  table > tbody > tr`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			sub_ctx := ctx.CreateSubContext()
			sub_ctx.PageType = Crawler.Expert
			sub_ctx.Name = element.ChildText("td > b")
			sub_ctx.Description = element.ChildText(`td:nth-last-child(1)[valign="top"]`)
		})
}
