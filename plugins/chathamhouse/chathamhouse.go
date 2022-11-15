package chathamhouse

import (
   "github.com/gocolly/colly/v2"
   "megaCrawler/Crawler"
   "strings"
)

func init() {
   w := Crawler.Register("chathamhouse", "查塔姆研究所", "https://www.chathamhouse.org/")

   w.SetStartingUrls([]string{
      "https://www.chathamhouse.org/topics/defence-and-security",
      "https://www.chathamhouse.org/topics/economics-and-trade",
      "https://www.chathamhouse.org/topics/environment",
      "https://www.chathamhouse.org/topics/health",
      "https://www.chathamhouse.org/topics/institutions",
      "https://www.chathamhouse.org/topics/major-powers",
      "https://www.chathamhouse.org/topics/politics-and-law",
      "https://www.chathamhouse.org/topics/society",
      "https://www.chathamhouse.org/topics/technology",
      "https://www.chathamhouse.org/topics/access-healthcare",
      "https://www.chathamhouse.org/topics/african-union-au",
      "https://www.chathamhouse.org/topics/agriculture-and-food",
      "https://www.chathamhouse.org/topics/americas-international-role",
      "https://www.chathamhouse.org/topics/arms-control",
      "https://www.chathamhouse.org/topics/brics-economies",
      "https://www.chathamhouse.org/topics/chinas-belt-and-road-initiative-bri",
      "https://www.chathamhouse.org/topics/chinas-domestic-politics",
      "https://www.chathamhouse.org/topics/chinas-foreign-relations",
      "https://www.chathamhouse.org/topics/circular-economy",
      "https://www.chathamhouse.org/topics/civil-society",
      "https://www.chathamhouse.org/topics/climate-policy",
      "https://www.chathamhouse.org/topics/coronavirus-response",
      "https://www.chathamhouse.org/topics/cyber-security",
      "https://www.chathamhouse.org/topics/data-governance-and-security",
      "https://www.chathamhouse.org/topics/democracy-and-political-participation",
      "https://www.chathamhouse.org/topics/demographics-and-politics",
      "https://www.chathamhouse.org/topics/digital-and-social-media",
      "https://www.chathamhouse.org/topics/disinformation",
      "https://www.chathamhouse.org/topics/drugs-and-organized-crime",
      "https://www.chathamhouse.org/topics/energy-transitions",
      "https://www.chathamhouse.org/topics/european-defence",
      "https://www.chathamhouse.org/topics/european-union-eu",
      "https://www.chathamhouse.org/topics/future-work",
      "https://www.chathamhouse.org/topics/future-work",
      "https://www.chathamhouse.org/topics/gender-and-equality",
      "https://www.chathamhouse.org/topics/health-strategy",
      "https://www.chathamhouse.org/topics/human-rights-and-security",
      "https://www.chathamhouse.org/topics/international-criminal-justice",
      "https://www.chathamhouse.org/topics/international-finance-system",
      "https://www.chathamhouse.org/topics/international-monetary-fund-imf",
      "https://www.chathamhouse.org/topics/international-trade",
      "https://www.chathamhouse.org/topics/investment-africa",
      "https://www.chathamhouse.org/topics/managing-natural-resources",
      "https://www.chathamhouse.org/topics/north-atlantic-treaty-organization-nato",
      "https://www.chathamhouse.org/topics/peacekeeping-and-intervention",
      "https://www.chathamhouse.org/topics/radicalization",
      "https://www.chathamhouse.org/topics/refugees-and-migration",
      "https://www.chathamhouse.org/topics/technology-governance",
      "https://www.chathamhouse.org/topics/terrorism",
      "https://www.chathamhouse.org/topics/uks-global-role",
      "https://www.chathamhouse.org/topics/united-nations-un",
      "https://www.chathamhouse.org/topics/us-domestic-politics",
      "https://www.chathamhouse.org/topics/us-foreign-policy",
      "https://www.chathamhouse.org/topics/world-health-organization-who",
      "https://www.chathamhouse.org/topics/world-trade-organization-wto",
      "https://www.chathamhouse.org/regions/africa",
      "https://www.chathamhouse.org/regions/americas",
      "https://www.chathamhouse.org/regions/asia-pacific",
      "https://www.chathamhouse.org/regions/europe",
      "https://www.chathamhouse.org/regions/middle-east-and-north-africa",
      "https://www.chathamhouse.org/regions/russia-and-eurasia",
      "https://www.chathamhouse.org/regions/asia-pacific/afghanistan",
      "https://www.chathamhouse.org/regions/russia-and-eurasia/belarus",
      "https://www.chathamhouse.org/regions/americas/canada",
      "https://www.chathamhouse.org/regions/africa/central-africa",
      "https://www.chathamhouse.org/regions/americas/central-america-and-caribbean",
      "https://www.chathamhouse.org/regions/europe/central-and-eastern-europe",
      "https://www.chathamhouse.org/regions/russia-and-eurasia/central-asia",
      "https://www.chathamhouse.org/regions/asia-pacific/china",
      "https://www.chathamhouse.org/regions/africa/east-africa",
      "https://www.chathamhouse.org/regions/middle-east-and-north-africa/egypt",
      "https://www.chathamhouse.org/regions/europe/eurozone",
      "https://www.chathamhouse.org/regions/europe/france",
      "https://www.chathamhouse.org/regions/europe/germany",
      "https://www.chathamhouse.org/regions/middle-east-and-north-africa/gulf-states",
      "https://www.chathamhouse.org/regions/africa/horn-africa",
      "https://www.chathamhouse.org/regions/asia-pacific/india",
      "https://www.chathamhouse.org/regions/middle-east-and-north-africa/iran",
      "https://www.chathamhouse.org/regions/middle-east-and-north-africa/iraq",
      "https://www.chathamhouse.org/regions/middle-east-and-north-africa/israel-and-palestine",
      "https://www.chathamhouse.org/regions/asia-pacific/japan",
      "https://www.chathamhouse.org/regions/asia-pacific/korean-peninsula",
      "https://www.chathamhouse.org/regions/middle-east-and-north-africa/libya",
      "https://www.chathamhouse.org/regions/americas/mexico",
      "https://www.chathamhouse.org/regions/asia-pacific/pakistan",
      "https://www.chathamhouse.org/regions/russia-and-eurasia/russia",
      "https://www.chathamhouse.org/regions/americas/south-america",
      "https://www.chathamhouse.org/regions/asia-pacific/south-asia",
      "https://www.chathamhouse.org/regions/asia-pacific/southeast-asia",
      "https://www.chathamhouse.org/regions/africa/southern-africa",
      "https://www.chathamhouse.org/regions/middle-east-and-north-africa/syria-and-levant",
      "https://www.chathamhouse.org/regions/asia-pacific/pacific",
      "https://www.chathamhouse.org/regions/europe/turkey",
      "https://www.chathamhouse.org/regions/russia-and-eurasia/ukraine",
      "https://www.chathamhouse.org/regions/europe/united-kingdom",
      "https://www.chathamhouse.org/regions/americas/united-states-america",
      "https://www.chathamhouse.org/regions/africa/west-africa",
      "https://www.chathamhouse.org/regions/middle-east-and-north-africa/yemen",
      "https://www.chathamhouse.org/events",
   })

   // 从 Index 访问 News & Report
   w.OnHTML(`a[class="no-external-link-icon teaser__link"]`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         if (strings.Contains(element.Attr("href"), "/events")) {
            w.Visit(element.Attr("href"), Crawler.Report)
         } else {
            w.Visit(element.Attr("href"), Crawler.News)
         }
      })

   // 从 Index 访问 News & Report
   w.OnHTML(`.event-teaser > a`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         if (strings.Contains(element.Attr("href"), "/events")) {
            w.Visit(element.Attr("href"), Crawler.Report)
         } else {
            w.Visit(element.Attr("href"), Crawler.News)
         }
      })

   // 从 News 访问 Expert
   w.OnHTML(`h3[class="h4 person-teaser__title"] > a`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         w.Visit(element.Attr("href"), Crawler.Expert)
      })

   // 获取 Title
   w.OnHTML(`h1[class="h2 hero__title"] > span`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Title = strings.TrimSpace(element.Text)
      })

   // 获取 SubTitle
   w.OnHTML(`div.hero__subtitle > p`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.SubTitle = strings.TrimSpace(element.Text)
      })

   // 获取 Publication Time（/events/）
   w.OnHTML(".event-details__date",
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.PublicationTime = strings.TrimSpace(element.Text)
      })

   // 获取 Publication Time
   w.OnHTML(".hero__meta-date > time",
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.PublicationTime = strings.TrimSpace(element.Text)
      })

   // 获取 CategoryText（/events/）
   w.OnHTML(".event-details__type-text",
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.CategoryText = strings.TrimSpace(element.Text)
      })

   // 获取 CategoryText
   w.OnHTML(".hero__meta-label > span",
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.CategoryText = strings.TrimSpace(element.Text)
      })

   // 获取 Location（/events/）
   w.OnHTML("address.event-details__location > span",
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Location = strings.TrimSpace(element.Text)
      })

   // 获取 Tags
   w.OnHTML(".sidebar-taxonomy__list > li > a",
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Tags = append(ctx.Tags, strings.TrimSpace(element.Text))
      })

   // 获取 Authors
   w.OnHTML(`h3[class="h4 person-teaser__title"]`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
      })

   // 获取 Content
   w.OnHTML("body > div.dialog-off-canvas-main-canvas > div.layout-container > main > section.bg-white.section-bottom-padding.body-content > div > article > div > div.wysiwyg",
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Content = element.Text
      })

   // Expert 获取 Name
   w.OnHTML(`.person-bio__header > .person-bio__title`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Name = strings.TrimSpace(element.Text)
      })

   // Expert 获取 Title
   w.OnHTML(`.person-bio__header > .person-bio__role`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Title = strings.TrimSpace(element.Text)
      })

   // Expert 获取 Phone
   w.OnHTML(`.person-bio__contact-item:nth-child(1) > a >span`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         ctx.Phone = strings.TrimSpace(element.Text)
      })

   // Expert 获取 Email
   w.OnHTML(`.person-bio__contact-item:nth-child(2) > a`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         e := strings.Replace(element.Attr("href"), "mailto:", "", 1)
         ctx.Email = strings.TrimSpace(e)
      })

   // Expert 获取 TwitterId
   w.OnHTML(`.person-bio__contact-item:nth-child(2) > a`,
      func(element *colly.HTMLElement, ctx *Crawler.Context) {
         t := strings.Replace(element.Attr("href"), "https://twitter.com/", "", 1)
         ctx.TwitterId = strings.TrimSpace(t)
      })
}
