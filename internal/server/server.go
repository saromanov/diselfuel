package server

import (
	"log"
	"net/http"
	"os"

	lucio "github.com/arriqaaq/server"
	"github.com/go-chi/chi"
	"github.com/saromanov/diselfuel/internal/config"
)

// New provides start of the server
func New(c *config.Config) {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("yes"))
	})
	server := lucio.NewServer(r, "0.0.0.0", 8080)
	err := server.Serve()
	log.Println("terminated", os.Getpid(), err)
}
