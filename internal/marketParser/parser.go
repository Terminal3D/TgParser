package marketParser

import (
	"TgParser/internal/marketParser/parsers"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func RunParser() {

	var inputURL string
	fmt.Scan(&inputURL)

	uri, err := url.ParseRequestURI(inputURL)
	if err != nil {
		log.Fatal("Invalid URL")
		return
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		log.Fatal("Failed GET-request")
		return
	}

	//t, _ := io.ReadAll(resp.Body)
	//
	//fmt.Println(string(t))

	sp, err := parsers.ParseSP(resp)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(sp)
}
