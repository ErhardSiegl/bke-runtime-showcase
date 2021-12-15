package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/hello", HelloHandler)

	http.ListenAndServe(":8080", mux)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Default().Print("RootHandler received request")
	_, err := fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	if err != nil {
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Default().Print("HelloHandler received request")
	keys, ok := r.URL.Query()["name"]
	var name string

	if !ok || len(keys[0]) < 1 {
		name = "Jane Doe"
	} else {
		name = keys[0]
	}

	fmt.Fprintf(w, "Hello, %s\n", name)

}
