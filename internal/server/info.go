package server

import (
	"net/http"
)
func (s*Server) Info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("yes"))
}