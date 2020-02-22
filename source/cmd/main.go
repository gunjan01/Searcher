package main

import (
	"io/ioutil"
	"logrus"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/gunjan01/searcher/source/search"
)

func indexMapping() string {
	filePath := "github.com/saltside/gunjan01/source/search/books.json"
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.WithError(err)
	}

	return string(raw)
}

func init() {
	client, err := search.NewES()
	if err != nil {
		err := client.EnsureIndex("literary_books", indexMapping())
		if err != nil {
			logrus.WithError(err)
		}
	}
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
		logrus.WithError(err).Fatalf("Failed to start the server")
	}
}
