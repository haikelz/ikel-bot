package entities

type BackgroundResult struct {
	Success bool   `json:"success"`
	Base64  string `json:"base64"`
	Message string `json:"message"`
}

type RemoveBgRequest struct {
	ImageFileB64 string `json:"image_file_b64"`
	BgColor      string `json:"bg_color"`
}

type RemoveBgResponse struct {
	Data struct {
		Result_b64 string `json:"result_b64"`
	} `json:"data"`
}

// Alternative response structure for some APIs
type RemoveBgResponseAlt struct {
	Result  string `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
