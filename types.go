package angelgo

type User struct {
	Name          string
	Bio           string
	BlogUrl       string `json:"blog_url"`
	OnlineBioUrl  string `json:"online_bio_url"`
	TwitterUrl    string `json:"twitter_url"`
	FacebookUrl   string `json:"facebook_url"`
	LinkedinUrl   string `json:"linkedin_url"`
    AngelListUrl  string `json:"angellist_url"`
	Image         string
    Id            int64
	FollowerCount int64  `json:"follower_count"`
	Investor      bool
	Locations     []Location
	Roles         []Role

}

type Location struct {
	Id           int64
	TagType      string `json:"tag_type"`
	Name         string
	DisplayName  string `json:"display_name"`
	AngelListUrl string `json:"angelist_url"`
}

type Role struct {
	Id           int64
	TagType      string `json:"tag_type"`
	Name         string
	DisplayName  string `json:"display_name"`
	AngelListUrl string `json:"angelist_url"`
}
