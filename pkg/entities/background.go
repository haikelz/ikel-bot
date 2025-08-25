package entities

type BackgroundResult struct {
	Success bool   `json:"success"`
	Base64  string `json:"base64"`
	Message string `json:"message"`
}

type RemoveBGRequest struct {
	ImageFileB64 string `json:"image_file_64"`
	BgColor      string `json:"bg_color"`
}

type RemoveBGResponse struct {
	Data struct {
		ResultB64 string `json:"result_b64"`
	} `json:"data"`
}
