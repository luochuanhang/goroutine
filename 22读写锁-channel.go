package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readaGo2(in <-chan int, index int) {
	for {
		num := <-in
		fmt.Printf("第%d个go程读到%d\n", index, num)
	}
}
func writeGo2(out chan<- int, index int) {

	for {
		num := rand.Intn(1000)
		fmt.Printf("第%d个go程写入%d\n", index, num)
		out <- num
		time.Sleep(time.Millisecond * 300)

	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)
	for i := 1; i < 5; i++ {
		go readaGo2(ch, i)
	}
	for i := 1; i < 3; i++ {
		go writeGo2(ch, i)
	}
	time.Sleep(time.Second * 5)
}
