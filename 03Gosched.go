package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for {
			fmt.Println("我是go程")
		}
	}()
	for {
		runtime.Gosched() //出让cpu时间 当再次获得cpu时从出让位置执行
		fmt.Println("我是main")
	}
}
