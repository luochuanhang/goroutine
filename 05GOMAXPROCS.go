package main

import (
	"fmt"
	"runtime"
)

//设置当前进程使用的最大cpu核数
func main() {
	n := runtime.GOMAXPROCS(2)
	fmt.Println("n=", n)
	for {
		go fmt.Print(0) //子go程
		fmt.Print(1)
	}
}
