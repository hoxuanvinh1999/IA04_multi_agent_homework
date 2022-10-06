package comsoc

import (
	"errors"
)

func CopelandSWF(p Profile) (count Count, err error) {
	var check, message = checkProfile(p)
	if !check {
		return count, message
	}
	var result = make(map[Alternative]int)
	var list_alts []Alternative
	for _, votant := range p {
		list_alts = append(list_alts, votant...)
		break
	}
	for i, alt := range list_alts {
		result[alt] = 0
		//fmt.Println("now i and alt = ", i, alt)
		for j, value := range list_alts {
			//fmt.Println("------now j and val = ", j, value)
			if j != i {
				var tiebreak, _ = TieBreak(alt, value, p)
				switch {
				case tiebreak == alt:
					result[alt] += 1
				case tiebreak == value:
					result[alt] += -1
				}
			}
		}
	}
	//fmt.Println(result)
	return result, errors.New("Finished")
}
func CopelandSCF(p Profile) (bestAlts []Alternative, err error) {
	var check, message = checkProfile(p)
	if !check {
		return bestAlts, message
	}
	var resultSWF, _ = CopelandSWF(p)
	var resultSCF = maxCount(resultSWF)
	switch {
	case len(resultSCF) == 1:
		return resultSCF, errors.New("Finished")
	case len(resultSCF) == 2:
		list_alts := []Alternative{resultSCF[0], resultSCF[1]}
		var tiebreak, err = TieBreak(resultSCF[0], resultSCF[1], compactAlts(list_alts, p))
		result := []Alternative{tiebreak}
		if len(result) == 1 {
			return result, errors.New("Finished")
		} else {
			return resultSCF, err
		}
	default:
		var compact_profile = compactAlts(resultSCF, p)
		var final_result, err = CondorcetWinner(compact_profile)
		return final_result, err
	}
}
