package data

import "ia04_projet1/pkg/models"

var Ballots = []models.Ballot{
	{
		BallotID: "vote1",
		Rule:     "majority",
		Deadline: "Tue Nov 10 23:00:00 UTC 2009",
		VoterIDs: []string{"agt1", "agt2", "agt3", "agt4"},
		Alts:     3,
	},
}

var OutputBallots = []models.OutputBallot{
	{
		BallotID: "vote1",
	},
}
