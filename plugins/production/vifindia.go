package production

import (
	"megaCrawler/crawlers"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// 用于把名字从职务中分离的函数。
func cutOutName(name string) string {
	// 将名字从职位中分割出来。
	name = strings.Split(name, ",")[0]
	// 去掉名字前后的空格。
	name = strings.TrimSpace(name)

	return name
}

func init() {
	w := crawlers.Register("vifindia",
		"维韦卡南达国际基金会", "https://www.vifindia.org/")

	w.SetStartingURLs([]string{
		"https://www.vifindia.org/articles/terrorism",
		"https://www.vifindia.org/articles/JammuAndKashmir",
		"https://www.vifindia.org/articles/LeftwingExtremism",
		"https://www.vifindia.org/articles/NorthEastInsurgency",
		"https://www.vifindia.org/areaofstudy/articles/illegal-immigration",
		"https://www.vifindia.org/articles/nuclear-Disarmament",
		"https://www.vifindia.org/articles/bcsecurity",
		"https://www.vifindia.org/articles/defence",
		"https://www.vifindia.org/articles/policelawandorder",
		"https://www.vifindia.org/articles/climatechange",
		"https://www.vifindia.org/articles/disastermanagement",
		"https://www.vifindia.org/articles/EnergySecurity",
		"https://www.vifindia.org/articles/cybersecurity",
		"https://www.vifindia.org/articles/National-Security-and-Strategic-Studies/others",
		"https://www.vifindia.org/articles/africa",
		"https://www.vifindia.org/articles/centralasia",
		"https://www.vifindia.org/articles/china/",
		"https://www.vifindia.org/articles/europe",
		"https://www.vifindia.org/articles/indopacific",
		"https://www.vifindia.org/articles/us",
		"https://www.vifindia.org/articles/russia",
		"https://www.vifindia.org/articles/westasia",
		"https://www.vifindia.org/articles/Indianoceanregion",
		"https://www.vifindia.org/articles/international-relationships/others",
		"https://www.vifindia.org/articles/afghanistan",
		"https://www.vifindia.org/articles/bangladesh",
		"https://www.vifindia.org/articles/bhutan",
		"https://www.vifindia.org/articles/maldives",
		"https://www.vifindia.org/articles/myanmar",
		"https://www.vifindia.org/articles/nepal",
		"https://www.vifindia.org/articles/pakistan",
		"https://www.vifindia.org/articles/srilanka",
		"https://www.vifindia.org/articles/tibet",
		"https://www.vifindia.org/articles/Neighbourhood-Studies/others",
		"https://www.vifindia.org/articles/governanceandpolitica",
		"https://www.vifindia.org/articles/economicstudies",
		"https://www.vifindia.org/articles/technological_scientific",
		"https://www.vifindia.org/articles/historicalncivilisationalstudies",
		"https://www.vifindia.org/article",
		"https://www.vifindia.org/reports",
		"https://www.vifindia.org/taxonomy/term/8117",
		"https://www.vifindia.org/bookreview",
		"https://www.vifindia.org/paper-transcriptions",
		"https://www.vifindia.org/viewpoint",
		"https://www.vifindia.org/informationalert",
		"https://www.vifindia.org/informationnewsdigest",
		// "https://www.vifindia.org/policies_and_Perspectives",
		"https://www.vifindia.org/interaction",
		"https://www.vifindia.org/taxonomy/term/56",
		"https://www.vifindia.org/conferences-seminars",
		"https://www.vifindia.org/question-for-expert",
	})

	// 从翻页器获取链接并访问。
	w.OnHTML(".item-list > .pager > .pager-next > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从所有的Index中访问文章（除了/question-for-expert）。
	w.OnHTML(".article_title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 将从带有/article路径访问的标记为News
		if strings.Contains(ctx.URL, "/article") {
			w.Visit(element.Attr("href"), crawlers.News)
		} else { // 其他标记为Report
			w.Visit(element.Attr("href"), crawlers.Report)
		}
	})

	// 从/question-for-expert中访问文章。
	w.OnHTML("a[class=\"article_title expert_title\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 拼接为URL并访问。
		questionURL := "https://www.vifindia.org/" + element.Attr("href")
		w.Visit(questionURL, crawlers.Report)
	})

	// 从文章中添加标题（Title）到ctx。
	w.OnHTML("a[class=\"heading md story_detail_title\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 从/question-for-expert的文章中添加标题（Title）到ctx。
	w.OnHTML(" div[class=\"heading text-center heading_main faq_title \"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 从News文章中添加作者（Authors）到ctx。
	w.OnHTML("div[class=\"article_meta story_detail_meta\"]>span[class=\"user article_author\"]>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 从“名字, 职务”中获取名字。
		ctx.Authors = append(ctx.Authors, cutOutName(element.Text))
	})

	// 从Report文章中添加作者（Authors）到ctx。
	w.OnHTML("div[class=\"author_name\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 从“名字, 职务”中获取名字。
		ctx.Authors = append(ctx.Authors, cutOutName(element.Text))
	})

	// 从/question-for-expert文章中添加作者（Authors）到ctx。
	w.OnHTML("div[class=\"content clear-block\"]>.wrap > :nth-child(2)", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 去除掉名字前面的 "Replied by "。
		outName := strings.Replace(element.Text, "Replied by ", "", 1)
		// 从“名字, 职务”中获取名字。
		ctx.Authors = append(ctx.Authors, cutOutName(outName))
	})

	// 从文章中添加正文（Content）到ctx。
	w.OnHTML("div[class=\"article_brief article_full_description \"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 从/question-for-expert的文章中添加正文（Content）到ctx。
	w.OnHTML(".node>div[class=\"content clear-block\"]>.wrap", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 添加标签（Tags）到ctx（仅带/article路径的News文章含有）
	w.OnHTML(".keywords_tags", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		// 获取所有tag子元素的文本组成的[]string。
		tagsList := element.ChildTexts("li > a")
		ctx.Tags = tagsList
	})

	// 从文章中添加观看数（ViewCount）和评论数（CommentCount）到ctx。
	w.OnHTML("div[class=\"article_meta story_detail_meta\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.ViewCount, _ = strconv.Atoi(element.ChildText(".article_view>strong"))
		ctx.CommentCount, _ = strconv.Atoi(element.ChildText(".article_comment>strong"))
	})

	w.OnHTML(".post_date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
}
