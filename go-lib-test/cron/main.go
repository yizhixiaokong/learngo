package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

const format = "15:04:05"

type JobA struct {
	flag bool
}

func NewA() *JobA {
	return &JobA{
		flag: false,
	}
}

type JobB struct {
	flag bool
}

func NewB() *JobB {
	return &JobB{
		flag: false,
	}
}

func (a *JobA) Run() {
	if !a.flag {
		a.flag = true
		fmt.Println("start a...")
	}
	fmt.Println(time.Now().Format(format), "a is running...")
}
func (b *JobB) Run() {
	if !b.flag {
		b.flag = true
		fmt.Println("start b...")
	}
	fmt.Println(time.Now().Format(format), "b is running...")
}

func main() {

	c := cron.New(
		cron.WithSeconds(), // 默认是不带秒的，加上这个就带秒了
		cron.WithChain(
			cron.Recover(cron.DefaultLogger),            // 允许定时器执行过程中捕获异常，不会导致定时器停止
			cron.SkipIfStillRunning(cron.DefaultLogger), // 如果上一次任务还没有执行完，就跳过这次任务
		),
	)

	wg := sync.WaitGroup{}

	// 添加任务
	id, err := c.AddFunc("*/1 * * * * *", func() { // 每秒执行一次
		fmt.Println(time.Now().Format(format), "c is running...")
	})
	if err != nil {
		panic(err)
	}
	wg.Add(1)

	c.Start()
	fmt.Println(time.Now().Format(format), "start cron")

	// id 用于停止任务
	go func() {
		time.Sleep(3 * time.Second)
		c.Remove(id)
		fmt.Println("remove job ", id)
		wg.Done()
	}()

	wg.Wait()

	// 也可以使用 AddJob 添加任务
	// Job接口需要实现Run方法
	a := NewA()
	idA, err := c.AddJob("@every 1s", a) // 也有一些内置的时间表达式，详见文档
	if err != nil {
		panic(err)
	}
	wg.Add(1)

	b := NewB()
	idB, err := c.AddJob("*/2 * * * * *", b) // 两秒执行一次
	if err != nil {
		panic(err)
	}
	wg.Add(1)

	c.Start()
	fmt.Println(time.Now().Format(format), "start cron job")

	go func() {
		time.Sleep(8 * time.Second)
		c.Remove(idA)
		fmt.Println("remove job ", idA)
		wg.Done()
		c.Remove(idB)
		fmt.Println("remove job ", idB)
		wg.Done()
	}()

	// 等待
	wg.Wait()
}
