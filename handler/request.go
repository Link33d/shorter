package handler

import (
	"fmt"
	"net/url"
	"strings"
)

type CreateShortUrlRequest struct {
	Url string `json:"url"`
}

func (request *CreateShortUrlRequest) Validate() error {

	// Trim spaces around the url
	request.Url = strings.TrimSpace(request.Url)

	// Check if url is empty
	if request.Url == "" {
		return fmt.Errorf("param: Url is required")
	}

	// Format and validate the URL
	if err := formatURL(&request.Url); err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	return nil
}

func formatURL(input *string) error {

	// Add "http://" if necessary
	if !strings.HasPrefix(*input, "http://") && !strings.HasPrefix(*input, "https://") {
		*input = "http://" + *input
	}

	// Parse string as URL
	parsedURL, err := url.Parse(*input)
	if err != nil {
		return fmt.Errorf("malformed URL: %w", err)
	}

	// Check if the parsed URL has a valid host
	if parsedURL.Host == "" {
		return fmt.Errorf("URL missing host")
	}

	// Normalize and format the path and query
	//parsedURL.Path = url.PathEscape(parsedURL.Path)
	//parsedURL.RawQuery = parsedURL.Query().Encode()

	// Change url value
	*input = parsedURL.String()

	return nil
}
