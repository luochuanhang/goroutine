package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwm sync.RWMutex
var value int

func readaGo1(index int) {
	for {
		rwm.RLock()
		num := value
		fmt.Printf("----第%d个go程读到%d\n", index, num)
		rwm.RUnlock()
	}
}
func writeGo1(index int) {

	for {
		num := rand.Intn(1000)
		rwm.Lock()
		fmt.Printf("第%d个go程写入%d\n", index, num)
		value = num
		time.Sleep(time.Second)
		rwm.Unlock()

	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < 5; i++ {
		go readaGo1(i)
	}
	for i := 1; i < 3; i++ {
		go writeGo1(i)
	}
	time.Sleep(time.Second * 5)
}
