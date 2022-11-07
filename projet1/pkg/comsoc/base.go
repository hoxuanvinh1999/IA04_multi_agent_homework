package comsoc

import (
	"errors"
	"sort"
)

type Alternative int
type Profile [][]Alternative
type Count map[Alternative]int

// renvoie l'indice ou se trouve alt dans prefs
func rank(alt Alternative, prefs []Alternative) int {
	for i, v := range prefs {
		if alt == v {
			return i
		}
	}
	print("alt n'est pas dan cet prefs")
	return 0
}

// renvoie vrai ssi alt1 est préférée à alt2
func isPref(alt1, alt2 Alternative, prefs []Alternative) bool {
	if rank(alt1, prefs) > rank(alt2, prefs) {
		return true
	}
	return false
}

// renvoie les meilleures alternatives pour un décomtpe donné
func maxCount(count Count) (bestAlts []Alternative) {
	keys := make([]Alternative, 0, len(count))
	for key := range count {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return count[keys[i]] > count[keys[j]]
	})
	var best_value = int(count[keys[0]])
	var result []Alternative
	for key, value := range count {
		if value == best_value {
			result = append(result, key)
		}
	}
	//fmt.Println("best value", best_value)
	//fmt.Println("keys", keys)
	//fmt.Println("copy_input", count)
	//fmt.Println(result)
	return result
}

func minCount(count Count) (worstAlts []Alternative) {
	keys := make([]Alternative, 0, len(count))
	for key := range count {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return count[keys[i]] < count[keys[j]]
	})
	var worst_value = int(count[keys[0]])
	var result []Alternative
	for key, value := range count {
		if value == worst_value {
			result = append(result, key)
		}
	}
	//fmt.Println("best value", best_value)
	//fmt.Println("keys", keys)
	//fmt.Println("copy_input", count)
	//fmt.Println(result)
	return result
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative n'apparaît qu'une seule fois par préférences
func checkProfile(prefs Profile) (bool, error) {
	switch {
	case prefs == nil:
		return false, errors.New("Profile est null")
	default:
		var check []Alternative
		for _, votant := range prefs {
			check = append(check, votant...)
			break
		}
		for _, votant := range prefs {
			if len(votant) == 0 {
				return false, errors.New("Il y a un votant null")
			}
			if !duplicateInArray(votant) {
				return false, errors.New("Duplicate en votant")
			}
			var answer, message = checkVotant(check, votant)
			if !answer {
				return false, errors.New(message)
			}
		}
	}

	return true, errors.New("Correct Profile")
}

func duplicateInArray(alts []Alternative) bool {
	visited := make(map[Alternative]bool, 0)
	for i := 0; i < len(alts); i++ {
		if visited[alts[i]] == true {
			return false
		} else {
			visited[alts[i]] = true
		}
	}
	return true
}

// vérifie le profil donné, par ex. qu'ils sont tous complets et que chaque alternative de alts apparaît exactement une fois par préférences
func CheckProfileAlternative(prefs Profile, alts []Alternative) (bool, error) {
	switch {
	case prefs == nil:
		return false, errors.New("Profile est null")
	case alts == nil:
		return false, errors.New("Alts est null")
	case !duplicateInArray(alts):
		return false, errors.New("Duplicate en Alts")
	default:
		for _, votant := range prefs {
			if len(votant) == 0 {
				return false, errors.New("Il y a un votant null")
			}
			if !duplicateInArray(votant) {
				return false, errors.New("Duplicate en votant")
			}
			var answer, message = checkVotant(alts, votant)
			if !answer {
				return false, errors.New(message)
			}
		}

	}
	return true, errors.New("Correct ProfileAlternative")
}

func checkVotant(pref []Alternative, alts []Alternative) (bool, string) {
	if len(pref) != len(alts) {
		return false, "Vote incorrect: different in numbers of Alternatives"
	}
	k := 0
	for _, val := range alts {
		for _, value := range pref {
			if val == value {
				k++
			}
		}
	}
	if k != len(pref) {
		return false, "Vote incorrect: Duplicate alternative in vote OR vote wrong alternative"
	} else {
		return true, "Correct Vote"
	}
}
