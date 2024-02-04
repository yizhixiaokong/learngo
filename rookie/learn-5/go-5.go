package main

import "fmt"

//接口interface
//go中可将一系列具有共性的方法定义在一起成为一个接口
//若某对象实现了某个接口的所有方法，即实现了这个接口

type human struct {
	name  string
	age   int
	phone string
}

type student struct {
	human
	school string
	money  float32
}

type employee struct {
	human
	campany string
	money   float32
}

//human对象的hello方法
func (h *human) hello() {
	fmt.Printf("I am %s.\n", h.name)
}

//human对象的sing方法
func (h *human) sing(lyrics string) {
	fmt.Println("lalalala", lyrics)
}

//student对象的borrowmoney方法
func (s *student) borrowmoney(mon float32) {
	s.money += mon
}

//ememployee对象的spendsalary方法
func (e *employee) spendsalary(mon float32) {
	e.money -= mon
}

//men接口
type men interface {
	hello()
	sing(lyrics string)
}

//young接口
type young interface {
	hello()
	sing(song string)
	borrowmoney(mon float32)
}

//older接口
type older interface {
	hello()
	sing(song string)
	spendsalary(mon float32)
}

//一个接口可以被任意多个对象实现，一个对象也可以实现任意多个接口
//所有对象都实现了空接口(interface{}),即实现了0个方法的接口

//employee重载的hello方法
func (e *employee) hello() {
	fmt.Printf("I am %s ,work on %s\n", e.name, e.campany)
}

func main() {
	a := student{human{"a", 18, "123456"}, "the A 大学", 0}
	b := student{human{"b", 19, "789101"}, "the B 大学", 100}
	c := employee{human{"c", 22, "12335224"}, "the C 公司", 1000}
	d := employee{human{"d", 29, "97941355"}, "the D 公司", 199999}

	var i men //由于human,student,employee都实现了men接口
	i = &a    //所以men类型的变量i可以接受student类型的变量
	i.hello()
	i.sing("balalala")

	i = &c //也可以接受employee类型的变量
	i.hello()
	i.sing("la lalalal ba ")

	xx := make([]men, 3)
	xx[0], xx[1], xx[2] = &b, &d, &a
	for _, val := range xx {
		val.hello()
	}
}
