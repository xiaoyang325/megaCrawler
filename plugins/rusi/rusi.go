package rusi

import (
	"encoding/json"
	"github.com/gocolly/colly/v2"
	"megaCrawler/Crawler"
	"regexp"
	"strings"
)

type PageData struct {
	ComponentChunkName string `json:"componentChunkName"`
	Path               string `json:"path"`
	Result             struct {
		Data struct {
			Article struct {
				Typename string `json:"__typename"`
				Id       string `json:"id"`
				Path     struct {
					Alias string `json:"alias"`
				} `json:"path"`
				FieldReadTime           string      `json:"field_read_time"`
				Title                   string      `json:"title"`
				FieldMetatag            interface{} `json:"field_metatag"`
				Created                 string      `json:"created"`
				FieldOutsideAuthor      string      `json:"field_outside_author"`
				FieldProfileDescriptors string      `json:"field_profile_descriptors"`
				FieldPremium            bool        `json:"field_premium"`
				FieldAbstract           struct {
					Value     string `json:"value"`
					Format    string `json:"format"`
					Processed string `json:"processed"`
				} `json:"field_abstract"`
				Body struct {
					Summary   string `json:"summary"`
					Processed string `json:"processed"`
					Value     string `json:"value"`
				} `json:"body"`
				FieldPublicationIssue    interface{} `json:"field_publication_issue"`
				FieldPublicationVolume   interface{} `json:"field_publication_volume"`
				FieldTandfPublic         bool        `json:"field_tandf_public"`
				FieldTaylorAndFrancisUrl interface{} `json:"field_taylor_and_francis_url"`
				Relationships            struct {
					FieldMediaEnquiry interface{}   `json:"field_media_enquiry_"`
					FieldSections     []interface{} `json:"field_sections"`
					FieldContentType  struct {
						Name string `json:"name"`
						Path struct {
							Alias string `json:"alias"`
						} `json:"path"`
					} `json:"field_content_type"`
					FieldAuthor []struct {
						Id                string      `json:"id"`
						Title             string      `json:"title"`
						FieldFirstNames   string      `json:"field_first_names"`
						FieldEmailAddress interface{} `json:"field_email_address"`
						Path              struct {
							Alias string `json:"alias"`
						} `json:"path"`
						FieldPosition string `json:"field_position"`
						Relationships struct {
							FieldStaffDepartment []interface{} `json:"field_staff_department"`
							FieldUserPicture     interface{}   `json:"field_user_picture"`
						} `json:"relationships"`
					} `json:"field_author"`
					FieldPdf []struct {
						Id       string `json:"id"`
						Filename string `json:"filename"`
						Filesize int    `json:"filesize"`
						Filemime string `json:"filemime"`
						Fields   struct {
							CdnUrl string `json:"cdn_url"`
						} `json:"fields"`
					} `json:"field_pdf"`
					FieldRegion         []struct{ name string } `json:"field_region"`
					FieldResearchGroups []interface{}           `json:"field_research_groups"`
					FieldTopics         []struct {
						Id   string `json:"id"`
						Name string `json:"name"`
						Path struct {
							Alias string `json:"alias"`
						} `json:"path"`
					} `json:"field_topics"`
					FieldRelatedProject []struct {
						Id   string `json:"id"`
						Name string `json:"name"`
						Path struct {
							Alias string `json:"alias"`
						} `json:"path"`
					} `json:"field_related_project"`
					FieldMediaImage struct {
						Relationships struct {
							FieldMediaImage struct {
								ChildImageKitAsset struct {
									Fluid struct {
										AspectRatio float64 `json:"aspectRatio"`
										Base64      string  `json:"base64"`
										Sizes       string  `json:"sizes"`
										Src         string  `json:"src"`
										SrcSet      string  `json:"srcSet"`
									} `json:"fluid"`
								} `json:"childImageKitAsset"`
								Relationships struct {
									MediaImage []struct {
										FieldCredit     string      `json:"field_credit"`
										FieldCreditLink interface{} `json:"field_credit_link"`
									} `json:"media__image"`
								} `json:"relationships"`
							} `json:"field_media_image"`
						} `json:"relationships"`
					} `json:"field_media_image"`
					FieldSignpostImage    interface{}   `json:"field_signpost_image"`
					FieldJournalSignposts []interface{} `json:"field_journal_signposts"`
				} `json:"relationships"`
			} `json:"article"`
		} `json:"data"`
		PageContext struct {
			Id         string `json:"id"`
			Title      string `json:"title"`
			IsHomepage bool   `json:"isHomepage"`
			Breadcrumb struct {
				Location string `json:"location"`
				Crumbs   []struct {
					Pathname   string `json:"pathname"`
					CrumbLabel string `json:"crumbLabel"`
				} `json:"crumbs"`
			} `json:"breadcrumb"`
		} `json:"pageContext"`
	} `json:"result"`
	StaticQueryHashes []string `json:"staticQueryHashes"`
}
type NewsData struct {
	ComponentChunkName string `json:"componentChunkName"`
	Path               string `json:"path"`
	Result             struct {
		Data struct {
			Node struct {
				Title                      string `json:"title"`
				Id                         string `json:"id"`
				FieldPublishDate           string `json:"field_publish_date"`
				MachineDate                string `json:"machineDate"`
				FieldPrimaryTag            string `json:"field_primary_tag"`
				FieldFocus                 string `json:"field_focus"`
				FieldExternalPubDescriptor string `json:"field_external_pub_descriptor"`
				FieldOrganisation          struct {
					Title string `json:"title"`
					Uri   string `json:"uri"`
				} `json:"field_organisation"`
				Relationships struct {
					FieldSections    []interface{} `json:"field_sections"`
					FieldContentType struct {
						Path struct {
							Alias string `json:"alias"`
						} `json:"path"`
						Name string `json:"name"`
					} `json:"field_content_type"`
					FieldExternalPublication struct {
						Name string `json:"name"`
						Path struct {
							Alias string `json:"alias"`
						} `json:"path"`
						Relationships struct {
							FieldExternalPublicationLogo struct {
								Relationships struct {
									FieldMediaImage struct {
										ChildImageKitAsset struct {
											Fluid struct {
												AspectRatio int    `json:"aspectRatio"`
												Base64      string `json:"base64"`
												Sizes       string `json:"sizes"`
												Src         string `json:"src"`
												SrcSet      string `json:"srcSet"`
											} `json:"fluid"`
										} `json:"childImageKitAsset"`
									} `json:"field_media_image"`
								} `json:"relationships"`
							} `json:"field_external_publication_logo"`
						} `json:"relationships"`
					} `json:"field_external_publication"`
					FieldAuthor []struct {
						Id              string `json:"id"`
						Title           string `json:"title"`
						FieldFirstNames string `json:"field_first_names"`
						FieldPosition   string `json:"field_position"`
						Relationships   struct {
							FieldUserPicture struct {
								ChildImageKitAsset struct {
									Fluid struct {
										AspectRatio float64 `json:"aspectRatio"`
										Base64      string  `json:"base64"`
										Sizes       string  `json:"sizes"`
										Src         string  `json:"src"`
										SrcSet      string  `json:"srcSet"`
									} `json:"fluid"`
								} `json:"childImageKitAsset"`
							} `json:"field_user_picture"`
						} `json:"relationships"`
						Path struct {
							Alias string `json:"alias"`
						} `json:"path"`
					} `json:"field_author"`
					FieldRegion []struct {
						Id   string `json:"id"`
						Name string `json:"name"`
						Path struct {
							Alias string `json:"alias"`
						} `json:"path"`
					} `json:"field_region"`
					FieldResearchGroups []interface{} `json:"field_research_groups"`
					FieldTopics         []struct {
						Id   string `json:"id"`
						Name string `json:"name"`
						Path struct {
							Alias string `json:"alias"`
						} `json:"path"`
					} `json:"field_topics"`
					FieldRelatedProject []interface{} `json:"field_related_project"`
				} `json:"relationships"`
			} `json:"node"`
		} `json:"data"`
		PageContext struct {
			Id         string `json:"id"`
			Title      string `json:"title"`
			IsHomepage bool   `json:"isHomepage"`
			Breadcrumb struct {
				Location string `json:"location"`
				Crumbs   []struct {
					Pathname   string `json:"pathname"`
					CrumbLabel string `json:"crumbLabel"`
				} `json:"crumbs"`
			} `json:"breadcrumb"`
		} `json:"pageContext"`
	} `json:"result"`
	StaticQueryHashes []string `json:"staticQueryHashes"`
}

