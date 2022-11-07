package main

import (
	"ia04_projet1/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ballot", handlers.GetAllBallots).Methods(http.MethodGet)
	router.HandleFunc("/all_vote", handlers.GetAllVotes).Methods(http.MethodGet)
	router.HandleFunc("/new_ballot", handlers.CreateNewBallot).Methods(http.MethodPost)
	router.HandleFunc("/vote", handlers.CreateNewVote).Methods(http.MethodPost)
	router.HandleFunc("/result", handlers.CreateNewResult).Methods(http.MethodPost)
	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
