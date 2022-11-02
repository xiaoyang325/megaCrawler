package megaCrawler

import (
	"encoding/json"
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
	PageType           PageType  `json:"page_type"`
	Id                 string    `json:"id"`
	Title              string    `json:"title"`
	Name               string    `json:"name"`
	SubTitle           string    `json:"sub_title"`
	Url                string    `json:"url"`
	Host               string    `json:"host"`
	Website            string    `json:"website"`
	CategoryText       string    `json:"category_text"`
	CategoryId         string    `json:"category_id"`
	Location           string    `json:"location"`
	CityISO            string    `json:"city_iso"`
	Language           string    `json:"language"`
	Authors            []string  `json:"authors"`
	PublicationTime    string    `json:"publication_time"`
	Description        string    `json:"description"`
	Content            string    `json:"content"`
	Image              []string  `json:"image"`
	Video              []string  `json:"video"`
	Audio              []string  `json:"audio"`
	File               []string  `json:"file"`
	Link               []string  `json:"link"`
	ViewCount          int       `json:"view_count"`
	LikeCount          int       `json:"like_count"`
	CommentCount       int       `json:"comment_count"`
	RepostCount        int       `json:"repost_count"`
	DislikeCount       int       `json:"dislike_count"`
	FavoriteCount      int       `json:"favorite_count"`
	Tags               []string  `json:"tags"`
	Keywords           []string  `json:"keywords"`
	Footnote           string    `json:"footnote"`
	Type               string    `json:"type"`
	LocationCityISO    string    `json:"location_city_iso"`
	NationalityCityISO string    `json:"nationality_city_iso"`
	Area               string    `json:"area"`
	Phone              string    `json:"phone"`
	Email              string    `json:"email"`
	Education          string    `json:"education"`
	TwitterId          string    `json:"twitter_id"`
	LinkedInId         string    `json:"linked_in_id"`
	FacebookId         string    `json:"facebook_id"`
	InstagramId        string    `json:"instagram_id"`
	WikipediaId        string    `json:"wikipedia_id"`
	ExpertWebsite      string    `json:"expert_website"`
	CrawlTime          time.Time `json:"crawl_time"`
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

func (c Context) process() (success bool) {
	var err error
	var marshal []byte
	now := time.Now()
	switch c.PageType {
	case Index:
		return true
	case News:
		n := news{
			Id:              c.Id,
			Title:           c.Title,
			SubTitle:        c.SubTitle,
			Url:             c.Url,
			Host:            c.Host,
			Website:         c.Website,
			CategoryText:    c.CategoryText,
			CategoryId:      c.CategoryId,
			Location:        c.Location,
			CityISO:         c.CityISO,
			Language:        c.Language,
			Authors:         c.Authors,
			PublicationTime: c.PublicationTime,
			Description:     c.Description,
			Content:         c.Content,
			Image:           c.Image,
			Video:           c.Video,
			Link:            c.Link,
			ViewCount:       c.ViewCount,
			LikeCount:       c.LikeCount,
			CommentCount:    c.CommentCount,
			RepostCount:     c.RepostCount,
			DislikeCount:    c.DislikeCount,
			FavoriteCount:   c.FavoriteCount,
			Tags:            c.Tags,
			Keywords:        c.Keywords,
			CrawlTime:       c.CrawlTime,
			CrawlTimestamp:  c.CrawlTime.Unix(),
			StoredTime:      now,
			StoredTimestamp: now.Unix(),
		}
		if n.Title == "" || n.Content == "" {
			return false
		}
		marshal, err = json.Marshal(n)
		if Debug {
			sugar.Debugw("Got News Type", spread(n)...)
		} else {
			newsChannel <- string(marshal)
		}
		return true
	case Report:
		n := report{
			Id:              c.Id,
			Title:           c.Title,
			SubTitle:        c.SubTitle,
			Url:             c.Url,
			Host:            c.Host,
			Website:         c.Website,
			CategoryText:    c.CategoryText,
			CategoryId:      c.CategoryId,
			CityISO:         c.CityISO,
			Language:        c.Language,
			Authors:         c.Authors,
			PublicationTime: c.PublicationTime,
			Description:     c.Description,
			Content:         c.Content,
			Image:           c.Image,
			Video:           c.Video,
			Audio:           c.Audio,
			File:            c.File,
			Link:            c.Link,
			ViewCount:       c.ViewCount,
			CommentCount:    c.CommentCount,
			Tags:            c.Tags,
			Keywords:        c.Keywords,
			CrawlTime:       c.CrawlTime,
			CrawlTimestamp:  c.CrawlTime.Unix(),
			StoredTime:      now,
			StoredTimestamp: now.Unix(),
		}
		if n.Title == "" || n.Content == "" {
			return false
		}
		marshal, err = json.Marshal(n)
		if Debug {
			sugar.Debugw("Got Report type", spread(n)...)
		} else {
			reportChannel <- string(marshal)
		}
		return true
	case Expert:
		image := ""
		if len(c.Image) > 0 {
			image = c.Image[0]
		}
		n := expert{
			Id:              c.Id,
			Title:           c.Title,
			Name:            c.Name,
			Url:             c.Url,
			Host:            c.Host,
			Website:         c.Website,
			CategoryText:    c.CategoryText,
			CategoryId:      c.CategoryId,
			Location:        c.Location,
			CityISO:         c.CityISO,
			Language:        c.Language,
			Description:     c.Description,
			Image:           image,
			Keywords:        c.Keywords,
			Type:            c.Type,
			Area:            c.Area,
			Phone:           c.Phone,
			Email:           c.Email,
			Education:       c.Education,
			TwitterId:       c.TwitterId,
			LinkedInId:      c.LinkedInId,
			FacebookId:      c.FacebookId,
			InstagramId:     c.InstagramId,
			WikipediaId:     c.WikipediaId,
			ExpertWebsite:   c.ExpertWebsite,
			CrawlTime:       c.CrawlTime,
			CrawlTimestamp:  c.CrawlTime.Unix(),
			StoredTime:      now,
			StoredTimestamp: now.Unix(),
		}
		if n.Name == "" {
			return false
		}
		marshal, err = json.Marshal(n)
		if Debug {
			sugar.Debugw("Got Expert type", spread(n)...)
		} else {
			expertChannel <- string(marshal)
		}
		return true
	}
	if err != nil {
		sugar.Error(err.Error())
	}
	return false
}
