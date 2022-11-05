package bihr

import (
   "github.com/gocolly/colly/v2"
   "megaCrawler/megaCrawler"
   "strings"
)

func init() {
   w := megaCrawler.Register("bihr", "", "https://www.bihr.org.uk/")

   w.SetStartingUrls([]string{
      "https://www.bihr.org.uk/the-human-rights-act-the-icescr",
      "https://www.bihr.org.uk/covid-19-vaccine-and-human-rights",
      "https://www.bihr.org.uk/the-mental-health-act-reform-and-human-rights",
      "https://www.bihr.org.uk/co-design-a-human-rights-support-solution-with-bihr",
      "https://www.bihr.org.uk/dnar-decision-making-2020-bihr",
   })

   // 添加 Title 到ctx
   w.OnHTML(".title",
      func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
         ctx.Title = element.Text
      })

   // 添加 Author 到ctx
   w.OnHTML("h2[class=\"memberName blogOwner\"]",
      func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
         ctx.Authors = append(ctx.Authors, element.Text)
      })

   // 添加 Content 到ctx
   w.OnHTML("div[class=\"content postContent pageContent\"]",
      func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
         ctx.Content = element.Text
      })

   // 通过图片获取 File 和 新的Report
   w.OnHTML("div[class=\"content postContent pageContent\"]>p>a",
      func(element *colly.HTMLElement, ctx *megaCrawler.Context) {
         url := element.Attr("href")
         if (strings.Contains(url, "Download")) {
            ctx.File = append(ctx.File, url)
         } else {
            w.Visit(url, megaCrawler.Report)
         }
      })
}