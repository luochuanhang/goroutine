package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 4)
	fmt.Println("len(ch)", len(ch), "cap(ch)", cap(ch))
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i //存满数据之前不会阻塞
			fmt.Println("子go程", i, "len(ch)", len(ch), "cap(ch)", cap(ch))
		}
	}()
	//time.Sleep(time.Second)
	for i := 0; i < 20; i++ {
		num := <-ch
		fmt.Println("主go程:", num)
	}
}
