package handler

import "fmt"

type CreateShortUrlRequest struct {
	Url string `json:"url"`
}

func (request *CreateShortUrlRequest) Validate() error {

	if request.Url == "" {
		return fmt.Errorf("param: Url is required")
	}

	return nil
}
