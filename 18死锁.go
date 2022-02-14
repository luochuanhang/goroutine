package main

import "fmt"

func main() {
	//死锁不是锁是一种错误使用导致的现象
	/*死锁 channel 必须用于两个go程中  单go程自己死锁
	ch := make(chan int)
	ch <- 456
	num := ch
	fmt.Println(num)
	*/
	//死锁 go程间对channel访问顺序导致死锁
	//使用一端读写必须保证另外一端同时读写，否则死锁
	ch := make(chan int)
	num := <-ch

	go func() {
		ch <- 789

		fmt.Println(num)
	}()

	//

	/*多个channel交叉死锁
	//ago程获取bgo程的值写入
	//bgo程获取ago程的值写入
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case num := <-ch2:
				ch1 <- num
			}
		}
	}()

	for {
		select {
		case num := <-ch1:
			ch2 <- num
		}
	}
	*/
}
