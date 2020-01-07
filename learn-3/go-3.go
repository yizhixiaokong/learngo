package main

import (
	"fmt"
)

func main() {
	a := 1
	for a < 1000 {
		a += a
	}
	fmt.Println(a)
	for b := 1; b < 10; b += b {
		fmt.Println(b)
	}

	ma := make(map[string]int)
	ma["one"] = 1
	ma["two"] = 2
	//range读取map等数据
	for k, v := range ma {
		fmt.Println(k, v)
	}
	//用_吃掉多余不需要的返回值
	for _, v := range ma {
		fmt.Println(v)
	}
	fmt.Println(max(100, 999))
	fmt.Println(sumAndMax(100, 999))
	fmt.Println(sumAndMin(100, 999))
	fun(1, 2)
	fun(4, 5, 6, 7, 8)

	//传指针
	x := 99
	fmt.Println("x=", x)
	fmt.Println("x+1=", add(&x)) //传入x的地址

	//defer
	deferfun()

	//函数类型参数
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	odd := filter(slice, isOdd) //由于声明了testInt函数类型，可以将isOdd当做值来传递了
	fmt.Println("Odd : ", odd)
	even := filter(slice, isEven) //同上
	fmt.Println("Even : ", even)

}

//函数声明
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//多个返回值的函数
func sumAndMax(a, b int) (int, int) {
	if a > b {
		return a + b, a
	}
	return a + b, b
}

//命名返回值，返回时省略(推荐)
func sumAndMin(a, b int) (sum int, min int) {
	sum = a + b
	if a > b {
		min = b
	} else {
		min = a
	}
	return
}

//变参函数
func fun(arg ...int) { //传入的所有参数为int
	//函数体中，arg为一个int类型的slice
	for _, n := range arg {
		fmt.Println(n)
	}
}

//传指针函数
func add(a *int) int {
	*a = *a + 1
	return *a
}

//延迟语句defer
func deferfun() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) //defer是后进先出的，所以输出顺序是4,3,2,1,0
	}
}

//用type将所有拥有相同的参数，相同的返回值的函数声明为一种函数类型
type testInt func(int) bool

//判断奇数
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

//判断偶数
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

//将testInt作为参数传入filter函数
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

/*
import中的若干操作
	1.点操作
	import(
     	. "fmt"
	)
	这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的fmt.Println("hello world")可以省略的写成Println("hello world")
	2.别名操作
	import(
		f "fmt"
	)
	调用包函数时前缀变成了我们的前缀，即f.Println("hello world")
	3._操作
	import (
		"database/sql"
		_ "github.com/ziutek/mymysql/godrv"
	)
	引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数
*/
