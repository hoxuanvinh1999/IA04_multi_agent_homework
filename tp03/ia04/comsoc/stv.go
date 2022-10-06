package comsoc

import (
	"errors"
)

func STV_SWF(p Profile) (count Count, err error) {
	var check, message = checkProfile(p)
	if !check {
		return count, message
	}
	var resultSWF, _ = MajoritySWF(p)
	var worstAlts = minCount(resultSWF)
	var dup_p [][]Alternative
	dup_p = append(dup_p, p...)
	var list_alts []Alternative
	for _, votant := range p {
		list_alts = append(list_alts, votant...)
		break
	}
	var tour_number = len(list_alts) - 1
	var list_remove_alts []Alternative
	for len(resultSWF) >= 2 || tour_number != 0 {
		list_remove_alts = nil
		switch {
		case len(worstAlts) == 2:
			var tiebreak_result, _ = TieBreak(worstAlts[0], worstAlts[1], compactAlts(worstAlts, p))
			if tiebreak_result == worstAlts[0] {
				list_remove_alts = append(list_remove_alts, worstAlts[1])
			} else if tiebreak_result == worstAlts[1] {
				list_remove_alts = append(list_remove_alts, worstAlts[0])
			} else {
				list_alts = append(list_alts, worstAlts...)
			}
			dup_p = removeAlts(list_remove_alts, dup_p)
			resultSWF, _ = MajoritySWF(dup_p)
			worstAlts = minCount(resultSWF)
			tour_number -= 1
		case len(worstAlts) == 1:
			dup_p = removeAlts(worstAlts, dup_p)
			resultSWF, _ = MajoritySWF(dup_p)
			worstAlts = minCount(resultSWF)
			tour_number -= 1
		default:
			var tiebreak_result, _ = CondorcetWinner(compactAlts(worstAlts, dup_p))
			if len(tiebreak_result) == 1 {
				var list_remove_alts []Alternative
				for _, value := range worstAlts {
					if value != tiebreak_result[0] {
						list_remove_alts = append(list_remove_alts, value)
					}
				}
				dup_p = removeAlts(worstAlts, dup_p)
				resultSWF, _ = MajoritySWF(dup_p)
				worstAlts = minCount(resultSWF)
				tour_number -= 1
			} else {
				dup_p = removeAlts(worstAlts, dup_p)
				resultSWF, _ = MajoritySWF(dup_p)
				worstAlts = minCount(resultSWF)
				tour_number -= 1
			}
		}
	}
	return resultSWF, errors.New("Finished")
}

func STV_SCF(p Profile) (bestAlts []Alternative, err error) {
	var check, message = checkProfile(p)
	if !check {
		return bestAlts, message
	}
	var resultSWF, _ = STV_SWF(p)
	switch {
	case len(resultSWF) == 1:
		var final_result []Alternative
		for key := range resultSWF {
			final_result = append(final_result, key)
			break
		}
		return final_result, errors.New("Finished")
	case len(resultSWF) == 2:
		var list_alts []Alternative
		for key := range resultSWF {
			list_alts = append(list_alts, key)
		}
		var tiebreak_result, err = TieBreak(list_alts[0], list_alts[1], compactAlts(list_alts, p))
		if tiebreak_result != -1 {
			bestAlts = append(bestAlts, tiebreak_result)
			return bestAlts, err
		} else {
			return bestAlts, errors.New("No best alternatives")
		}
	default:
		var list_alts []Alternative
		for key := range resultSWF {
			list_alts = append(list_alts, key)
		}
		return CondorcetWinner(compactAlts(list_alts, p))
	}
}
