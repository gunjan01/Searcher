package main

import (
	"encoding/json"
	"io"
	"logrus"
	"net/http"

	"github.com/gunjan01/searcher/source/search"
)

// Handler implements the handler.
type Handler struct {
	client *search.Es
}

// PingHandler is a health check endpoint.
func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"ping": pong}`)
}

// Searcher makes a call to elastic search and returns the relevant book.
func (h *Handler) Searcher(w http.ResponseWriter, r *http.Request) {
	req := search.GetLiteraryBooksRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logrus.WithError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	searchSource, err := h.client.SearchQuery(req)
	if err != nil {
		h.client.ExtractBooks(searchSource)
		// construct response
	}

	w.WriteHeader(http.StatusOK)
}
