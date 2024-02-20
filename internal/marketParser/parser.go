package marketParser

import (
	"TgParser/internal/marketParser/models"
	"TgParser/internal/marketParser/parsers"
	"net/http"
	"net/url"
)

func RunParser(inputURL string) (*models.ProductData, error) {

	uri, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		return nil, err
	}

	sp, err := parsers.ParseSP(resp)
	if err != nil {
		return nil, err
	}

	return &sp, nil
}
