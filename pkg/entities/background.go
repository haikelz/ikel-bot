package entities

type BackgroundResult struct {
	Success bool   `json:"success"`
	Base64  string `json:"base64"`
	Message string `json:"message"`
}

type RemoveBgRequest struct {
	ImageUrl string `json:"image_url"`
	BgColor  string `json:"bg_color"`
}

type RemoveBgResponse struct {
	Data struct {
		Result string `json:"result"`
	} `json:"data"`
}
