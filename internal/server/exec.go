package server

import (
	"net/http"
)

// Exec provides execution of commands
func (s *Server) Exec(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["query"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "query attribute is missed", http.StatusBadRequest)
		return
	}
	if err := s.a.Exec(keys[0]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
