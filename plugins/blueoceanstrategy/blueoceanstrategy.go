package blueoceanstrategy

import (
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"strings"
)

// 这个函数用于分隔使用 ",", "&" 和 "and" 的字符串
// 并返回分割开的 []string
func cutToList(input_str string) ([]string)  {
	name_str := strings.Replace(input_str, "and", ",", -1)
	name_str = strings.Replace(name_str, "&", ",", -1)
	name_list := strings.Split(name_str, ",")
	for index, value := range name_list {
		name_list[index] = strings.TrimSpace(value)
	}

	return name_list
}

func init() {
	w := Crawler.Register("blueoceanstrategy", "海洋战略研究所",
			"https://www.blueoceanstrategy.com/")
	
	w.SetStartingUrls([]string{
		"https://www.blueoceanstrategy.com/blue-ocean-strategy-examples",
		"https://www.blueoceanstrategy.com/teaching-materials/all-cases/",
	})

	// 访问 Report 从 Index
	w.OnHTML(`div > div > div > div.et_pb_section.et_pb_section_2.et_pb_with_background.et_section_regular > div > div > div > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 访问 Report 从 Index
	w.OnHTML(`div > div > div > div.et_pb_section.et_pb_section_4.et_section_regular > div > div.et_pb_css_mix_blend_mode_passthrough > div > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 访问 Report 从 Index
	w.OnHTML(`#post-260622 > div > div > div > div.et_pb_section.et_pb_section_4.et_section_regular > div.et_pb_row.et_pb_row_4.et_pb_row_5col > div > div > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			w.Visit(element.Attr("href"), Crawler.Report)
		})

	// 获取 Title
	w.OnHTML(`#main-content > div > div > div.et_pb_section.et_pb_section_1_tb_body.et_section_regular > div > div > div.et_pb_module.et_pb_text.et_pb_text_0_tb_body.et_pb_text_align_center.et_pb_bg_layout_dark > div > h1`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Title
	w.OnHTML(`div > div > div > div.et_pb_section.et_pb_section_0.et_section_regular > div > div.et_pb_column.et_pb_column_3_4.et_pb_column_1.et_pb_css_mix_blend_mode_passthrough.et-last-child > div.et_pb_module.et_pb_text.et_pb_text_0.et_pb_text_align_left.et_pb_bg_layout_light > div > h1`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 Title
	w.OnHTML(`div > div > div > div > div.et_pb_row.et_pb_row_3 > div > div.et_pb_module.et_pb_text.et_pb_text_6.et_pb_text_align_left.et_pb_bg_layout_light > div > div.casestudy-title > h2`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Title = strings.TrimSpace(element.Text)
		})

	// 获取 SubTitle
	w.OnHTML(`#main-content > div > div > div.et_pb_section.et_pb_section_1_tb_body.et_section_regular > div > div > div.et_pb_module.et_pb_text.et_pb_text_1_tb_body.et_pb_text_align_left.et_pb_bg_layout_light > div`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.SubTitle = strings.TrimSpace(element.Text)
		})

	// 获取 SubTitle
	w.OnHTML(`div > div > div > div > div.et_pb_row.et_pb_row_3 > div > div.et_pb_module.et_pb_text.et_pb_text_7.et_pb_text_align_left.et_pb_bg_layout_light > div > h3`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.SubTitle = strings.TrimSpace(element.Text)
		})

	// 获取 Authors
	w.OnHTML(`#main-content > div > div > div.et_pb_section.et_pb_section_1_tb_body.et_section_regular > div > div > div.et_pb_module.et_pb_blurb.et_pb_blurb_0_tb_body.et_pb_text_align_left.et_pb_blurb_position_left.et_pb_bg_layout_light > div > div.et_pb_blurb_container > h4 > span > a`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			if (strings.Contains(element.Text, "&")) {
				ctx.Authors = append(ctx.Authors, cutToList(element.Text)...)
			} else {
				ctx.Authors = append(ctx.Authors, strings.TrimSpace(element.Text))
			}
		})

	// 获取 Authors
	w.OnHTML(`div > div > div > div > div.et_pb_row.et_pb_row_3 > div > div.et_pb_module.et_pb_text.et_pb_text_7.et_pb_text_align_left.et_pb_bg_layout_light > div > p`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			raw_str := strings.Replace(element.Text, "Author(s):", "", 1)
			if (strings.Contains(raw_str, "&")) {
				ctx.Authors = append(ctx.Authors, cutToList(raw_str)...)
			} else {
				ctx.Authors = append(ctx.Authors, strings.TrimSpace(raw_str))
			}
		})

	// 获取 Content
	w.OnHTML(`#main-content > div > div > div.et_pb_section.et_pb_section_2_tb_body.et_section_regular > div > div.et_pb_column.et_pb_column_3_4.et_pb_column_2_tb_body.et_pb_css_mix_blend_mode_passthrough > div > div > div > div > div.et_pb_row.et_pb_row_0 > div`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})

	// 获取 Content
	w.OnHTML(`div > div > div > div.et_pb_section.et_pb_section_4.et_section_regular > div.et_pb_row.et_pb_row_4 > div > div > div`,
		func(element *colly.HTMLElement, ctx *Crawler.Context) {
			ctx.Content = strings.TrimSpace(element.Text)
		})
}
