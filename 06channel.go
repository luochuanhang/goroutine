package main

import (
	"fmt"
	"time"
)

//全局定义channel，用来数据同步
var channel = make(chan int)

func printer(s string) {
	for _, ch := range s {
		fmt.Printf("%c", ch)
		time.Sleep(time.Millisecond)
	}
}
func person1() {
	printer("hellohello")
	channel <- 8
}
func person2() {
	<-channel //没有数据写入程序阻塞
	printer("nihao")
}
func main() {
	go person1()

	go person2()
	for {

	}
}
