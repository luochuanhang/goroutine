package main

import (
	"fmt"
	"sync"
	"time"
)

/*使用channel完成同步
func printl(s string) {
	for _, a := range s {
		fmt.Printf("%c", a)
		time.Sleep(time.Second)
	}
}
func peron1(c chan int) {
	printl("hello")
	c <- 9
}
func peron2(c chan int) {
	<-c
	printl("world")
}
func main() {
	ch := make(chan int)
	go peron1(ch)
	go peron2(ch)
	for {

	}
}
*/
var m sync.Mutex

//创建mutex，访问数据之前加锁 访问数据结束解锁
//在ago程加锁期间，bgo程会加锁失败阻塞
func printl(s string) {
	m.Lock()
	for _, a := range s {
		fmt.Printf("%c", a)
		time.Sleep(time.Second)
	}
	m.Unlock()
}
func peron1() {
	printl("hello")
}
func peron2() {
	printl("world")
}
func main() {
	go peron1()
	go peron2()
	for {

	}
}
