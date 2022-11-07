package data

import (
	"ia04_projet1/pkg/comsoc"
	"ia04_projet1/pkg/models"
)

var Votes = []models.Vote{
	{
		VoteID:  "vote1",
		Prefs:   []comsoc.Alternative{1, 3, 2, 4},
		AgentID: "agt1",
		Options: []comsoc.Alternative{3},
	},
	{
		VoteID:  "vote1",
		Prefs:   []comsoc.Alternative{1, 4, 2, 3},
		AgentID: "agt2",
		Options: []comsoc.Alternative{3},
	},
	{
		VoteID:  "vote1",
		Prefs:   []comsoc.Alternative{1, 4, 2, 3},
		AgentID: "agt3",
		Options: []comsoc.Alternative{3},
	},
	{
		VoteID:  "vote1",
		Prefs:   []comsoc.Alternative{2, 4, 1, 3},
		AgentID: "agt4",
		Options: []comsoc.Alternative{3},
	},
}
