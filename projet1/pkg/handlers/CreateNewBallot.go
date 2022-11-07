package handlers

import (
	"encoding/json"
	"ia04_projet1/pkg/data"
	"ia04_projet1/pkg/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func createBallotID() string {
	var str string = "vote"
	var value = rand.Intn(100)
	result := str + strconv.Itoa(value)
	return result
}
func CreateNewBallot(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var ballot models.Ballot
	json.Unmarshal(body, &ballot)

	if ballot.Rule != "approval" && ballot.Rule != "borda" && ballot.Rule != "copeland" && ballot.Rule != "majoriry" && ballot.Rule != "condorcet" {
		json.NewEncoder(w).Encode("The Rule not Valid")
		return
	}
	_, er := time.Parse("Jan 2, 2006 at 3:04pm", ballot.Deadline)
	if er != nil {
		json.NewEncoder(w).Encode("The Deadline not Valid")
	}

	ballot.BallotID = createBallotID()
	var output_ballot models.OutputBallot
	output_ballot.BallotID = ballot.BallotID
	data.OutputBallots = append(data.OutputBallots, output_ballot)
	data.Ballots = append(data.Ballots, ballot)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created New Ballot")
	json.NewEncoder(w).Encode(output_ballot)
}
