package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
			fmt.Println("子go程", i)
		}
		close(ch) //关闭channel
	}()

	for n := range ch {
		fmt.Println("读到数据", n)
	}

	for {
		if n, ok := <-ch; ok == true {
			fmt.Println("主go读到数据", n)
		} else {
			fmt.Println("子go程关闭channel")
			break
		}
	}

}
