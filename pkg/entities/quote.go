package entities

type Qoute struct {
	ID        int    `json:"id"`
	English   string `json:"english"`
	Indo      string `json:"indo"`
	Character string `json:"character"`
	Anime     string `json:"anime"`
}
