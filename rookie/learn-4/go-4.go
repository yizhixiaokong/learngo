package main

import (
	. "fmt"
)

//结构体struct
type person struct {
	name  string
	age   int
	phone string
}

//匿名字段
type student struct {
	person //匿名字段，默认student包含person的所有字段
	number string
	phone  string //同名字段最外层优先访问
	int           //内置类型和其他自定义类型都能作为匿名字段
}

//方法method
func (s student) print() {
	println(s.name, s.age, s.person.phone, s.number, s.phone, s.int)
}

//传指针
func (s *student) setnumber(str string) {
	s.number = str //不需要*s.number = str
	//当你用指针去访问相应的字段时(虽然指针没有任何的字段)，Go知道你要通过指针去获取这个值
}

//方法继承
func (p *person) myname() {
	Printf("my name is " + p.name + "\n")
}

//方法重写
func (p *person) myphone() {
	Printf("my phone number is " + p.phone + "\n")
}
func (s *student) myphone() {
	Printf("my school phone number is " + s.phone + "\n")
}

func main() {
	a := student{person{"hxx", 22, "7758258"}, "5712345679", "18166668888", 99}
	Println(a.name, a.age, a.number, a.int)
	Println(a.person.name, a.person.age) //以person作为student的字段名
	Println(a.phone, a.person.phone)     //同名优先最外层，通过字段访问内层

	//方法
	a.print()

	a.setnumber("5700000000") //如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
	//类似，如果一个method的receiver是T，你可以在一个*T类型的变量P上面调用这个method，而不需要 *P去调用这个method
	a.print()

	a.myname()         //person是student的匿名字段，student继承了person的myname方法
	a.myphone()        //student重写了person的myphone方法
	a.person.myphone() //通过字段调用person的方法

	//方法也是通过首字母大小写区分共有与私有，大写为共有

}
