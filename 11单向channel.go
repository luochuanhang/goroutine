package main

import "fmt"

func send(out chan<- int) { //只写channel
	out <- 123
	close(out)
}
func recv(in <-chan int) { //只读channel
	num := <-in
	fmt.Println("读到数据", num)
}
func main() {
	ch := make(chan int) //双向channel

	go send(ch) //双向channel转为写channel

	recv(ch) //双向channel转为写channel
}
