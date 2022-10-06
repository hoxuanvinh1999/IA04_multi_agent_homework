package comsoc

import (
	"errors"
)

func contains(s []Alternative, e Alternative) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func compactAlts(alts []Alternative, p Profile) (compact_p Profile) {

	var result [][]Alternative
	var null_votant []Alternative
	for i, votant := range p {
		result = append(result, null_votant)
		for _, value := range votant {
			if contains(alts, value) {
				result[i] = append(result[i], value)
			}
		}
	}
	return result
}

func removeAlts(alts []Alternative, p Profile) (compact_p Profile) {

	var result [][]Alternative
	var null_votant []Alternative
	for i, votant := range p {
		result = append(result, null_votant)
		for _, value := range votant {
			if !contains(alts, value) {
				result[i] = append(result[i], value)
			}
		}
	}
	return result
}

func TieBreak(alt1 Alternative, alt2 Alternative, p Profile) (Alternative, error) {
	list_alts := []Alternative{alt1, alt2}
	var point_alt1 = 0
	var point_alt2 = 0
	for _, votant := range compactAlts(list_alts, p) {
		for _, value := range votant {
			if value == alt1 {
				point_alt1 += 1
				break
			} else {
				point_alt2 += 1
				break
			}
		}
	}

	switch {
	case point_alt1 > point_alt2:
		return alt1, errors.New("Finished")
	case point_alt2 > point_alt1:
		return alt2, errors.New("Finished")
	default:
		return -1, errors.New("Equal")
	}

}
