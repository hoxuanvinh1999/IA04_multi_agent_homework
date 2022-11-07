package handlers

import (
	"encoding/json"
	"fmt"
	"ia04_projet1/pkg/comsoc"
	"ia04_projet1/pkg/data"
	"ia04_projet1/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/exp/maps"
)

func CreateNewResult(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result models.Result
	json.Unmarshal(body, &result)

	ballotid := result.BallotID
	var currentBallot models.Ballot
	var find_ballot = 0
	var current_profile comsoc.Profile
	var current_thresholds []int
	var swf_result comsoc.Count
	var scf_result []comsoc.Alternative
	var winner int
	var ranking []comsoc.Alternative
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

	var list_voter_id = currentBallot.VoterIDs
	var found_voter = 0
	for _, voter_id := range list_voter_id {
		found_voter = 0
		for _, voter := range data.Votes {
			if voter_id == voter.AgentID {
				found_voter = 1
				current_profile = append(current_profile, voter.Prefs)
				current_thresholds = append(current_thresholds, int(voter.Options[0]))
			}
		}
		if found_voter == 0 {
			msg := fmt.Sprintf("%s not voted", voter_id)
			json.NewEncoder(w).Encode(msg)
			return
		}
	}
	deadline_time, _ := time.Parse("2006-01-02", currentBallot.Deadline)
	if deadline_time.Sub(time.Now()) < 0 {
		switch currentBallot.Rule {
		case "majority":
			swf_result, err = comsoc.MajoritySWF(current_profile)
			ranking = maps.Keys(swf_result)
			scf_result, err = comsoc.MajoritySCF(current_profile)
			if len(scf_result) == 1 {
				winner = int(scf_result[0])
			}
		case "approval":
			swf_result, err = comsoc.ApprovalSWF(current_profile, current_thresholds)
			ranking = maps.Keys(swf_result)
			scf_result, err = comsoc.ApprovalSCF(current_profile, current_thresholds)
			if len(scf_result) == 1 {
				winner = int(scf_result[0])
			}
		case "borda":
			swf_result, err = comsoc.BordaSWF(current_profile)
			ranking = maps.Keys(swf_result)
			scf_result, err = comsoc.BordaSCF(current_profile)
			if len(scf_result) == 1 {
				winner = int(scf_result[0])
			}
		case "copeland":
			swf_result, err = comsoc.CopelandSWF(current_profile)
			ranking = maps.Keys(swf_result)
			scf_result, err = comsoc.CopelandSCF(current_profile)
			if len(scf_result) == 1 {
				winner = int(scf_result[0])
			}
		case "condorcet":
			ranking, err = comsoc.CondorcetWinner(current_profile)
			json.NewEncoder(w).Encode(err)
			if len(ranking) == 1 {
				winner = int(ranking[0])
			}
		}
		data.Results = append(data.Results, result)
		var output_result models.OutputResult
		output_result.Ranking = ranking
		output_result.Winner = winner

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Created New Result")
		json.NewEncoder(w).Encode(output_result)
	} else {
		json.NewEncoder(w).Encode("Too Early")
	}

}
