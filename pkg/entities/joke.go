package entities

type JokeTextResponse struct {
	Status    int    `json:"status"`
	End_point string `json:"end_point"`
	Method    string `json:"method"`
	Data      string `json:"data"`
}

type JokeImageData struct {
	Url    string `json:"url"`
	Source string `json:"source"`
}

type JokeImageResponse struct {
	Status    int           `json:"status"`
	End_point string        `json:"end_point"`
	Method    string        `json:"method"`
	Data      JokeImageData `json:"data"`
}
