package server

import (
	_ "embed"
	"net/http"
)

//go:embed index.html
var html string

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	}
}
