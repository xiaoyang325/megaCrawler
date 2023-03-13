package production

import (
	"strings"

	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("cgai", "全球事务研究所", "https://www.cgai.ca/")

	w.SetStartingURLs([]string{"https://www.cgai.ca/defence_innovation",
		"https://www.cgai.ca/defence_policy",
		"https://www.cgai.ca/defence_resources",
		"https://www.cgai.ca/defence_operations",
		"https://www.cgai.ca/procurement",
		"https://www.cgai.ca/natotag",
		"https://www.cgai.ca/north_america_norad",
		"https://www.cgai.ca/cyber_tech",
		"https://www.cgai.ca/hybrid_threats",
		"https://www.cgai.ca/space",
		"https://www.cgai.ca/intelligence",
		"https://www.cgai.ca/terrorism",
		"https://www.cgai.ca/counter_insurgency_operations",
		"https://www.cgai.ca/wmds",
		"https://www.cgai.ca/borders",
		"https://www.cgai.ca/environment_energy",
		"https://www.cgai.ca/health",
		"https://www.cgai.ca/food_water",
		"https://www.cgai.ca/international_trade",
		"https://www.cgai.ca/natural_resources",
		"https://www.cgai.ca/supply_chain",
		"https://www.cgai.ca/human_rights",
		"https://www.cgai.ca/international_law",
		"https://www.cgai.ca/migration",
		"https://www.cgai.ca/development",
		"https://www.cgai.ca/international_institutions",
		"https://www.cgai.ca/international_politics",
		"https://www.cgai.ca/elections",
		"https://www.cgai.ca/canada",
		"https://www.cgai.ca/united_states",
		"https://www.cgai.ca/mexico",
		"https://www.cgai.ca/caribbean_latin_america",
		"https://www.cgai.ca/south_america",
		"https://www.cgai.ca/arctic",
		"https://www.cgai.ca/western_europe",
		"https://www.cgai.ca/eastern_europe",
		"https://www.cgai.ca/central_southern_europe",
		"https://www.cgai.ca/turkey",
		"https://www.cgai.ca/iran",
		"https://www.cgai.ca/russia",
		"https://www.cgai.ca/afghanistan",
		"https://www.cgai.ca/iraq",
		"https://www.cgai.ca/syria",
		"https://www.cgai.ca/israel",
		"https://www.cgai.ca/saudi_arabia",
		"https://www.cgai.ca/africa",
		"https://www.cgai.ca/australia",
		"https://www.cgai.ca/china",
		"https://www.cgai.ca/japan",
		"https://www.cgai.ca/the_koreas",
		"https://www.cgai.ca/indochina",
		"https://www.cgai.ca/india",
		"https://www.cgai.ca/pakistan",
		"https://www.cgai.ca/taiwan",
		"https://www.cgai.ca/global",
		"https://www.cgai.ca/policy_papers",
		"https://www.cgai.ca/policy_perspectives",
		"https://www.cgai.ca/primers",
		"https://www.cgai.ca/the_global_exchange",
		"https://www.cgai.ca/2022_defence_and_security_series",
		"https://www.cgai.ca/2022_supply_chains_series",
		"https://www.cgai.ca/2021_taiwan_series",
		"https://www.cgai.ca/2021_indo_pacific_series",
		"https://www.cgai.ca/2020_energy_security_series",
		"https://www.cgai.ca/2019_lng_series",
		"https://www.cgai.ca/2018_international_trade_series",
		"https://www.cgai.ca/2017_nato_series",
		"https://www.cgai.ca/2017_energy_series",
		"https://www.cgai.ca/2017_foreign_policy_series",
		"https://www.cgai.ca/2016_defence_policy_series",
		"https://www.cgai.ca/commentary",
		"https://www.cgai.ca/committee_testimony",
		"https://www.cgai.ca/books",
		"https://www.cgai.ca/energy_security_forum",
		"https://www.cgai.ca/book_reviews",
		"https://www.cgai.ca/the_global_exchange_podcast",
		"https://www.cgai.ca/defence_deconstructed_podcast",
		"https://www.cgai.ca/battle_rhythm_podcast",
		"https://www.cgai.ca/conseils_de_securite",
		"https://www.cgai.ca/energy_security3_podcast"})

	// 尝试寻找下载pdf的按钮，并如果存在则将页面类型转换为报告
	w.OnHTML("a.button", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), ".pdf") {
			ctx.File = append(ctx.File, element.Attr("href"))
			ctx.PageType = crawlers.Report
		}
	})

	// 从翻页器获取链接并访问
	w.OnHTML("div.pagination>ul>li>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})
	// 从index访问新闻
	w.OnHTML("h3 > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("td.top>p>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	// report.title
	w.OnHTML("h2.headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// report.author
	w.OnHTML("#intro > div > p > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("#intro > div.text-content > p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})

	// 内含Expert
	w.OnHTML("#intro > div > p:nth-child(4) > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})
	// expert.Name
	w.OnHTML("h2.headline", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = element.Text
	})

	// expert.area
	w.OnHTML("div#fellowsheadertag>p>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Area = ctx.Area + "," + element.Text
	})
	// expert.description
	w.OnHTML("div.text-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	// report .content
	w.OnHTML("div.text-content", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
}
