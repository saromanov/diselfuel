package server

import (
	"encoding/json"
	"net/http"

	"github.com/saromanov/diselfuel/internal/models"
)

// List returns list of nodes
func (s *Server) List(w http.ResponseWriter, r *http.Request) {
	service := s.a.GetService()
	nodes, err := service.ListNodes(models.FilterNodes{})
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(nodes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
