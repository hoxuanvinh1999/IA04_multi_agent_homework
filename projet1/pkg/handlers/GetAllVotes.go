package handlers

import (
	"encoding/json"
	"ia04_projet1/pkg/data"
	"net/http"
)

func GetAllVotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Votes)
}
