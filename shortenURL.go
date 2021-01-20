package main

import (
	"fmt"
	"github.com/kare/base62"
)

func main() {
	var urlVal int64 = 2222222
	encodeURL := base62.Encode(urlVal)
	fmt.Println(encodeURL)

	byteURL, _ := base62.Decode(encodeURL)
	fmt.Println(byteURL)
}