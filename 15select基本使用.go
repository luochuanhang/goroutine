package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)    //用来进行通讯
	quit := make(chan bool) //用来判断是否退出

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true //通知主go程退出
		runtime.Goexit()
	}()
	for { //主go程监听数据流动
		select {
		case num := <-ch: //不可读阻塞，可读使用数据
			fmt.Println("主go程读到:", num) //模拟使用数据
		case <-quit: //不可读阻塞，可读将主go程退出
			fmt.Println("退出")
			return //终止进程
		}
		fmt.Println("----------") //select不支持循环机制需使用for循环
	}
}
