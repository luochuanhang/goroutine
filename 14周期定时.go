package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan bool) //创建一个判断是否终止的channel

	fmt.Println("当前时间:", time.Now())
	myTicker := time.NewTicker(time.Second)
	i := 0
	go func() {
		for {
			i++
			nowTime := <-myTicker.C
			fmt.Println("现在时间:", nowTime)
			if i == 4 {
				quit <- true //解除主go程阻塞
			}
		}
	}()
	<-quit
}
