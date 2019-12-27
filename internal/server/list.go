package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) List(w http.ResponseWriter, r *http.Request) {
	service := s.a.GetService()
	nodes, err := service.ListNodes()
	if err != nil {
		panic(err)
	}
	json.Marshal(nodes)
}
