package extractors

import (
	"megaCrawler/crawlers"

	"github.com/gocolly/colly/v2"
)

func getMetaImgURL(dom *colly.HTMLElement) string {
	topMetaImage := GetMetaContent(dom, "meta[property=\"og:image\"]")
	if topMetaImage == "" {
		topMetaImage = dom.ChildAttr("link[rel=\"img_src|image_src\"]", "href")
	}
	if topMetaImage == "" {
		topMetaImage = GetMetaContent(dom, "meta[name=\"og:image\"]")
	}
	if topMetaImage == "" {
		topMetaImage = dom.ChildAttr("link[rel=\"icon\"]", "href")
	}
	if url, err := dom.Request.URL.Parse(topMetaImage); err != nil {
		return url.String()
	} else {
		return topMetaImage
	}
}

func getImgURLs(dom *colly.HTMLElement) (images []string) {
	return dom.ChildAttrs("img", "src")
}

func Image(ctx *crawlers.Context, dom *colly.HTMLElement) {
	ctx.Image = append([]string{getMetaImgURL(dom)}, getImgURLs(dom)...)
}
