package comsoc

import (
	"errors"
)

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	var check, message = checkProfile(p)
	if !check {
		return count, message
	}
	if len(thresholds) != len(p) {
		return count, errors.New("Invalid thresholds: length is not equal")
	}
	var list_alts []Alternative
	var dup_p [][]Alternative
	var result = make(map[Alternative]int)
	for _, votant := range p {
		list_alts = append(list_alts, votant...)
		break
	}
	for _, alt := range list_alts {
		result[Alternative(alt)] = 0
	}
	dup_p = append(dup_p, p...)
	for i, votant := range dup_p {
		if thresholds[i] > len(votant) || thresholds[i] < 0 {
			return count, errors.New("Invalid thresholds: element in thresholds")
		} else {
			votant = votant[0:thresholds[i]]
			for _, value := range votant {
				result[value] += 1
			}
		}
		//fmt.Println(i, votant)
	}
	return result, errors.New("Finished")
}

func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	var check, message = checkProfile(p)
	if !check {
		return bestAlts, message
	}

	if len(thresholds) != len(p) {
		return bestAlts, errors.New("Invalid thresholds: length is not equal")
	}
	var list_alts []Alternative
	for _, votant := range p {
		list_alts = append(list_alts, votant...)
		break
	}
	var dup_p [][]Alternative
	var resultSWF = make(map[Alternative]int)
	for _, alt := range list_alts {
		resultSWF[Alternative(alt)] = 0
	}
	dup_p = append(dup_p, p...)
	for i, votant := range dup_p {
		if thresholds[i] > len(votant) || thresholds[i] < 0 {
			return bestAlts, errors.New("Invalid thresholds: element in thresholds")
		} else {
			votant = votant[0:thresholds[i]]
			for _, value := range votant {
				resultSWF[value] += 1
			}
		}
	}
	var resultSCF = maxCount(resultSWF)
	switch {
	case len(resultSCF) == 1:
		return resultSCF, errors.New("Finished")
	case len(resultSCF) == 2:
		var tiebreak, err = TieBreak(resultSCF[0], resultSCF[1], dup_p)
		result := []Alternative{tiebreak}
		if len(result) == 1 {
			return result, errors.New("Finished")
		} else {
			return resultSCF, err
		}
	default:
		return
	}
}
