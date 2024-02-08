package parsers

import (
	"TgParser/internal/marketParser/data"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func findTag(tokenizer *html.Tokenizer, blockName string, id string) error {

	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			return fmt.Errorf("id not found")

		case tt == html.StartTagToken:
			token := tokenizer.Token()

			if token.Data != blockName {
				continue
			}

			for _, attr := range token.Attr {
				if attr.Key == "id" && attr.Val == id {
					return nil
				}
			}
		}
	}
}

func getBlockTextByID(tokenizer *html.Tokenizer, blockName string, id string) (string, error) {
	err := findTag(tokenizer, blockName, id)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(getText(tokenizer)), nil
}

func getSizesListSP(tokenizer *html.Tokenizer, listType string, id string) ([]data.SizeData, error) {
	err := findTag(tokenizer, listType, id)
	if err != nil {
		return nil, err
	}

	var sizes []data.SizeData
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return nil, fmt.Errorf("end of HTML input")
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data == "li" {
				var item data.SizeData
				for _, attr := range token.Attr {
					if attr.Key == "data-text" {
						item.Size = attr.Val
					}

					if attr.Key == "data-stock-qty" {
						item.Quantity = attr.Val
					}
				}
				sizes = append(sizes, item)
			}
		case html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == listType {
				return sizes, nil
			}
		}
	}
}

func getText(tokenizer *html.Tokenizer) string {
	var text strings.Builder
	depth := 1
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken, html.EndTagToken:
			depth--
			if depth <= 0 {
				return text.String()
			}
		case html.StartTagToken, html.SelfClosingTagToken:
			depth++
		case html.TextToken:
			if depth == 1 {
				text.WriteString(tokenizer.Token().Data)
			}
		}
	}
}
