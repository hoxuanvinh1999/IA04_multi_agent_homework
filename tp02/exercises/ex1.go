package exercises

import (
	"fmt"
	"strconv"
	"time"
)

// exercice 1
func Compte(n int) {
	/*
		if n <= 0 || reflect.TypeOf(n).Kind() == reflect.Int {
			fmt.Println("n is not a int")
	}*/
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func CompteMsg(n int, msg string) {
	for i := 0; i < n; i++ {
		fmt.Println(msg)
		fmt.Println(i)
	}
}

func CompteMsgFromTo(start int, end int, msg string) {
	for i := start; i < end+1; i++ {
		fmt.Println(msg)
		fmt.Println(i)
	}
}

func Ex1() {
	//fmt.Println("testing")
	// go compteMsg(10, "world")
	// go compteMsg(10, "hello")
	for i := 1; i <= 10; i++ {
		var msg = "gouroutine" + strconv.Itoa(i)
		go CompteMsgFromTo((i-1)*10, i*10-1, msg)
	}
	time.Sleep(time.Second)
}
