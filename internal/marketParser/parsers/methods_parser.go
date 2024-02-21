package parsers

import (
	"TgParser/internal/data"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strconv"
	"strings"
)

const ItemSizeAttrSP = "data-text"
const ItemQuantityAttrSP = "data-stock-qty"

func findTag(tokenizer *html.Tokenizer, blockName string, id string) error {

	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			return fmt.Errorf("%s for %s block not found", id, blockName)

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

func blockTextByID(tokenizer *html.Tokenizer, blockName string, id string) (string, error) {
	err := findTag(tokenizer, blockName, id)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text(tokenizer)), nil
}

func sizesListSP(tokenizer *html.Tokenizer, listType string, id string) ([]data.SizeData, error) {
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
					if attr.Key == ItemSizeAttrSP {
						item.Size = attr.Val
					}

					if attr.Key == ItemQuantityAttrSP {
						item.Quantity, err = strconv.Atoi(attr.Val)
						if err != nil {
							log.Println("Error parsing quantity for size " + item.Size)
							item.Quantity = 0
						}
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

func text(tokenizer *html.Tokenizer) string {
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
