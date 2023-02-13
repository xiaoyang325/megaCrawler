// Package production contain plugins that are ready for production
package production

import (
	"megaCrawler/crawlers"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

// 从名字列表中分离出所有名字，并加入到[]string中返回。
func cutOutNames(nameStr string) []string {
	reg := regexp.MustCompile("(?U)[(].*[)]")
	outAnd := regexp.MustCompile("and")

	result := reg.ReplaceAllString(nameStr, "")
	result2 := outAnd.ReplaceAllString(result, ",")
	pts := strings.Split(result2, ",")

	for index, value := range pts {
		pts[index] = strings.TrimSpace(value)
	}

	return pts
}

func init() {
	w := crawlers.Register("pile",
		"彼得森国际经济研究所", "https://www.piie.com")

	w.SetStartingURLs([]string{
		"https://www.piie.com/research/organizations/association-southeast-asian-nations",
		"https://www.piie.com/research/economic-issues/coronavirus",
		"https://www.piie.com/research/economic-issues/economic-outlook",
		"https://www.piie.com/research/economic-issues/environment",
		"https://www.piie.com/research/economic-issues/gender",
		"https://www.piie.com/research/economic-issues/globalization",
		"https://www.piie.com/research/economic-issues/inequality",
		"https://www.piie.com/research/economic-issues/inflation",
		"https://www.piie.com/research/economic-issues/labor",
		"https://www.piie.com/research/finance/monetary-policy",
		"https://www.piie.com/research/trade-investment/sanctions",
		"https://www.piie.com/research/economic-issues/technology",
		"https://www.piie.com/research/trade-investment",
		"https://www.piie.com/research/organizations/world-trade-organization",
		"https://www.piie.com/research/china",
		"https://www.piie.com/research/european-union",
		"https://www.piie.com/research/india",
		"https://www.piie.com/research/japan",
		"https://www.piie.com/research/latin-america-caribbean",
		"https://www.piie.com/research/mexico",
		"https://www.piie.com/research/middle-east-north-africa",
		"https://www.piie.com/research/south-korea",
		"https://www.piie.com/research/ukraine",
		"https://www.piie.com/research/united-kingdom",
		"https://www.piie.com/research/united-states",
		"https://www.piie.com/research/argentina",
		"https://www.piie.com/research/brazil",
		"https://www.piie.com/research/burma",
		"https://www.piie.com/research/east-asia-pacific",
		"https://www.piie.com/research/egypt",
		"https://www.piie.com/research/europe-central-asia",
		"https://www.piie.com/research/former-soviet-economies",
		"https://www.piie.com/research/france",
		"https://www.piie.com/research/germany",
		"https://www.piie.com/research/greece",
		"https://www.piie.com/research/iran",
		"https://www.piie.com/research/ireland",
		"https://www.piie.com/research/italy",
		"https://www.piie.com/research/middle-east-north-africa",
		"https://www.piie.com/research/north-korea",
		"https://www.piie.com/research/puerto-rico",
		"https://www.piie.com/research/russia",
		"https://www.piie.com/research/south-asia",
		"https://www.piie.com/research/sub-saharan-africa",
		"https://www.piie.com/research/syria",
		"https://www.piie.com/research/trade-investment/agriculture",
		"https://www.piie.com/research/trade-investment/commodities",
		"https://www.piie.com/research/trade-investment/competition",
		"https://www.piie.com/research/trade-investment/cptpp-tpp",
		"https://www.piie.com/research/trade-investment/disputes",
		"https://www.piie.com/research/trade-investment/foreign-direct-investment",
		"https://www.piie.com/research/trade-investment/free-trade-agreements",
		"https://www.piie.com/research/trade-investment/intellectual-property-rights",
		"https://www.piie.com/research/trade-investment/manufacturing",
		"https://www.piie.com/research/trade-investment/multinational-corporations",
		"https://www.piie.com/research/trade-investment/protectionism",
		"https://www.piie.com/research/trade-investment/services",
		"https://www.piie.com/research/trade-investment/trade-deficit",
		"https://www.piie.com/research/trade-investment/trade-policy",
		"https://www.piie.com/research/trade-investment/transatlantic-trade-and-investment-partnership",
		"https://www.piie.com/research/trade-investment/us-china-trade-war",
		"https://www.piie.com/research/trade-investment/usmca-nafta",
		"https://www.piie.com/research/political-economy/corruption",
		"https://www.piie.com/research/political-economy/fiscal-policy",
		"https://www.piie.com/research/political-economy/foreign-aid",
		"https://www.piie.com/research/political-economy/governance",
		"https://www.piie.com/research/political-economy/government",
		"https://www.piie.com/research/political-economy/health",
		"https://www.piie.com/research/political-economy/human-rights",
		"https://www.piie.com/research/political-economy/migration",
		"https://www.piie.com/research/political-economy/nuclear",
		"https://www.piie.com/research/political-economy/politics",
		"https://www.piie.com/research/political-economy/security",
		"https://www.piie.com/research/political-economy/washington-consensus",
		"https://www.piie.com/research/organizations/asia-pacific-economic-cooperation",
		"https://www.piie.com/research/organizations/association-southeast-asian-nations",
		"https://www.piie.com/research/organizations/central-banks",
		"https://www.piie.com/research/organizations/european-central-bank",
		"https://www.piie.com/research/organizations/european-commission",
		"https://www.piie.com/research/organizations/export-import-bank",
		"https://www.piie.com/research/organizations/g20",
		"https://www.piie.com/research/organizations/g7",
		"https://www.piie.com/research/organizations/international-monetary-fund",
		"https://www.piie.com/research/organizations/g8",
		"https://www.piie.com/research/organizations/organization-economic-cooperation-and-development",
		"https://www.piie.com/research/organizations/peoples-bank-china",
		"https://www.piie.com/research/organizations/united-nations",
		"https://www.piie.com/research/organizations/us-federal-reserve",
		"https://www.piie.com/research/organizations/world-bank",
		"https://www.piie.com/research/organizations/world-trade-organization",
		"https://www.piie.com/research/finance/capital-markets",
		"https://www.piie.com/research/finance/currency",
		"https://www.piie.com/research/finance/currency-manipulation",
		"https://www.piie.com/research/finance/financial-crises",
		"https://www.piie.com/research/finance/monetary-policy",
		"https://www.piie.com/research/finance/regulations",
		"https://www.piie.com/research/economic-issues/economic-policy-pandemic-age",
		"https://www.piie.com/research/finance/taxes",
		"https://www.piie.com/research/economic-issues/education",
		"https://www.piie.com/research/economic-issues/emerging-markets",
		"https://www.piie.com/research/economic-issues/energy",
		"https://www.piie.com/research/economic-issues/fiscal-deficit",
		"https://www.piie.com/research/economic-issues/growth",
		"https://www.piie.com/research/economic-issues/infrastructure",
		"https://www.piie.com/research/economic-issues/productivity",
		"https://www.piie.com/research/economic-issues/rebuilding-global-economy",
		"https://www.piie.com/research/economic-issues/secular-stagnation",
		"https://www.piie.com/research/economic-issues/womens-economic-empowerment-research-initiative",
		"https://www.piie.com/events",
		"https://www.piie.com/blogs",
		"https://www.piie.com/research/publications",
		"https://www.piie.com/research/commentary/op-eds",
		"https://www.piie.com/research/commentary/testimonies",
		"https://www.piie.com/research/commentary/speeches-papers",
		"https://www.piie.com/experts/podcasts/trade-talks",
	})

	// 从翻页器中获取新的Index
	w.OnHTML("a[title=\"Go to next page\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		toURL := strings.Split(ctx.URL, "?")[0] + element.Attr("href")
		w.Visit(toURL, crawlers.Index)
	})

	// 从Index中进入文章（情况一）
	w.OnHTML(".view__row>article>.teaser__body>.teaser__title>span>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		toURL := "https://www.piie.com/" + element.Attr("href")
		if strings.Contains(toURL, "/blogs/") || strings.Contains(toURL, "/events/") {
			w.Visit(element.Attr("href"), crawlers.News)
		} else {
			w.Visit(element.Attr("href"), crawlers.Report)
		}
	})

	// 从Index中进入文章（情况二）
	w.OnHTML(".view__row>article>.teaser__body>.teaser__title>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		if strings.Contains(element.Attr("href"), "https:") {
			// 对于网站外的链接，什么也不做。
		} else {
			toURL := "https://www.piie.com/" + element.Attr("href")
			if strings.Contains(toURL, "/blogs/") || strings.Contains(toURL, "/events/") {
				w.Visit(element.Attr("href"), crawlers.News)
			} else {
				w.Visit(element.Attr("href"), crawlers.Report)
			}
		}
	})

	// 从文章中添加标题（Title）到ctx。（/events/）
	w.OnHTML(".hero-banner-event__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 从文章中添加标题（Title）到ctx。（/blogs/）（/publications/）（/research/）（/commentary/）
	w.OnHTML(".hero-banner-publication__title", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Title = element.Text
	})

	// 从文章中添加位置（Location）到ctx。（/events/）
	w.OnHTML(".location>span", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Location = element.Text
	})

	// 从文章中添加作者（Authors）到ctx。（/events/）
	w.OnHTML(".hero-banner-event__speakers>div>.field__item>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, cutOutNames(element.Text)...)
	})

	// 从文章中添加作者（Authors）到ctx。（/blogs/）（/publications/）（/research/）（/commentary/）
	w.OnHTML(".hero-banner-publication__authors>div>div>p", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, cutOutNames(element.Text)...)
	})

	// 从文章中添加作者（Authors）到ctx。（/blogs/）（/publications/）（/research/）（/commentary/）
	w.OnHTML(".hero-banner-publication__authors>.author-list", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, cutOutNames(element.Text)...)
	})

	// 从文章中添加作者（Authors）到ctx。（/events/）
	w.OnHTML(".hero-banner-event__speakers>.author-list>.author-list__author>.author-list__link", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
	})

	// 从文章中添加正文（Content）到ctx。（/events/）
	w.OnHTML(".content-block__inner", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 从文章中添加正文（Content）到ctx。（/blogs/）（/publications/）（/research/）（/commentary/）
	w.OnHTML(".content-block__inner>div>div", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Content = element.Text
	})

	// 从文章中添加文件（File）到ctx。（/publications/）（/commentary/）
	w.OnHTML(".download-button>a", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := "https://www.piie.com" + element.Attr("href")
		ctx.File = append(ctx.File, fileURL)
	})

	// 从文章中添加文件（File）到ctx。（/events/）
	w.OnHTML("a[type=\"application/pdf\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		fileURL := "https://www.piie.com" + element.Attr("href")
		ctx.File = append(ctx.File, fileURL)
	})

	w.OnHTML("meta[property=\"og:description\"]", func(element *colly.HTMLElement, ctx *crawlers.Context) {
		ctx.Description = element.Attr("content")
	})
}
