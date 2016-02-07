package angelgo

type User struct {
	Name          string
	Bio           string
	BlogUrl       string `json:"blog_url"`
	WhatIDo       string `json:"what_i_do"`
	WhatIveBuilt  string `json:"what_ive_built"`
	OnlineBioUrl  string `json:"online_bio_url"`
	GithubUrl     string `json:"github_url"`
	ResumeUrl     string `json:"resume_url"`
	TwitterUrl    string `json:"twitter_url"`
	FacebookUrl   string `json:"facebook_url"`
	LinkedinUrl   string `json:"linkedin_url"`
	AngelListUrl  string `json:"angellist_url"`
	AboutMeUrl    string `json:"aboutme_url"`
	BehanceUrl    string `json:"behance_url"`
	DribbbleUrl   string `json:"dribbble_url"`
	Criteria      string
	Image         string
	Id            int64
	FollowerCount int64 `json:"follower_count"`
	Investor      bool
	Locations     []Location
	Roles         []Role
	Skills        []Skill
}

type CommonType struct {
	Id           int64
	TagType      string `json:"tag_type"`
	Name         string
	DisplayName  string `json:"display_name"`
	AngelListUrl string `json:"angelist_url"`
}

type Location CommonType

type Role CommonType

type Skill struct {
	CommonType
	Level string
}

type StartupRole struct {
    
}

type Startup struct {
    
}
