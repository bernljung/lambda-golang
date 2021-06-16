package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"

	handlers "github.com/bernljung/lambda-golang"
)

func callHandler(w http.ResponseWriter, r *http.Request) {
	res, err := handlers.HandlerFunc(context.Background(), strings.TrimPrefix(r.URL.Path, "/"))
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, res)
}

func main() {
	http.HandleFunc("/", callHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
