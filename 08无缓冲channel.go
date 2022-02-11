package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			ch <- i
		}
	}()

	for i := 0; i < 5; i++ {

		n := <-ch
		fmt.Println(n)

	}
}
