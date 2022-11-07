package comsoc

import "errors"

func BordaSWF(p Profile) (count Count, err error) {
	var check, message = checkProfile(p)
	if !check {
		return count, message
	}
	var list_alts []Alternative
	for _, votant := range p {
		list_alts = append(list_alts, votant...)
		break
	}
	var result = make(map[Alternative]int)
	for _, alt := range list_alts {
		result[Alternative(alt)] = 0
	}
	for _, alt := range p {
		for i := 0; i <= len(alt)-1; i++ {
			result[Alternative(alt[i])] += len(alt) - 1 - i
		}
	}
	return result, errors.New("Finished")
}

func BordaSCF(p Profile) (bestAlts []Alternative, err error) {
	var check, message = checkProfile(p)
	if !check {
		return bestAlts, message
	}
	var resultSWF, _ = BordaSWF(p)

	return maxCount(resultSWF), errors.New("Finished")
}
