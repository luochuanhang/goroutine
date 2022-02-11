package main

import "fmt"

type OrderInfo struct {
	id int
}

func producer(out chan<- OrderInfo) { ///生成订单--生产者
	for i := 1; i < 10; i++ { //循环生成10个订单
		out <- OrderInfo{i} // 写入channel
	}
	close(out) //写完关闭channel
}

func consumer(in <-chan OrderInfo) { //处理订单--消费者
	for order := range in { //从channel中取出订单
		fmt.Println("消费者拿到", order.id) //模拟处理订单
	}
}
func main() {
	ch := make(chan OrderInfo)

	go producer(ch) //子go程生产者
	consumer(ch)    //主go程消费者
}
