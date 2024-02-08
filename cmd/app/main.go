package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	data, err := http.Get("https://md.sportsdirect.com/under-armour-armour-hustle-50-backpack-710979#colcode=71097941")
	if err != nil {
		return
	}
	d, err := io.ReadAll(data.Body)
	fmt.Println(string(d))
}
