package exercises

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var n = 0

func Ex2() {

	for i := 0; i < 10000; i++ {
		go func() {
			mu.Lock()
			n++
			mu.Unlock()
		}()
	}

	fmt.Println("Appuyez sur entrée")
	fmt.Scanln()
	fmt.Println("n:", n)
}

func Ex2_parallel() {
	done := make(chan int, 10000)
	for i := 0; i < cap(done); i++ {
		go func() {
			n++
			done <- 1
		}()
	}
	for i := 0; i < cap(done); i++ {
		<-done
	}
	fmt.Println("Appuyez sur entrée")
	fmt.Scanln()
	fmt.Println("n:", n)
}

func approve(ch chan bool) {
	for {
		if <-ch {
			n++
		}
	}
}

func Ex2_3() {
	ch := make(chan bool)
	go approve(ch)
	for i := 0; i < 10000; i++ {
		go func() {
			ch <- true
		}()
	}
	fmt.Println("Appuyez sur entrée")
	fmt.Scanln()
	fmt.Println("n:", n)
}
