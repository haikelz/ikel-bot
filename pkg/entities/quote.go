package entities

type Qoute struct {
	ID        int    `json:"id"`
	English   string `json:"english"`
	Indo      string `json:"indo"`
	Character string `json:"character"`
	Anime     string `json:"anime"`
}

type QuoteResponse struct {
	Sukses bool    `json:"sukses"`
	Next   bool    `json:"next"`
	Result []Qoute `json:"result"`
}

type QuotesResponse struct {
	Sukses bool    `json:"sukses"`
	Result []Qoute `json:"result"`
}
