package production

import (
	"strconv"
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("china_org_cn", "中国新闻办公室", "http://www.china.org.cn/")

	w.SetStartingURLs([]string{
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
	w.OnHTML(`.columns1 > #autopage > center`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 仅在第一页访问所有其他 Index
		if element.ChildText(`span`) == "1" {
			urlList := element.ChildAttrs("a", "href")
			for _, url := range urlList {
				w.Visit(element.Attr(url), crawlers.Index)
			}
		}
	})

	// 访问 News 从 Index
	w.OnHTML(`.columns1 > ul > li > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 获取 Title
	w.OnHTML(`.wrapper #title`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Authors & PublicationTime
	w.OnHTML(`.wrapper #guild > dd`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		raw := strings.Replace(element.Text, ",", "*", 1)
		outList := strings.Split(raw, "*")
		var date string
		if len(outList) >= 2 {
			ctx.Authors = append(ctx.Authors, strings.TrimSpace(outList[0]))
			date = outList[1]
		} else {
			date = outList[0]
		}
		ctx.PublicationTime = strings.TrimSpace(date)
	})

	// 获取 CommentCount
	w.OnHTML(`.wrapper #guild > dd > span font#pinglun`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		str := strings.TrimSpace(element.Text)
		num, _ := strconv.Atoi(str)
		ctx.CommentCount = num
	})

	// 获取 Content
	w.OnHTML(`#container_txt`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p"))
	})

	// 获取 Expert 通过 SubContext 在 Index
	w.OnHTML(`div.apDiv1 > div>  table > tbody > tr`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		subCtx := ctx.CreateSubContext()
		subCtx.PageType = crawlers.Expert
		subCtx.Name = element.ChildText("td > b")
		subCtx.Description = element.ChildText(`td:nth-last-child(1)[valign="top"]`)
	})
}
