package parsers

import (
	"TgParser/internal/marketParser/data"
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

	wg.Add(5)

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
		parsedData.Available = parseAvailable(bodyReader)
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

	return parsedData, nil
}

func parseName(resp *bytes.Reader) string {
	tokenizer := html.NewTokenizer(resp)
	name, err := getBlockTextByID(tokenizer, "span", "lblProductName")
	if err != nil {
		log.Println(err)
		return ""
	}
	return name
}

func parseBrand(resp *bytes.Reader) string {

	tokenizer := html.NewTokenizer(resp)
	brand, err := getBlockTextByID(tokenizer, "span", "lblProductBrand")
	if err != nil {
		log.Println(err)
		return ""
	}
	return brand
}

func parseAvailable(resp *bytes.Reader) bool {
	/* TODO checking availability by verifying fields name and qty from sizes list */
	return true
}

func parsePrice(resp *bytes.Reader) float64 {
	tokenizer := html.NewTokenizer(resp)
	price, err := getBlockTextByID(tokenizer, "span", "lblSellingPrice")
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
	sizes, err := getSizesListSP(tokenizer, "ul", "ulSizes")
	if err != nil {
		log.Println(err)
		return []data.SizeData{}
	}
	fmt.Println(sizes)
	return sizes
}
