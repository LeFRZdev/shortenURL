package main

import (
	"fmt"
	"kkn.fi/base62"
	"github.com/gorilla/mux"
)

func main() {
	var urlVal int64 = 2222222
	encodeURL := base62.Encode(urlVal)
	fmt.Println(encodeURL)

	byteURL, _ := base62.Decode(encodeURL)
	fmt.Println(byteURL)
}