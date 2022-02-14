package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("获得数据：", num)
			case <-time.After(time.Second * 3):
				fmt.Println("超时")
				quit <- true
				fmt.Println("超时3秒，程序结束")
				return
			}
		}
	}()
	for i := 0; i < 4; i++ {
		ch <- i
		time.Sleep(time.Second * 4)
		fmt.Println(i)
	}
	<-quit
	fmt.Println("要完了")
}
