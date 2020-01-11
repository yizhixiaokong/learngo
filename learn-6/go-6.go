package main

import (
	"fmt"
	"time"
)

//轻量级线程/协程 goroutine
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	//关键字go,启动一个goroutine
	go say("wow")
	say("yoo")

	//通道 channel
	//用chan关键字声明,同时必须用make创建
	ci := make(chan int, 1) //声明通道时要同时声明传输的数据类型,同时给大小为1的缓冲区

	//用<-操作符来接受和发送数据
	i := 5
	ci <- i
	fmt.Println(<-ci)

	c := make(chan int, 2) //修改2为1就报错，修改2为3可以正常运行，因为为1时缓冲区大小不够
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("")
	//可通过range读取channel
	ch := make(chan int)
	go func() {
		for i = 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for ii := range ch {
		fmt.Println(ii)
	}

	//用语法 v, ok := <-ch 检测channel是否被关闭
	//如果ok返回false，那么说明channel已经没有任何数据并且已经被关闭

	chh := make(chan int)
	go func() {
		for i = 0; i < 5; i++ {
			chh <- i
		}
		close(chh)
	}()

	for {
		if v, ok := <-chh; ok == true {
			fmt.Println("接收", v)
		} else {
			fmt.Println("chh已关闭")
			break
		}
	}
}
