package megaCrawler

import "github.com/gocolly/colly/v2"

type Template struct {
	htmlHandlers     []CollyHTMLPair
	xmlHandlers      []XMLPair
	responseHandlers []func(response *colly.Response, ctx *Context)
}

func (t *Template) OnHTML(selector string, callback func(element *colly.HTMLElement, ctx *Context)) *Template {
	t.htmlHandlers = append(t.htmlHandlers, CollyHTMLPair{func(element *colly.HTMLElement) {
		callback(element, element.Request.Ctx.GetAny("ctx").(*Context))
	}, selector})
	return t
}

func (t *Template) OnXML(selector string, callback func(element *colly.XMLElement, ctx *Context)) *Template {
	t.xmlHandlers = append(t.xmlHandlers, XMLPair{callback, selector})
	return t
}

func (t *Template) OnResponse(callback func(response *colly.Response, ctx *Context)) *Template {
	t.responseHandlers = append(t.responseHandlers, callback)
	return t
}
