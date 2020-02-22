package main

import (
	"io/ioutil"
	"logrus"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/gunjan01/searcher/source/config"
	"github.com/gunjan01/searcher/source/search"
)

func indexMapping() string {
	raw, err := ioutil.ReadFile(config.Filepath)
	if err != nil {
		logrus.WithError(err)
	}

	return string(raw)
}

func init() {
	client, err := search.NewES()
	if err != nil {
		err := client.EnsureIndex(config.IndexName, indexMapping())
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
	err := http.ListenAndServe(config.Port, router())

	if err != nil {
		logrus.WithError(err).Fatalf("Failed to start the server")
	}
}
