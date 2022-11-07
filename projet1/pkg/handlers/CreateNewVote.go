package handlers

import (
	"encoding/json"
	"ia04_projet1/pkg/data"
	"ia04_projet1/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CreateNewVote(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var vote models.Vote
	json.Unmarshal(body, &vote)
	ballotid := vote.VoteID
	voterid := vote.AgentID

	for _, vote := range data.Votes {
		if voterid == vote.AgentID {
			json.NewEncoder(w).Encode("403: Vote already made")
			return
		}
	}

	var currentBallot models.Ballot
	var find_ballot = 0
	for _, ballot := range data.Ballots {
		if ballot.BallotID == ballotid {
			currentBallot = ballot
			find_ballot = 1
		}
	}
	if find_ballot == 0 {
		json.NewEncoder(w).Encode("Ballot not found")
		return
	}
	var voter_check = 0
	for _, id_voter := range currentBallot.VoterIDs {
		if voterid == id_voter {
			voter_check = 1
			break
		}
	}
	if voter_check == 0 {
		json.NewEncoder(w).Encode("Voter not in this ballot")
		return
	}

	deadline_time, _ := time.Parse("2006-01-02", currentBallot.Deadline)
	if deadline_time.Sub(time.Now()) > 0 {
		data.Votes = append(data.Votes, vote)
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Created New Vote")
	} else {
		json.NewEncoder(w).Encode("Deadline passed")
	}

}
