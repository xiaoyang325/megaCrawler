package megaCrawler

type PageType int8

const (
	Index PageType = iota
	News
	Expert
	Report
)

type Context struct {
	PageType           PageType
	Id                 string
	Title              string
	Name               string
	SubTitle           string
	Url                string
	Host               string
	Website            string
	CategoryText       string
	CategoryId         string
	Location           string
	CityISO            string
	Language           string
	Authors            []string
	PublicationTime    string
	Description        string
	Content            string
	Image              []string
	Video              []string
	Audio              []string
	File               []string
	Link               []string
	ViewCount          int
	LikeCount          int
	CommentCount       int
	RepostCount        int
	DislikeCount       int
	FavoriteCount      int
	Tags               []string
	Keywords           []string
	Type               string
	LocationCityISO    string
	NationalityCityISO string
	Area               string
	Phone              string
	Email              string
	Education          string
	TwitterId          string
	LinkedInId         string
	FacebookId         string
}
