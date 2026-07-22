package model

type Link struct {
	Label string
	URL   string
}

type Profile struct {
	Name     string
	Title    string
	Company  string
	Location string
	Bio      string
	Photo    string
	Links    []Link
}

type Experience struct {
	Company     string
	Role        string
	Start       string
	End         string
	Highlights  []string
}

type SkillGroup struct {
	Category string
	Items    []string
}

type Education struct {
	Institution string
	Degree      string
	Year        string
	Image       string
	VerifyURL   string
}

type PageData struct {
	Profile    Profile
	Experience []Experience
	Skills     []SkillGroup
	Education  []Education
	Year       string
}
