package comsoc

import (
	"errors"
)

func getBestAlternative(alts []Alternative) Alternative {
	for _, value := range alts {
		return value
	}
	return 0
}

func MajoritySWF(p Profile) (count Count, err error) {
	var check, message = checkProfile(p)
	if !check {
		return count, message
	}
	var list_alts []Alternative
	for _, votant := range p {
		list_alts = append(list_alts, votant...)
		break
	}
	//fmt.Println(list_alts)
	var result = make(map[Alternative]int)
	for _, alt := range list_alts {
		result[Alternative(alt)] = 0
	}
	//fmt.Println(result)
	for _, alt := range p {
		result[getBestAlternative(alt)] += 1
	}
	//fmt.Println(result)
	return result, errors.New("Finished")
}

func MajoritySCF(p Profile) (bestAlts []Alternative, err error) {
	var check, message = checkProfile(p)
	if !check {
		return bestAlts, message
	}
	var resultSWF, _ = MajoritySWF(p)

	return maxCount(resultSWF), errors.New("Finished")
}
