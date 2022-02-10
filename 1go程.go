package main

import (
	"fmt"
	"time"
)

func sing() {
	for i := 0; i < 50; i++ {
		fmt.Println("正在唱歌")
		time.Sleep(time.Millisecond)
	}
}
func dance() {
	for i := 0; i < 50; i++ {
		fmt.Println("正在跳舞...")
		time.Sleep(time.Millisecond)
	}
}

func main() {
	go sing()
	go dance()
	for {

	}
}
