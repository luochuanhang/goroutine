package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccccccccccc")
	//return  返回当前函数调用  return之前的defer注册有效
	runtime.Goexit() //退出当前go程 之前注册的defer都有效
	fmt.Println("ddddddddddddd")
}
func main() {
	go func() {
		defer fmt.Println("aaaaaaaa")
		test()
		fmt.Println("bbbbbbbb")
	}()
	for {

	}
}
