package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cond sync.Cond

func pro(out chan<- int, index int) { //生产者
	for {
		cond.L.Lock()
		for len(out) == 5 {
			cond.Wait()
		}
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("%d生产者生产%d\n", index, num)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 200)
	}

}
func con(in <-chan int, index int) { //消费者
	for {
		cond.L.Lock()
		for len(in) == 0 {
			cond.Wait()
		}
		num := <-in
		fmt.Printf("---%d消费者取出%d\n", index, num)
		cond.L.Unlock()
		cond.Signal()
		time.Sleep(time.Millisecond * 200)
	}
}
func main() {
	ch := make(chan int, 5)
	cond.L = new(sync.Mutex)
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 5; i++ {
		go pro(ch, i)
	}
	for i := 1; i < 4; i++ {
		go con(ch, i)
	}
	time.Sleep(time.Second * 10)
}
