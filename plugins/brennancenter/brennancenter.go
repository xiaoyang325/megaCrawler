package brennancenter

import (
   "github.com/gocolly/colly/v2"
   "megaCrawler/Crawler"
   "strings"
)

func init() {
   w := Crawler.Register("brennancenter", "布伦南司法中心",
      "https://www.brennancenter.org/")

   w.SetStartingUrls([]string{
      "https://www.brennancenter.org/issues/reform-money-politics/foreign-spending",
      "https://www.brennancenter.org/issues/defend-our-elections/independent-state-legislature-theory",
      "https://www.brennancenter.org/issues/strengthen-our-courts/scotus-federal-courts",
      "https://www.brennancenter.org/issues/end-mass-incarceration/cutting-jail-prison-populations",
      "https://www.brennancenter.org/issues/end-mass-incarceration/accurate-crime-data",
      "https://www.brennancenter.org/issues/bolster-checks-balances/effective-congress",
      "https://www.brennancenter.org/issues/advance-constitutional-change/electoral-college-reform",
      "https://www.brennancenter.org/issues/advance-constitutional-change/equal-rights-amendment",
      "https://www.brennancenter.org/issues/advance-constitutional-change/first-amendment",
      "https://www.brennancenter.org/issues/advance-constitutional-change/second-amendment",
      "https://www.brennancenter.org/series/abortion-rights-are-essential-democracy",
      "https://www.brennancenter.org/series/punitive-excess",
      "https://www.brennancenter.org/our-work/protests-insurrection-and-second-amendment",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/voting-reform/automatic-voter-registration",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/voting-reform/strengthening-voting-rights-act",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/voting-reform/ballot-design",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/voting-reform/election-administration",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/voting-reform/state-voting-laws",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/voting-reform/people-act-democracy-reform",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/voting-rights-restoration/state-reform",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/voting-rights-restoration/disenfranchisement-laws",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/voting-reform/state-voting-laws",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/vote-suppression/voter-purges",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/vote-suppression/myth-voter-fraud",
      "https://www.brennancenter.org/issues/ensure-every-american-can-vote/vote-suppression/voter-id",
      "https://www.brennancenter.org/issues/defend-our-elections/election-security/voting-machines-infrastructure",
      "https://www.brennancenter.org/our-work/responding-coronavirus-crisis",
      "https://www.brennancenter.org/issues/defend-our-elections/election-security/post-election-audits",
      "https://www.brennancenter.org/issues/defend-our-elections/election-security/funding-election-security",
      "https://www.brennancenter.org/issues/gerrymandering-fair-representation/redistricting/fight-fair-maps",
      "https://www.brennancenter.org/issues/gerrymandering-fair-representation/redistricting/redistricting-reform",
      "https://www.brennancenter.org/issues/gerrymandering-fair-representation/redistricting/redistricting-courts",
      "https://www.brennancenter.org/issues/gerrymandering-fair-representation/fair-accurate-census/2020-census-litigation",
      "https://www.brennancenter.org/issues/gerrymandering-fair-representation/fair-accurate-census/citizenship-question",
      "https://www.brennancenter.org/issues/gerrymandering-fair-representation/fair-accurate-census/census-confidentiality",
      "https://www.brennancenter.org/issues/reform-money-politics/influence-big-money/super-pacs-coordination",
      "https://www.brennancenter.org/issues/reform-money-politics/influence-big-money/dark-money",
      "https://www.brennancenter.org/issues/reform-money-politics/influence-big-money/enforcement-fec",
      "https://www.brennancenter.org/issues/reform-money-politics/public-campaign-financing/small-donor-public-financing",
      "https://www.brennancenter.org/issues/reform-money-politics/public-campaign-financing/campaign-finance-new-york-state",
      "https://www.brennancenter.org/issues/reform-money-politics/campaign-finance-courts/citizens-united",
      "https://www.brennancenter.org/issues/reform-money-politics/foreign-spending",
      "https://www.brennancenter.org/issues/strengthen-our-courts/promote-fair-courts/choosing-state-court-judges",
      "https://www.brennancenter.org/issues/strengthen-our-courts/promote-fair-courts/money-judicial-elections",
      "https://www.brennancenter.org/issues/strengthen-our-courts/promote-fair-courts/diversity-bench",
      "https://www.brennancenter.org/issues/strengthen-our-courts/promote-fair-courts/judicial-ethics-recusal",
      "https://www.brennancenter.org/issues/strengthen-our-courts/promote-fair-courts/assaults-courts",
      "https://www.brennancenter.org/issues/end-mass-incarceration/changing-incentives/accountable-private-prisons",
      "https://www.brennancenter.org/issues/end-mass-incarceration/changing-incentives/prosecutorial-reform",
      "https://www.brennancenter.org/issues/end-mass-incarceration/changing-incentives/fees-fines",
      "https://www.brennancenter.org/issues/protect-liberty-security/government-targeting-minority-communities/muslim-ban-extreme",
      "https://www.brennancenter.org/issues/protect-liberty-security/government-targeting-minority-communities/domestic-terrorism-hate",
      "https://www.brennancenter.org/issues/protect-liberty-security/government-targeting-minority-communities/countering-violent",
      "https://www.brennancenter.org/issues/protect-liberty-security/privacy-free-expression/foreign-intelligence-surveillance",
      "https://www.brennancenter.org/issues/protect-liberty-security/privacy-free-expression/policing-technology",
      "https://www.brennancenter.org/issues/protect-liberty-security/social-media/government-social-media-surveillance",
      "https://www.brennancenter.org/issues/protect-liberty-security/social-media/schools-social-media-surveillance",
      "https://www.brennancenter.org/issues/protect-liberty-security/social-media/police-social-media-surveillance",
      "https://www.brennancenter.org/issues/protect-liberty-security/transparency-oversight/secret-law",
      "https://www.brennancenter.org/issues/protect-liberty-security/transparency-oversight/rethinking-intelligence",
      "https://www.brennancenter.org/issues/bolster-checks-balances/ethics-rule-law/national-task-force-rule-law-democracy",
      "https://www.brennancenter.org/issues/bolster-checks-balances/executive-power/domestic-deployment-military",
      "https://www.brennancenter.org/issues/bolster-checks-balances/executive-power/emergency-powers",
   })

   // 从 Index 访问 Work & Resources 的 banner 的 Report
   w.OnHTML(".half-banner__text .half-banner__title a", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         w.Visit("https://www.brennancenter.org" + element.Attr("href"), Crawler.Report)
      })

   // 从 Index 访问 Work & Resources 的 Report
   w.OnHTML(".collection-teaser__text .collection-teaser__title a", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         w.Visit("https://www.brennancenter.org" + element.Attr("href"), Crawler.Report)
      })

   // 从 Index 访问 Work & Resources 的 Report
   w.OnHTML(".link__title a", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         w.Visit("https://www.brennancenter.org" + element.Attr("href"), Crawler.Report)
      })

   // 从 Index 访问 Work & Resources 的 Report
   w.OnHTML(".teaser__title a", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         w.Visit("https://www.brennancenter.org" + element.Attr("href"), Crawler.Report)
      })

   // 从 Index 访问 /news-analysis 的 Report
   w.OnHTML(".search-result-link__title > a", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         w.Visit("https://www.brennancenter.org" + element.Attr("href"), Crawler.Report)
      })


   // 从 Index 访问 Experts
   w.OnHTML(".expert__name > a", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         w.Visit("https://www.brennancenter.org" + element.Attr("href"), Crawler.Expert)
      })

   // 获取 Report 的 CategoryText
   w.OnHTML(".page-info-header__tophat-area .page-info-header__tophat", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.CategoryText = element.Text
      })

   // 获取 Report 的 Title
   w.OnHTML(".page-info-header__title span", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Title = element.Text
      })

   // 获取 Report 的 Description
   w.OnHTML(".page-info-header__description", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Description = element.Text
      })

   // 获取 Report 的 Authors
   w.OnHTML(".page-info-header__author-title", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Authors = append(ctx.Authors, element.Text)
      })

   // 获取 Report 的 Publication Time
   w.OnHTML(".page-info-header__dates", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         if (strings.Contains(element.ChildText(":nth-child(2)"), "Published: ")) {
            time := strings.ReplaceAll(element.ChildText(":nth-child(2)"), "Published: ", "")
            ctx.PublicationTime = strings.TrimSpace(time)
         } else
         {
            ctx.PublicationTime = strings.TrimSpace(element.ChildText(":nth-child(1)"))
         }
      })

   // 获取 Report 的 Content
   w.OnHTML(".page-body", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         if (ctx.PageType == Crawler.Report) {
            ctx.Content = element.ChildText("p, h2")
         }
      })

   // 获取 Expert 的 Name
   w.OnHTML(".page-bio-header__title", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Name = element.Text
      })

   // 获取 Expert 的 TwitterId
   w.OnHTML(".page-bio-header__twitter > span > a", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.TwitterId = strings.Replace(element.Attr("href"), "https://twitter.com/", "", 1)
      })

   // 获取 Expert 的 CategoryText
   w.OnHTML(".page-bio-header__tophat", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.CategoryText = strings.TrimSpace(element.Text)
      })

   // 获取 Expert 的 Title
   w.OnHTML(".page-bio-header__role", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Title = strings.TrimSpace(element.Text)
      })

   // 获取 Expert 的 Description
   w.OnHTML(".page-body", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         if (ctx.PageType == Crawler.Expert) {
            ctx.Description = strings.TrimSpace(element.ChildText("p, h2"))
         }
      })

   // 获取 Expert 的 Email
   w.OnHTML(".related-bio-group__contact-content :nth-child(2)", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Email = strings.TrimSpace(element.Text)
      })

   // 获取 Expert 的 Phone
   w.OnHTML(".related-bio-group__contact-content :nth-child(4)", 
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Phone = strings.TrimSpace(element.Text)
      })
}