var PageTypeMap = map[string]Crawler.PageType{
	"sitemap":              Crawler.Index,
	"explore-our-research": Crawler.Report,
	"people":               Crawler.Expert,
	"news-and-comment":     Crawler.News,
	"in-the-news":          Crawler.News,
	"podcasts":             Crawler.News,
	"publication":          Crawler.Report,
}

func init() {
	w := Crawler.Register("rusi", "皇家联合服务研究所", "https://rusi.org/")
	w.SetStartingUrls([]string{"https://www.rusi.org/sitemap/sitemap-index.xml"})

	w.OnXML("//loc", func(element *colly.XMLElement, ctx *Crawler.Context) {
		reg := regexp.MustCompile("rusi.org/([\\w-]+)/")
		if matches := reg.FindStringSubmatch(element.Text); len(matches) == 2 {
			pageType, ok := PageTypeMap[matches[1]]
			if !ok {
				return
			}
			switch pageType {
			case Crawler.Index:
				w.Visit(element.Text, pageType)
			case Crawler.Expert:
				w.Visit(element.Text, pageType)
			default:
				url := strings.ReplaceAll(element.Text, "https://www.rusi.org/", "https://www.rusi.org/page-data/") + "/page-data.json"
				w.Visit(url, pageType)
			}
		}
	})
	//获取人物姓名
	w.OnHTML("[class^=\"ProfileTitleBlock-module--text\"] > h1", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Name = element.Text
	})

	//获取人物头衔
	w.OnHTML("[class^=\"ProfileTitleBlock-module--text\"] > samll", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Title = element.Text
	})

	//获取人物领域
	w.OnHTML("aside > ul > li > a > span", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Area = element.Text
	})

	//获取人物描述
	w.OnHTML("[class^=\"Section-module--content\"] > div > div > p", func(element *colly.HTMLElement, ctx *Crawler.Context) {
		ctx.Description = element.Text
	})

	w.OnResponse(func(response *colly.Response, ctx *Crawler.Context) {
		if strings.Contains(ctx.Url, "page-data.json") {
			var obj PageData
			_ = json.Unmarshal(response.Body, &obj)
			if obj.Result.Data.Article.Title != "" {
				art := obj.Result.Data.Article
				ctx.Title = art.Title
				ctx.Content = Crawler.HTML2Text(art.Body.Value)
				if ctx.Content == "" {
					ctx.Content = Crawler.HTML2Text(art.FieldAbstract.Value)
				}
				ctx.PublicationTime = art.Created
				for _, s := range obj.Result.Data.Article.Relationships.FieldAuthor {
					ctx.Authors = append(ctx.Authors, s.FieldFirstNames+" "+s.Title)
				}
				ctx.Area = art.Relationships.FieldRegion[0].name
				for _, topic := range art.Relationships.FieldTopics {
					ctx.Tags = append(ctx.Tags, topic.Name)
				}
				for _, pdf := range art.Relationships.FieldPdf {
					ctx.File = append(ctx.File, pdf.Fields.CdnUrl)
				}
				return
			}

			var obj2 NewsData
			_ = json.Unmarshal(response.Body, &obj)
			if obj2.Result.Data.Node.Title != "" {
				art := obj2.Result.Data.Node
				ctx.Title = art.Title
				ctx.Content = Crawler.HTML2Text(obj2.Result.Data.Node.FieldFocus)
				ctx.PublicationTime = art.FieldPublishDate
				for _, s := range obj.Result.Data.Article.Relationships.FieldAuthor {
					ctx.Authors = append(ctx.Authors, s.FieldFirstNames+" "+s.Title)
				}
				ctx.Area = art.Relationships.FieldRegion[0].Name
				for _, topic := range art.Relationships.FieldTopics {
					ctx.Tags = append(ctx.Tags, topic.Name)
				}
				return
			}
		}
	})
}
