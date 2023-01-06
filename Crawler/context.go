package Crawler

import (
	"encoding/json"
	"strings"
	"time"
)

type PageType int8

const (
	Index PageType = iota
	News
	Expert
	Report
)

type Context struct {
	PageType           PageType   `json:"page_type"`
	Id                 string     `json:"id"`
	Title              string     `json:"title"`
	Name               string     `json:"name"`
	SubTitle           string     `json:"sub_title"`
	Url                string     `json:"url"`
	Host               string     `json:"host"`
	Website            string     `json:"website"`
	CategoryText       string     `json:"category_text"`
	CategoryId         string     `json:"category_id"`
	Location           string     `json:"location"`
	CityISO            string     `json:"city_iso"`
	Language           string     `json:"language"`
	Authors            []string   `json:"authors"`
	PublicationTime    string     `json:"publication_time"`
	Description        string     `json:"description"`
	Content            string     `json:"content"`
	Image              []string   `json:"image"`
	Video              []string   `json:"video"`
	Audio              []string   `json:"audio"`
	File               []string   `json:"file"`
	Link               []string   `json:"link"`
	ViewCount          int        `json:"view_count"`
	LikeCount          int        `json:"like_count"`
	CommentCount       int        `json:"comment_count"`
	RepostCount        int        `json:"repost_count"`
	DislikeCount       int        `json:"dislike_count"`
	FavoriteCount      int        `json:"favorite_count"`
	Tags               []string   `json:"tags"`
	Keywords           []string   `json:"keywords"`
	Footnote           string     `json:"footnote"`
	Type               string     `json:"type"`
	LocationCityISO    string     `json:"location_city_iso"`
	NationalityCityISO string     `json:"nationality_city_iso"`
	Area               string     `json:"area"`
	Phone              string     `json:"phone"`
	Email              string     `json:"email"`
	Education          string     `json:"education"`
	TwitterId          string     `json:"twitter_id"`
	LinkedInId         string     `json:"linked_in_id"`
	FacebookId         string     `json:"facebook_id"`
	InstagramId        string     `json:"instagram_id"`
	WikipediaId        string     `json:"wikipedia_id"`
	ExpertWebsite      string     `json:"expert_website"`
	CrawlTime          time.Time  `json:"crawl_time"`
	SubContext         []*Context `json:"subContext"`
}

