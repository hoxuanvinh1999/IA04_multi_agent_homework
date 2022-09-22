package exercises

import (
	"fmt"
	"sync"
	"time"
)

const length int = 10000

func Fill_normal(tab []int, v int) []int {
	fmt.Println("Function Fill Normal")
	for i := 0; i < length; i++ {
		tab[i] = v
	}
	return tab
}

func Fill_parallel(tab []int, v int) []int {

	fmt.Println("Function Fill Parallel")
	var wc sync.WaitGroup
	for i := range tab {
		wc.Add(1)
		i2 := i
		go func() {
			defer wc.Done()
			tab[i2] = v
		}()
	}
	wc.Wait()
	return tab
}

func double(n int) int {
	return 2 * n
}

func isPositive(n int) bool {
	if n > 0 {
		return true
	}
	return false
}

func ForEach_normal(tab []int, f func(int) int) []int {
	fmt.Println("Function For Each Normal")
	Fill_normal(tab, 2)
	for i, v := range tab {
		tab[i] = f(v)
	}
	return tab
}

func ForEach_parallel(tab []int, f func(int) int) []int {
	fmt.Println("Function For Each Parallel")
	var wc sync.WaitGroup
	for i, v := range tab {
		i2 := i
		v2 := v
		wc.Add(1)
		go func() {
			defer wc.Done()
			tab[i2] = f(v2)
		}()
	}
	wc.Wait()
	return tab
}

func Copy_normal(src []int, dest []int) {
	fmt.Println("Function Copy Normal")
	var wc sync.WaitGroup
	for i := range src {
		i2 := i
		wc.Add(1)
		go func() {
			defer wc.Done()
			dest[i2] = src[i2]
		}()
	}
	wc.Wait()
	return
}

func Copy_parallel(src []int, dest []int) {
	fmt.Println("Function Copy Parallel")
	for i := range src {
		dest[i] = src[i]
	}
}

func Equal(tab1 []int, tab2 []int) bool {
	if len(tab1) != len(tab2) {
		return false
	}
	for i := range tab1 {
		if tab1[i] != tab2[i] {
			return false
		}
	}
	return true
}

func Find_normal(tab []int, f func(int) bool) (index int, val int) {
	for i, v := range tab {
		if f(tab[i]) {
			index = i
			val = v
			break
		}
	}
	return
}

func Ex4() {
	fmt.Println("This is execrice 4")
	tab_1 := make([]int, length)
	tab_2 := make([]int, length)
	tab_3 := make([]int, length)

	tab_1 = Fill_normal(tab_1, 2)

	tab_2 = Fill_parallel(tab_2, 2)

	Copy_parallel(tab_1, tab_3)

	tab_1[1] = 0
	tab_1[2] = -1
	fmt.Println(Find_normal(tab_1, isPositive))

	time.Sleep(2 * time.Second)
}
