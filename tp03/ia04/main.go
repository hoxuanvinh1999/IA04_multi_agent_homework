package main

import (
	"fmt"
	"tp03/comsoc"
)

func main() {
	comsoc.Test(1, 2)
	//basket := map[comsoc.Alternative]int{1231: 9, 7: 7,3: 3, 9: 9}

	prefs := [][]comsoc.Alternative{{1, 3, 2}, {1, 2, 3}, {3, 2, 1}}
	prefs2 := [][]comsoc.Alternative{{2, 3, 1}, {1, 2, 3}, {3, 1, 2}}
	thresholds := []int{2, 1, 2}
	//test_alts := []comsoc.Alternative{1, 2}
	//alts := []comsoc.Alternative{1, 2, 3}
	//fmt.Println(comsoc.CheckProfileAlternative(prefs, alts))
	//fmt.Println(comsoc.CheckProfile(prefs))
	fmt.Print("MajoritySWF: ")
	fmt.Println(comsoc.MajoritySWF(prefs))
	fmt.Print("MajoritySCF: ")
	fmt.Println(comsoc.MajoritySCF(prefs))
	fmt.Print("BordaSWF: ")
	fmt.Println(comsoc.BordaSWF(prefs))
	fmt.Print("BordaSCF: ")
	fmt.Println(comsoc.BordaSCF(prefs))
	fmt.Print("ApprovalSWF: ")
	fmt.Println(comsoc.ApprovalSWF(prefs, thresholds))
	fmt.Print("ApprovalSCF: ")
	fmt.Println(comsoc.ApprovalSCF(prefs, thresholds))
	fmt.Print("CondorcetWinner: ")
	fmt.Println(comsoc.CondorcetWinner(prefs))
	fmt.Print("CondorcetWinner: ")
	fmt.Println(comsoc.CondorcetWinner(prefs2))
	fmt.Print("CopelandSWF: ")
	fmt.Println(comsoc.CopelandSWF(prefs))
	fmt.Print("CopelandSCF: ")
	fmt.Println(comsoc.CopelandSCF(prefs2))
	fmt.Print("STV_SWF: ")
	fmt.Println(comsoc.STV_SWF(prefs))
	fmt.Print("STV_SCF:")
	fmt.Println(comsoc.STV_SCF(prefs))

	prefs3 := [][]comsoc.Alternative{{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{2, 3, 4, 1, 5},
		{2, 3, 4, 1, 5},
		{4, 5, 3, 2, 1},
		{4, 5, 3, 2, 1},
		{4, 5, 3, 2, 1},
		{3, 4, 2, 5, 1}}
	fmt.Print("CondorcetWinner Ex3 tp04 ")
	fmt.Println(comsoc.CondorcetWinner(prefs3))
}
