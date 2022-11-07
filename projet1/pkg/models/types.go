package models

import (
	"ia04_projet1/pkg/comsoc"
)

type Ballot struct {
	BallotID string   `json:"ballotID"`
	Rule     string   `json:"rule"`
	Deadline string   `json:"deadline"`
	VoterIDs []string `json:"voter-ids"`
	Alts     int      `json:"#alts"`
}

type Vote struct {
	VoteID  string               `json:"vote-id"`
	Prefs   []comsoc.Alternative `json:"prefs"`
	AgentID string               `json:"agent-id"`
	Options []comsoc.Alternative `json:"options"`
}

type Result struct {
	BallotID string `json:"ballot-id"`
}

type OutputBallot struct {
	BallotID string `json:"ballot-id"`
}

type OutputResult struct {
	Winner  int                  `json:"winner"`
	Ranking []comsoc.Alternative `json:"ranking"`
}
