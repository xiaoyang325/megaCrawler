package crawlers

import (
	"encoding/json"
	"strings"
	"time"

	"megaCrawler/crawlers/tester"
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
	ID                 string     `json:"id"`
	Title              string     `json:"title"`
	Name               string     `json:"name"`
	SubTitle           string     `json:"sub_title"`
	URL                string     `json:"url"`
	Host               string     `json:"host"`
	Website            string     `json:"website"`
	CategoryText       string     `json:"category_text"`
	CategoryID         string     `json:"category_id"`
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
	TwitterID          string     `json:"twitter_id"`
	LinkedInID         string     `json:"linked_in_id"`
	FacebookID         string     `json:"facebook_id"`
	InstagramID        string     `json:"instagram_id"`
	WikipediaID        string     `json:"wikipedia_id"`
	ExpertWebsite      string     `json:"expert_website"`
	CrawlTime          time.Time  `json:"crawl_time"`
	SubContext         []*Context `json:"subContext"`
}

type news struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	SubTitle        string    `json:"sub_title"`
	URL             string    `json:"url"`
	Host            string    `json:"host"`
	Website         string    `json:"website"`
	CategoryText    string    `json:"category_text"`
	CategoryID      string    `json:"category_id"`
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
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	SubTitle        string    `json:"sub_title"`
	URL             string    `json:"url"`
	Host            string    `json:"host"`
	Website         string    `json:"website"`
	CategoryText    string    `json:"category_text"`
	CategoryID      string    `json:"category_id"`
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
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Name            string    `json:"name"`
	URL             string    `json:"url"`
	Host            string    `json:"host"`
	Website         string    `json:"website"`
	CategoryText    string    `json:"category_text"`
	CategoryID      string    `json:"category_id"`
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
	TwitterID       string    `json:"twitter_user_id"`
	LinkedInID      string    `json:"linkedin_user_id"`
	FacebookID      string    `json:"facebook_user_id"`
	InstagramID     string    `json:"instagram_user_id"`
	WikipediaID     string    `json:"wikipedia_url"`
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
		URL:        ctx.URL,
		Host:       ctx.Host,
		Website:    ctx.ID,
		CrawlTime:  time.Time{},
	}
	ctx.SubContext = append(ctx.SubContext, k)
	return k
}

func (ctx *Context) process(tester *tester.Tester) (success bool) {
	var err error
	var marshal []byte
	now := time.Now()
	success = true

	if tester != nil && tester.Report.Count+tester.News.Count+tester.Expert.Count > 100 && !tester.Done {
		tester.Done = true
		Sugar.Info("tester finished, limit reached")
		tester.WG.Done()
		return false
	}

	for _, context := range ctx.SubContext {
		context.process(tester)
	}

	switch ctx.PageType {
	case Index:
		if tester != nil {
			tester.Index.Add(1)
			tester.Index.AddFilled(1)
		}
		return
	case News:
		n := news{
			ID:              ctx.ID,
			Title:           strings.TrimSpace(ctx.Title),
			SubTitle:        strings.TrimSpace(ctx.SubTitle),
			URL:             ctx.URL,
			Host:            ctx.Host,
			Website:         ctx.Website,
			CategoryText:    strings.TrimSpace(ctx.CategoryText),
			CategoryID:      ctx.CategoryID,
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
		if tester != nil {
			tester.News.Add(1)
		}
		if n.Title == "" || n.Content == "" {
			return false
		}
		if tester != nil {
			tester.News.AddFilled(1)
		}
		marshal, err = json.Marshal(n)
		if err != nil {
			Sugar.Error(err)
			return
		}
		if !Kafka {
			Sugar.Debugw("Got News Type", spread(n)...)
		} else {
			newsChannel <- string(marshal)
		}
		return
	case Report:
		n := report{
			ID:              ctx.ID,
			Title:           strings.TrimSpace(ctx.Title),
			SubTitle:        strings.TrimSpace(ctx.SubTitle),
			URL:             ctx.URL,
			Host:            ctx.Host,
			Website:         ctx.Website,
			CategoryText:    ctx.CategoryText,
			CategoryID:      ctx.CategoryID,
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
		if tester != nil {
			tester.Report.Add(1)
		}
		if n.Title == "" || (n.Content == "" && len(n.File) == 0) {
			return false
		}
		if tester != nil {
			tester.Report.AddFilled(1)
		}
		marshal, err = json.Marshal(n)
		if err != nil {
			Sugar.Error(err)
			return
		}
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
			ID:              ctx.ID,
			Title:           strings.TrimSpace(ctx.Title),
			Name:            strings.TrimSpace(ctx.Name),
			URL:             ctx.URL,
			Host:            ctx.Host,
			Website:         ctx.Website,
			CategoryText:    strings.TrimSpace(ctx.CategoryText),
			CategoryID:      ctx.CategoryID,
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
			TwitterID:       ctx.TwitterID,
			LinkedInID:      ctx.LinkedInID,
			FacebookID:      ctx.FacebookID,
			InstagramID:     ctx.InstagramID,
			WikipediaID:     ctx.WikipediaID,
			ExpertWebsite:   ctx.ExpertWebsite,
			CrawlTime:       ctx.CrawlTime,
			CrawlTimestamp:  ctx.CrawlTime.Unix(),
			StoredTime:      now,
			StoredTimestamp: now.Unix(),
		}
		if tester != nil {
			tester.Expert.Add(1)
		}
		if n.Name == "" {
			return false
		}
		if tester != nil {
			tester.Expert.AddFilled(1)
		}
		marshal, err = json.Marshal(n)
		if err != nil {
			Sugar.Error(err)
			return
		}
		if !Kafka {
			Sugar.Debugw("Got Expert type", spread(n)...)
		} else {
			expertChannel <- string(marshal)
		}
		return
	}
	return false
}
