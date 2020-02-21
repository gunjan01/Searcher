package main

import (
	"net/http"

	"github.com/go-zoo/bone"
)

func init() {
	client := NewES()
	client.EnsureIndex("literary_books", "books.json")

}

func router() http.Handler {
	mux := bone.New()
	handler := Handler{}

	mux.Get(
		"/ping",
		http.HandlerFunc(handler.PingHandler),
	)

	mux.Post(
		"/search",
		http.HandlerFunc(handler.Searcher),
	)

	return mux
}

func main() {
	err := http.ListenAndServe(":8080", router())

	if err != nil {
		// Log the error
	}
}
