package server

import (
	"encoding/json"
	"fmt"
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
	response, err := s.a.Exec(query[0], command[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("RESPONSE: ", response)

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
