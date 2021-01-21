package main

import (
	"fmt"
	"kkn.fi/base62"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
	"sync"
	"github.com/pkg/browser"
)

var mu sync.Mutex
var nb int64 = 1
var addr string = "http://localhost:8000/"
var URLs map[string]int64

func URLHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	URLs[vars["URL"]] = nb
	encodeURL := base62.Encode(nb)
	fmt.Fprintf(w, "URL : %v\n", encodeURL)
	fmt.Fprintf(w, "nb : %v\n", nb)
	mu.Lock()
	nb++
	mu.Unlock()
}

func URL2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	URLs[vars["URL"]] = nb
	encodeURL := base62.Decode(nb)
	fmt.Fprintf(w, "URL : %v\n", encodeURL)
	fmt.Fprintf(w, "nb : %v\n", nb)
}

func main() {
	URLs = make(map[string]int64)
	r := mux.NewRouter()
	r.HandleFunc("/{URL}", URL2)
	r.HandleFunc("/new/{URL}", URLHandler)
	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}