type news struct {
	Id              string    `json:"id"`
	Title           string    `json:"title"`
	SubTitle        string    `json:"sub_title"`
	Url             string    `json:"url"`
	Host            string    `json:"host"`
	Website         string    `json:"website"`
	CategoryText    string    `json:"category_text"`
	CategoryId      string    `json:"category_id"`
	Location        string    `json:"location"`
	CityISO         string    `json:"city_iso_cd"`
	Language        string    `json:"language"`
	Authors         []string  `json:"author_name"`
	PublicationTime string    `json:"publication_time"`
	Description     string    `json:"description"`
	Content         string    `json:"content"`
	Image           []string  `json:"image"`
	Video           []string  `json:"video"`
	Link            []string  `json:"link"`
	ViewCount       int       `json:"view_count"`
	LikeCount       int       `json:"like_count"`
	CommentCount    int       `json:"comment_count"`
	RepostCount     int       `json:"repost_count"`
	DislikeCount    int       `json:"dislike_count"`
	FavoriteCount   int       `json:"favorite_count"`
	Tags            []string  `json:"tags"`
	Keywords        []string  `json:"keywords"`
	CrawlTime       time.Time `json:"crawl_time"`
	CrawlTimestamp  int64     `json:"crawl_timestamp"`
	StoredTime      time.Time `json:"stored_time"`
	StoredTimestamp int64     `json:"stored_timestamp"`
}
type report struct {
	Id              string    `json:"id"`
	Title           string    `json:"title"`
	SubTitle        string    `json:"sub_title"`
	Url             string    `json:"url"`
	Host            string    `json:"host"`
	Website         string    `json:"website"`
	CategoryText    string    `json:"category_text"`
	CategoryId      string    `json:"category_id"`
	CityISO         string    `json:"city_iso_cd"`
	Language        string    `json:"language"`
	Authors         []string  `json:"author_name"`
	PublicationTime string    `json:"publication_time"`
	Description     string    `json:"description"`
	Content         string    `json:"content"`
	Image           []string  `json:"image"`
	Video           []string  `json:"video"`
	Audio           []string  `json:"audio"`
	File            []string  `json:"file"`
	Link            []string  `json:"link"`
	ViewCount       int       `json:"view_count"`
	CommentCount    int       `json:"comment_count"`
	Tags            []string  `json:"tags"`
	Keywords        []string  `json:"keywords"`
	CrawlTime       time.Time `json:"crawl_time"`
	CrawlTimestamp  int64     `json:"crawl_timestamp"`
	StoredTime      time.Time `json:"stored_time"`
	StoredTimestamp int64     `json:"stored_timestamp"`
}
type expert struct {
	Id              string    `json:"id"`
	Title           string    `json:"title"`
	Name            string    `json:"name"`
	Url             string    `json:"url"`
	Host            string    `json:"host"`
	Website         string    `json:"website"`
	CategoryText    string    `json:"category_text"`
	CategoryId      string    `json:"category_id"`
	Location        string    `json:"location"`
	CityISO         string    `json:"city_iso"`
	Language        string    `json:"language"`
	Description     string    `json:"description"`
	Image           string    `json:"image"`
	Keywords        []string  `json:"keywords"`
	Type            string    `json:"type"`
	Area            string    `json:"area"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	Education       string    `json:"education"`
	TwitterId       string    `json:"twitter_user_id"`
	LinkedInId      string    `json:"linkedin_user_id"`
	FacebookId      string    `json:"facebook_user_id"`
	InstagramId     string    `json:"instagram_user_id"`
	WikipediaId     string    `json:"wikipedia_url"`
	ExpertWebsite   string    `json:"expert_website"`
	CrawlTime       time.Time `json:"crawl_time"`
	CrawlTimestamp  int64     `json:"crawl_timestamp"`
	StoredTime      time.Time `json:"stored_time"`
	StoredTimestamp int64     `json:"stored_timestamp"`
}

func (ctx *Context) CreateSubContext() (k *Context) {
	k = &Context{
		PageType:   ctx.PageType,
		Authors:    []string{},
		Image:      []string{},
		Video:      []string{},
		Audio:      []string{},
		File:       []string{},
		Link:       []string{},
		Tags:       []string{},
		Keywords:   []string{},
		SubContext: []*Context{},
		Url:        ctx.Url,
		Host:       ctx.Host,
		Website:    ctx.Id,
		CrawlTime:  time.Time{},
	}
	ctx.SubContext = append(ctx.SubContext, k)
	return k
}

func (ctx *Context) process() (success bool) {
	var err error
	var marshal []byte
	now := time.Now()
	success = true

	if Test.Report.Count+Test.News.Count+Test.Expert.Count > 100 && !Test.Done {
		Test.Done = true
		Sugar.Info("Test limit reached")
		Test.WG.Done()
		return false
	}

	for _, context := range ctx.SubContext {
		go context.process()
	}

	switch ctx.PageType {
	case Index:
		if Test != nil {
			Test.Index.Add(1)
		}
		return
	case News:
		n := news{
			Id:              ctx.Id,
			Title:           strings.TrimSpace(ctx.Title),
			SubTitle:        strings.TrimSpace(ctx.SubTitle),
			Url:             ctx.Url,
			Host:            ctx.Host,
			Website:         ctx.Website,
			CategoryText:    strings.TrimSpace(ctx.CategoryText),
			CategoryId:      ctx.CategoryId,
			Location:        ctx.Location,
			CityISO:         ctx.CityISO,
			Language:        ctx.Language,
			Authors:         Unique(ctx.Authors),
			PublicationTime: strings.TrimSpace(ctx.PublicationTime),
			Description:     strings.TrimSpace(ctx.Description),
			Content:         strings.TrimSpace(ctx.Content),
			Image:           ctx.Image,
			Video:           ctx.Video,
			Link:            ctx.Link,
			ViewCount:       ctx.ViewCount,
			LikeCount:       ctx.LikeCount,
			CommentCount:    ctx.CommentCount,
			RepostCount:     ctx.RepostCount,
			DislikeCount:    ctx.DislikeCount,
			FavoriteCount:   ctx.FavoriteCount,
			Tags:            ctx.Tags,
			Keywords:        ctx.Keywords,
			CrawlTime:       ctx.CrawlTime,
			CrawlTimestamp:  ctx.CrawlTime.Unix(),
			StoredTime:      now,
			StoredTimestamp: now.Unix(),
		}
		Test.News.Add(1)
		if n.Title == "" || n.Content == "" {
			return false
		}
		Test.News.AddFilled(1)
		marshal, err = json.Marshal(n)
		if !Kafka {
			Sugar.Debugw("Got News Type", spread(n)...)
		} else {
			newsChannel <- string(marshal)
		}
		return
	case Report:
		n := report{
			Id:              ctx.Id,
			Title:           strings.TrimSpace(ctx.Title),
			SubTitle:        strings.TrimSpace(ctx.SubTitle),
			Url:             ctx.Url,
			Host:            ctx.Host,
			Website:         ctx.Website,
			CategoryText:    ctx.CategoryText,
			CategoryId:      ctx.CategoryId,
			CityISO:         ctx.CityISO,
			Language:        ctx.Language,
			Authors:         Unique(ctx.Authors),
			PublicationTime: strings.TrimSpace(ctx.PublicationTime),
			Description:     strings.TrimSpace(ctx.Description),
			Content:         strings.TrimSpace(ctx.Content),
			Image:           ctx.Image,
			Video:           ctx.Video,
			Audio:           ctx.Audio,
			File:            ctx.File,
			Link:            ctx.Link,
			ViewCount:       ctx.ViewCount,
			CommentCount:    ctx.CommentCount,
			Tags:            ctx.Tags,
			Keywords:        ctx.Keywords,
			CrawlTime:       ctx.CrawlTime,
			CrawlTimestamp:  ctx.CrawlTime.Unix(),
			StoredTime:      now,
			StoredTimestamp: now.Unix(),
		}
		Test.Report.Add(1)
		if n.Title == "" || (n.Content == "" && len(n.File) == 0) {
			return false
		}
		Test.Report.AddFilled(1)
		marshal, err = json.Marshal(n)
		if !Kafka {
			Sugar.Debugw("Got Report type", spread(n)...)
		} else {
			reportChannel <- string(marshal)
		}
		return
	case Expert:
		image := ""
		if len(ctx.Image) > 0 {
			image = ctx.Image[0]
		}
		n := expert{
			Id:              ctx.Id,
			Title:           strings.TrimSpace(ctx.Title),
			Name:            strings.TrimSpace(ctx.Name),
			Url:             ctx.Url,
			Host:            ctx.Host,
			Website:         ctx.Website,
			CategoryText:    strings.TrimSpace(ctx.CategoryText),
			CategoryId:      ctx.CategoryId,
			Location:        ctx.Location,
			CityISO:         ctx.CityISO,
			Language:        ctx.Language,
			Description:     strings.TrimSpace(ctx.Description),
			Image:           image,
			Keywords:        ctx.Keywords,
			Type:            ctx.Type,
			Area:            ctx.Area,
			Phone:           ctx.Phone,
			Email:           ctx.Email,
			Education:       ctx.Education,
			TwitterId:       ctx.TwitterId,
			LinkedInId:      ctx.LinkedInId,
			FacebookId:      ctx.FacebookId,
			InstagramId:     ctx.InstagramId,
			WikipediaId:     ctx.WikipediaId,
			ExpertWebsite:   ctx.ExpertWebsite,
			CrawlTime:       ctx.CrawlTime,
			CrawlTimestamp:  ctx.CrawlTime.Unix(),
			StoredTime:      now,
			StoredTimestamp: now.Unix(),
		}
		Test.Expert.Add(1)
		if n.Name == "" {
			return false
		}
		Test.Expert.AddFilled(1)
		marshal, err = json.Marshal(n)
		if !Kafka {
			Sugar.Debugw("Got Expert type", spread(n)...)
		} else {
			expertChannel <- string(marshal)
		}
		return
	}
	if err != nil {
		Sugar.Error(err.Error())
	}
	return false
}
