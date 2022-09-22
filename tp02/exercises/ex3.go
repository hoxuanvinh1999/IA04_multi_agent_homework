package exercises

import (
	"fmt"
	"time"
)

func Ex3() {
	/*
		for i := 5; i >= 0; i-- {
			if i != 0 {
				fmt.Println(i)
				time.Sleep(time.Second)
			} else {
				fmt.Println("Bonne année ! ")
			}
		}
	*/
	/*
		for i := 5; i >= 0; i-- {
			select {
			case <-time.After(time.Second):
				if i != 0 {
					fmt.Println(i)
				} else {
					fmt.Println("Bonne année ! ")
				}
			}
		}
	*/
	i := 5
	for tick := range time.Tick(time.Second) {
		_ = tick
		if i != 0 {
			fmt.Println(i)
		} else {
			fmt.Println("Bonne année ! ")
			return
		}
		i--
	}
}
