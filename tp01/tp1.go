package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
)

// account for TD: ia04a041   KNt2tYgi

type Vertex struct {
	X, Y float64
}

func Abs(v *Vertex) float64 {
	v.X = 4
	v.Y = 3
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func Fill(sl []int) []int {
	var res []int
	for i := 0; i < len(sl); i++ {
		res = append(res, rand.Intn(100))
	}
	return res
}
func Moyenne(sl []int) float64 {
	var sum int
	if len(sl) == 0 {
		fmt.Println("null slice")
		return 0
	} else {
		for _, v := range sl {
			sum += v
		}
		result := sum / len(sl)
		return float64(result)
	}
}
func ValeursCentrales(sl []int) int {
	concat := sl
	length := len(concat)
	sort.Ints(concat)
	if len(concat) == 0 {
		fmt.Println("slice null")
		return 0
	} else if math.Mod(float64(len(concat)), 2) == 0 {
		return concat[int(float64((length/2))-0.5)]
	} else {
		return (concat[length/2] + concat[(length/2)-1]) / 2
	}

}

func Plus1(sl []int) []int {
	if len(sl) == 0 {
		fmt.Println("null slice")
		return sl
	} else {
		concat := sl
		for i := 0; i < len(concat); i++ {
			concat[i] += 1
		}
		return concat
	}
}
func compte(n int, tab []int) {
	if len(tab) == 0 {
		fmt.Println("tab null")
	} else if len(tab) == 1 {
		fmt.Println(tab[0])
		return
	} else {
		concat := tab[0:]
		sort.Ints(concat)
		check := concat[0]
		fmt.Println(tab[0])
		for i := 1; i < len(concat); i++ {
			if tab[i] > check {
				check = tab[i]
				fmt.Println(tab[i])
			}
		}

	}
}

// ------------------------------------------------------------
type Point2D struct {
	X, Y float64
}

func NewPoint2D(x, y float64) *Point2D {
	return &Point2D{x, y}
}
func (p *Point2D) SetX(X float64) {
	p.X = X
}

func (p *Point2D) getX() float64 {
	return p.X
}

func (p *Point2D) SetY(Y float64) {
	p.Y = Y
}

func (p *Point2D) Gety() float64 {
	return p.Y
}
func Module(p Point2D) float64 {
	return float64(math.Sqrt(p.X*p.X + p.Y*p.Y))
}

type Rectangle struct {
	hag, bad Point2D
}

func NewRectangle(x, y Point2D) *Rectangle {
	return &Rectangle{x, y}
}

func (p *Rectangle) Gethag() Point2D {
	return p.hag
}

func (p *Rectangle) Sethag(hag Point2D) {
	p.hag = hag
}

func (p *Rectangle) Getbad() Point2D {
	return p.bad
}

func (p *Rectangle) Setbad(bad Point2D) {
	p.bad = bad
}

type Sprite struct {
	position    Point2D
	hitbox      float64
	factor      float64
	nom_fichier string
}

//--------------------------------------------------------------

func IsPalindrome(word string) bool {
	reversedWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversedWord += string(word[i])
	}
	for i := range word {
		if word[i] != reversedWord[i] {
			return false
		}
	}
	return true
}
func Palindromes(words []string) (l []string) {
	var result []string
	for i := 0; i < len(words); i++ {
		if IsPalindrome(words[i]) {
			result = append(result, words[i])
		}
	}
	return result
}
func Footprint(s string) (footprint string) {
	var input []string
	for i := 0; i < len(s); i++ {
		input = append(input, string(s[i]))
	}
	sort.Strings(input)
	result := ""
	for i := 0; i < len(s); i++ {
		result += string(input[i])
	}
	return result
}

