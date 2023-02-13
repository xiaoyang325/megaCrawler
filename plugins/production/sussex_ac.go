package production

import (
	"megaCrawler/crawlers"
	"strings"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("sussex_ac", "萨塞克斯大学腐败问题研究所", "https://www.sussex.ac.uk/")

	w.SetStartingUrls([]string{
		"https://www.sussex.ac.uk/news/research",
		"https://www.sussex.ac.uk/news/people",
		"https://www.sussex.ac.uk/news/university",
		"https://www.sussex.ac.uk/research/centres/centre-for-study-of-corruption/about/team",
	})

	// 访问下一页 Index ***
	w.OnHTML(`[class="paginationControl bottom clear"] span[name="next"] > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 访问 News 从 Index ***
	w.OnHTML(`div[class="row gutter-small row-wrap"] > div > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})

	// 访问 Expert 从 Index ***
	w.OnHTML(`[class="small-12 large-9 large-push-3 columns wysiwyg content-region"] div.content > p:nth-child(4) > a`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})

	// 获取 Title ***
	w.OnHTML(`h2.lineless`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Title ***
	w.OnHTML(`[class="userInfo__userInfoGroup___Wv_dB"] div > div`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = strings.TrimSpace(element.Text)
	})

	// 获取 Email ***
	w.OnHTML(`.oneLineText__oneLineText____Igu4`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Email = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Description ***
	w.OnHTML(`.whiteBox__body___nZQwU > p`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = strings.TrimSpace(element.Text)
	})

	// 获取 Phone ***
	w.OnHTML(`#app > div > main > div.profile__widthRestrictedContent___GTpH5.profile__profileComponentInnerContainer___QqE0i > div.profile__sidebar___cm8bT > div > div > div.userInfo__userInfoDesktop___Li7iR > div:nth-child(6) > ul:nth-child(1) > li > div.iconAndContentRow__content___ZD5ge.iconAndContentRow__leftMargin___bpZTS.iconAndContentRow__topMargin___SN4Wk`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Phone = strings.TrimSpace(element.Text)
	})

	// 获取 Expert's Tags ***
	w.OnHTML(`.tag__body___B87S8 > .tag__value___ePSCW`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
	})

	// 获取 TwitterId ***
	w.OnHTML(`[aria-label="Twitter profile"]`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		raw := element.Attr("href")
		raw = strings.Replace(raw, "https://twitter.com/", "", 1)
		ctx.TwitterId = strings.TrimSpace(raw)
	})

	// 获取 Name ***
	w.OnHTML(`.hero__header___xfv2U`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = strings.TrimSpace(element.Text)
	})

	// 获取 PublicationTime ***
	w.OnHTML(`#main > .row > div > p:nth-child(4)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		date := strings.Replace(element.Text, "Last updated:", "", 1)
		ctx.PublicationTime = strings.TrimSpace(date)
	})

	// 获取 CategoryText ***
	w.OnHTML(`#main .row h1`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.CategoryText = strings.TrimSpace(element.Text)
	})

	// 获取 Authors ***
	w.OnHTML(`#main > .row > div > p:nth-child(3)`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		name := strings.Replace(element.Text, "By:", "", 1)
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(name))
	})

	// 获取 Content ***
	w.OnHTML(`.social-wrapper > #news-item`, func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = strings.TrimSpace(element.ChildText("p, h1, h2, h3"))
	})
}
