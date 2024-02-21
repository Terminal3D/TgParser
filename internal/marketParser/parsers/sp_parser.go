package parsers

import (
	"TgParser/internal/data"
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func ParseSP(resp *http.Response) (data.ProductData, error) {
	var parsedData data.ProductData

	var wg sync.WaitGroup

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()

	wg.Add(4)

	go func() {
		defer wg.Done()
		bodyReader := bytes.NewReader(bodyBytes)
		parsedData.Name = parseName(bodyReader)
	}()

	go func() {
		defer wg.Done()
		bodyReader := bytes.NewReader(bodyBytes)
		parsedData.Brand = parseBrand(bodyReader)
	}()

	go func() {
		defer wg.Done()
		bodyReader := bytes.NewReader(bodyBytes)
		parsedData.Price = parsePrice(bodyReader)
	}()

	go func() {
		defer wg.Done()
		bodyReader := bytes.NewReader(bodyBytes)
		parsedData.Sizes = parseSizes(bodyReader)
	}()

	wg.Wait()

	if parsedData.Available = parseAvailable(&parsedData); !parsedData.Available {
		return parsedData, fmt.Errorf("item is not available")
	}

	return parsedData, nil
}

func parseName(resp *bytes.Reader) string {
	tokenizer := html.NewTokenizer(resp)
	name, err := blockTextByKey(tokenizer, "span", "lblProductName", "id")
	if err != nil {
		log.Print(err)
		log.Println(" for name")
		return ""
	}
	return name
}

func parseBrand(resp *bytes.Reader) string {

	tokenizer := html.NewTokenizer(resp)
	brand, err := blockTextByKey(tokenizer, "span", "lblProductBrand", "id")
	if err != nil {
		log.Print(err)
		log.Println(" for brand")
		return ""
	}
	return brand
}

func parseAvailable(parsedData *data.ProductData) bool {

	if parsedData.Name == "" {
		return false
	}

	for _, size := range parsedData.Sizes {

		if size.Quantity != 0 {
			return true
		}
	}

	return false
}

func parsePrice(resp *bytes.Reader) float64 {
	tokenizer := html.NewTokenizer(resp)
	price, err := blockTextByKey(tokenizer, "span", "lblSellingPrice", "id")
	if err != nil {
		log.Println(err)
		return -1.0
	}
	price = strings.TrimSpace(price)
	price = strings.ReplaceAll(price, ",", ".")
	price = strings.Split(price, " ")[0]
	floatPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println(err)
		return 0.0
	}
	return floatPrice
}

func parseSizes(resp *bytes.Reader) []data.SizeData {
	tokenizer := html.NewTokenizer(resp)
	sizes, err := sizesListSP(tokenizer, "ul", "ulSizes")
	if err != nil {
		log.Println(err)
		return []data.SizeData{}
	}
	return sizes
}