func Anagrams(words []string) (anagrams map[string][]string) {

	result := make(map[string][]string)
	var check string
	null := []string{}
	for i := 0; i < len(words); i++ {
		check = Footprint(words[i])
		if _, ok := result[check]; ok {
			result[check] = append(result[check], words[i])
		} else {
			result[check] = null
			result[check] = append(result[check], words[i])
		}
	}
	//stupid solution-------------------------------------------
	/*
		var footprints []string
		var uniquefootprint []string
		for i := 0; i < len(words); i++ {
			footprints = append(footprints, Footprint(words[i]))
		}
		sort.Strings(footprints)
		check := footprints[0]
		uniquefootprint = append(uniquefootprint, check)
		for i := 0; i < len(words); i++ {
			if check != footprints[i] {
				check = footprints[i]
				uniquefootprint = append(uniquefootprint, check)
			}
		}
		null := []string{}
		fmt.Println(len(uniquefootprint))
		for i := 0; i < len(uniquefootprint); i++ {
			result[uniquefootprint[i]] = null
			for j := 0; j < len(words); j++ {
				if Footprint(words[j]) == uniquefootprint[i] {
					result[uniquefootprint[i]] = append(result[uniquefootprint[i]], words[j])
				}
			}
		}
	*/
	return result
}
func longestPalindrome(words []string) string {
	var result string
	for i := 0; i < len(words); i++ {
		if IsPalindrome(words[i]) && len(words[i]) > len(result) {
			result = words[i]
		}
	}
	return result
}

func whatAnagrams(words []string, keyword string) (anagrams []string) {

	result := make(map[string][]string)
	var check string
	null := []string{}
	for i := 0; i < len(words); i++ {
		check = Footprint(words[i])
		if _, ok := result[check]; ok {
			result[check] = append(result[check], words[i])
		} else {
			result[check] = null
			result[check] = append(result[check], words[i])
		}
	}

	return result[Footprint(keyword)]
}

func longestAnagrams(words []string) []string {

	result := make(map[string][]string)
	var check string
	null := []string{}
	for i := 0; i < len(words); i++ {
		check = Footprint(words[i])
		if _, ok := result[check]; ok {
			result[check] = append(result[check], words[i])
		} else {
			result[check] = null
			result[check] = append(result[check], words[i])
		}
	}
	var check_length int = 0
	var key_result string
	for key, value := range result {
		if len(value) > check_length {
			check_length = len(value)
			key_result = key
		}
	}
	return result[key_result]
}

func PalindromesAnagrams(words []string) (anagrams map[string][]string) {

	result := make(map[string][]string)
	var check string
	null := []string{}
	for i := 0; i < len(words); i++ {
		check = Footprint(words[i])
		if IsPalindrome(words[i]) {
			if _, ok := result[check]; ok {
				result[check] = append(result[check], words[i])
			} else {
				result[check] = null
				result[check] = append(result[check], words[i])
			}
		}
	}
	return result
}

//--------------------------------------------------------------

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(&v))
	fmt.Println(v)
	//var i float64 = 0
	//for i < 1000 {
	//	if math.Mod(i, 2) == 0 {
	//		fmt.Println(i)
	//	}
	//	i += 1
	//}
	//exercices 2
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//test := [12]int{2, 1, 4, 5, 4, 9, 11, 14, 1, 2, 4, 5}
	//fmt.Println(Fill(primes[:]))
	//fmt.Println(Moyenne(test[:]))
	//sort.Ints(test[:])
	//fmt.Println(test)
	//fmt.Println(ValeursCentrales(primes[:]))
	//compte(1, test[:])
	//------------------------------------------------------------------
	//exercices 3
	//a := NewPoint2D(4, 5)
	//fmt.Println(*a)
	//a.SetX(7)
	//a.SetY(6)
	//fmt.Println(*a)
	//fmt.Println(Module(*a))
	//------------------------------------------------------------------
	//exercices 4
	dict := [...]string{"AGENT", "CHIEN", "COLOC", "ETANG", "ELLE", "GEANT", "NICHE", "RADAR"}
	for i := 0; i < len(dict); i++ {
		fmt.Println(IsPalindrome(dict[i]))
	}
	fmt.Println(Palindromes(dict[:]))
	for i := 0; i < len(dict); i++ {
		fmt.Println(Footprint(dict[i]))
	}
	fmt.Print(Anagrams(dict[:]))
	var data []string
	f, err := os.Open("dico-scrabble-fr.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Print(Anagrams(data))
	//fmt.Println(longestPalindrome(data))
	//fmt.Println(whatAnagrams(data, "AGENTS"))
	//fmt.Println(longestAnagrams(data))
	fmt.Println(PalindromesAnagrams(data))

}
