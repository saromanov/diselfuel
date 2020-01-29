package server

import (
	"encoding/json"
	"net/http"

	"github.com/saromanov/diselfuel/internal/models"
)

// Apply provides execution of the list of commands
func (s *Server) Apply(w http.ResponseWriter, r *http.Request) {

	var a *models.Execution
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := s.a.Apply(a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
