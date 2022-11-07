package comsoc

import (
	"errors"
)

func CondorcetWinner(p Profile) (bestAlts []Alternative, err error) {
	var check, message = checkProfile(p)
	if !check {
		return bestAlts, message
	}
	var result = make(map[Alternative][]Alternative)
	var list_alts []Alternative
	for _, votant := range p {
		list_alts = append(list_alts, votant...)
		break
	}
	for i, alt := range list_alts {
		result[alt] = make([]Alternative, 0)
		//fmt.Println("now i and alt = ", i, alt)
		for j, value := range list_alts {
			//fmt.Println("------now j and val = ", j, value)
			switch {
			case j < i:
				result[alt] = append(result[alt], -result[value][i])
			case j == i:
				result[alt] = append(result[alt], 0)
			default:
				var tiebreak, _ = TieBreak(alt, value, p)
				switch {
				case tiebreak == alt:
					result[alt] = append(result[alt], 1)
				case tiebreak == value:
					result[alt] = append(result[alt], -1)
				case tiebreak == -1:
					result[alt] = append(result[alt], 0)
				}
			}
		}
	}
	//fmt.Println(result)
	var point = 0
	var result_condorcet []Alternative
	for key, value := range result {
		point = 0
		for _, val := range value {
			if val == -1 {
				break
			}
			point += 1
		}
		if point == len(list_alts) {
			result_condorcet = append(result_condorcet, key)
			return result_condorcet, errors.New("Finished")
		}
	}
	return result_condorcet, errors.New("No best alternative")
}
