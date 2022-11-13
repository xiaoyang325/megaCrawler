package stanford

import (
   "megaCrawler/Crawler"
)

func init() {
   w := Crawler.Register("stanford", "斯坦福大学",
      "https://stanford.edu/")

   w.SetStartingUrls([]string{
      "https://news.stanford.edu/section/science-technology/",
      "https://news.stanford.edu/section/social-sciences/",
      "https://news.stanford.edu/section/law-policy/a",
      "https://www.gsb.stanford.edu/insights",
      "https://med.stanford.edu/news.html",
      "https://ed.stanford.edu/news-media",
   })

   partNews(w)
   partGsb(w)
   partMed(w)
   partEd(w)
}