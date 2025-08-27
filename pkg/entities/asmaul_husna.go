package entities

type AsmaulHusna struct {
	Urutan int    `json:"urutan"`
	Latin  string `json:"latin"`
	Arab   string `json:"arab"`
	Arti   string `json:"arti"`
}

type AsmaulHusnaResponse struct {
	StatusCode int           `json:"statusCode"`
	Total      int           `json:"total"`
	Data       []AsmaulHusna `json:"data"`
}

type AsmaulHusnaByLatinOrUrutanResponse struct {
	StatusCode int         `json:"statusCode"`
	Total      int         `json:"total"`
	Data       AsmaulHusna `json:"data"`
}
