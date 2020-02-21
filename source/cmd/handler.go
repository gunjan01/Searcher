package main

import (
	"io"
	"net/http"
)

// Handler implements the handler.
type Handler struct {
}

// PingHandler is a health check endpoint.
func (h *Handler) PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

// Searcher makes a call to elastic search and returns the relevant book.
func (h *Handler) Searcher(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
