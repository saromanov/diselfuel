package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/saromanov/diselfuel/internal/config"
)

// New provides start of the server
func New(c *config.Config) {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yes"))
	})
	http.ListenAndServe(":3000", r)
}
