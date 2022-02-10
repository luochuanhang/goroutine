package main

import (
	"fmt"
)

//主go程结束 子go程随之退出
func main() {
	go func() { //创建一个子go程
		for i := 0; i < 5; i++ {
			fmt.Println("我是go程")
			//time.Sleep(time.Millisecond)
		}
	}()
	for i := 0; i < 5; i++ {
		fmt.Println("我是main")
	}
}
