package production

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func init() {
	w := crawlers.Register("carnegieendowment", "卡内基国际和平基金会",
		"https://carnegieendowment.org/")

	w.SetStartingUrls([]string{
		"https://www.nupi.no/en/our-research/topics/defence-and-security/intelligence",
		"https://www.nupi.no/en/our-research/topics/defence-and-security/cyber",
		"https://www.nupi.no/en/our-research/topics/defence-and-security/nato",
		"https://www.nupi.no/en/our-research/topics/defence-and-security/terrorism-and-extremism",
		"https://www.nupi.no/en/our-research/topics/defence-and-security/security-policy",
		"https://www.nupi.no/en/our-research/topics/defence-and-security/defence",
		"https://www.nupi.no/en/our-research/topics/natural-resources-and-climate/oceans",
		"https://www.nupi.no/en/our-research/topics/natural-resources-and-climate/energy",
		"https://www.nupi.no/en/our-research/topics/natural-resources-and-climate/climate",
		"https://www.nupi.no/en/our-research/topics/global-governance/the-african-union-au",
		"https://www.nupi.no/en/our-research/topics/global-governance/human-rights",
		"https://www.nupi.no/en/our-research/topics/global-governance/united-nations",
		"https://www.nupi.no/en/our-research/topics/global-governance/the-eu",
		"https://www.nupi.no/en/our-research/topics/global-governance/international-organizations",
		"https://www.nupi.no/en/our-research/topics/global-governance/governance",
		"https://www.nupi.no/en/our-research/topics/diplomacy-and-foreign-policy/foreign-policy",
		"https://www.nupi.no/en/our-research/topics/diplomacy-and-foreign-policy/diplomacy",
		"https://www.nupi.no/en/our-research/topics/diplomacy-and-foreign-policy/development-policy",
		"https://www.nupi.no/en/our-research/topics/peace-crisis-and-conflict/pandemics",
		"https://www.nupi.no/en/our-research/topics/peace-crisis-and-conflict/conflict",
		"https://www.nupi.no/en/our-research/topics/peace-crisis-and-conflict/humanitarian-issues",
		"https://www.nupi.no/en/our-research/topics/peace-crisis-and-conflict/peace-operations",
		"https://www.nupi.no/en/our-research/topics/global-economy/international-investments",
		"https://www.nupi.no/en/our-research/topics/global-economy/regional-integration",
		"https://www.nupi.no/en/our-research/topics/global-economy/trade",
		"https://www.nupi.no/en/our-research/topics/global-economy/economic-growth",
		"https://www.nupi.no/en/our-research/topics/global-economy/international-economics",
		"https://www.nupi.no/en/our-research/topics/theory-and-method/historical-international-relations",
		"https://www.nupi.no/en/our-research/topics/theory-and-method/comparative-methods",
		"https://www.nupi.no/en/our-research/regions/the-nordic-countries",
		"https://www.nupi.no/en/our-research/regions/europe",
		"https://www.nupi.no/en/our-research/regions/russia-and-eurasia",
		"https://www.nupi.no/en/our-research/regions/the-middle-east-and-north-africa",
		"https://www.nupi.no/en/our-research/regions/africa",
		"https://www.nupi.no/en/our-research/regions/asia",
		"https://www.nupi.no/en/our-research/regions/the-arctic",
		"https://www.nupi.no/en/our-research/regions/south-and-central-america",
		"https://www.nupi.no/en/our-research/regions/north-america",
		"https://www.nupi.no/en/our-research/research-centres/centre-for-ocean-governance",
		"https://www.nupi.no/en/our-research/research-centres/nupi-centre-for-european-studies",
		"https://www.nupi.no/en/our-research/research-centres/centre-for-energy-research",
		"https://www.nupi.no/en/our-research/research-centres/nupi-s-centre-for-digitalization-and-cyber-security-studies",
		"https://www.nupi.no/en/our-research/research-centres/nupi-center-for-asia-research",
		"https://www.nupi.no/en/our-research/research-centres/consortium-for-research-on-terrorism-and-international-crime",
		"https://www.nupi.no/en/our-research/research-centres/nupi-center-for-un-and-global-governance",
		"https://www.nupi.no/en/our-research/research-centres/the-taxcapdev-network",
		"https://www.nupi.no/en/our-research/research-centres/norwegian-centre-for-humanitarian-studies",
	})
	// 从翻页器获取链接并访问
	w.OnHTML("a.page-link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Index)
	})

	// 从index访问新闻
	w.OnHTML("div.media-body>h5>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	// 从index访问新闻
	w.OnHTML("div.ezrichtext-field>p>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	// 从index访问新闻
	w.OnHTML("div.card-body>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	w.OnHTML("section.events > div > div > a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.News)
	})
	// report.title
	w.OnHTML("body > div.container > div.row.main.content > div > div:nth-child(2) > div > h1 > font > font", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	// report.title
	w.OnHTML("div.content>div>div>h1", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	w.OnHTML("h1.pt-4", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// report .content
	w.OnHTML("div.article-body", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	// report.publish time
	w.OnHTML("body > div.container > div.row.main.content > div > div:nth-child(2) > div > div > span > font > font", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	w.OnHTML("span.pub-date", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	w.OnHTML("div.fact>h4", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.PublicationTime = element.Text
	})
	// report.description
	w.OnHTML("div.eztext-field", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Text
	})

	// report.author
	w.OnHTML("section.person-section.py-3 > div > div > div > div > div.front > div > div > h5 > strong > a > font > font", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("div.byline>ul>li.list-inline-item>a.author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	w.OnHTML("a.author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, element.Text)
	})
	// report .content
	w.OnHTML("div.col-md-8", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 内含Expert
	w.OnHTML("a.author", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		w.Visit(element.Attr("href"), crawlers.Expert)
	})
	// expert.Name
	w.OnHTML("h2.resume-name", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Name = element.Text
	})
	// expert.title
	w.OnHTML("span.ezstring-field", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})
	// expert.description
	w.OnHTML("div.text-start>div.ezrichtext-field", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})
	// expert.link
	w.OnHTML("li.email>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
	w.OnHTML("li.phone", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Link = append(ctx.Link, element.Attr("href"))
	})
}
