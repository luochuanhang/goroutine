package main

import (
	"fmt"
	"time"
)

//3种定时方法
func main() {
	//1.time.sleep
	time.Sleep(time.Second)
	fmt.Println("当前时间:", time.Now())

	//2.Timer.c
	myTimer := time.NewTimer(time.Second * 2)
	nowTime := <-myTimer.C
	fmt.Println("现在时间:", nowTime)

	//3 time.after
	fmt.Println("当前时间：", time.Now())
	timer := <-time.After(time.Second * 2)
	fmt.Println("现在时间：", timer)

	mytime := time.NewTimer(time.Second) //创建定时器
	mytime.Reset(time.Second * 2)        //更新定时时长
	go func() {
		<-mytime.C
		fmt.Println("子go程结束")
	}()
	mytime.Stop() //设置定时停止
	for {

	}
}
