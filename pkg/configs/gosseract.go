package configs

import (
	"github.com/otiai10/gosseract/v2"
)

func NewGoserract() *gosseract.Client {
	client := gosseract.NewClient()
	defer client.Close()

	return client
}
