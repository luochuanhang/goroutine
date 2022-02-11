package main

import (
	"fmt"
)

func fibonacci(ch chan int, quit chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Print(num, " ")
		case <-quit:
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go fibonacci(ch, quit)
	x := 1
	y := 1
	for i := 0; i < 40; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}
