package server

import (
	"os"

	lucio "github.com/arriqaaq/server"
	"github.com/go-chi/chi"
	"github.com/saromanov/diselfuel/internal/app"
	"github.com/saromanov/diselfuel/internal/config"
	"github.com/sirupsen/logrus"
)

// Server provides definition
type Server struct {
	a *app.App
}

// New provides start of the server
func New(a *app.App, c *config.Config, log *logrus.Logger) {
	r := chi.NewRouter()
	s := Server{
		a: a,
	}
	r.Get("/v1/info", s.Info)
	r.Get("/v1/nodes", s.List)

	log.Info("starting of the server")
	server := lucio.NewServer(r, "0.0.0.0", 8080)
	err := server.Serve()
	log.Println("terminated", os.Getpid(), err)
}
