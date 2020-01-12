package server

import (
	"net/http"
)

// Exec provides execution of commands
func (s *Server) Exec(w http.ResponseWriter, r *http.Request) {
	query, ok := r.URL.Query()["query"]
	if !ok || len(query[0]) < 1 {
		http.Error(w, "query attribute is missed", http.StatusBadRequest)
		return
	}
	command, ok := r.URL.Query()["command"]
	if !ok || len(command[0]) < 1 {
		http.Error(w, "command attribute is missed", http.StatusBadRequest)
		return
	}
	if err := s.a.Exec(query[0], command[0]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
