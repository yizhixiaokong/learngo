package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hellow,yizhixiaokong，你好一只小空")
	//变量的几种声明方法：
	var a, b, c int = 1, 3, 2
	var d, e, f = 6, 4, 5
	g, h, i := 8, 9, 7 //只能在函数内
	fmt.Println(a, b, c, d, e, f, g, h, i)
	//常量
	const Pi = 3.1415926
	//布尔值
	no := false
	if no {

	}

	//几种数值类型：
	//不同类型不能互相赋值

	//整型
	var j int = 1
	var k uint = 2
	var l rune = 3 //是int32的别称
	var m int8 = 4
	var n int16 = 5
	var o int32 = 6
	var p int64 = 7
	var q byte = 8 //是uint8的别称
	var r uint8 = 9
	var s uint16 = 10
	var t uint32 = 11
	fmt.Println(j, k, l, m, n, o, p, q, r, s, t)

	//浮点，默认为float64
	var u float32 = 12.12
	var v float64 = 13.14520666
	fmt.Println(u, v)

	//复数,默认为complex128
	var w complex64 = 15 + 16i  //32位实数+32位虚数
	var x complex128 = 17 + 18i //64位实数+64位虚数
	fmt.Println(w, x)

	//字符串
	var y string = "十九"
	var z string = `二十`
	//go字符串不可直接变
	//var s string = "hello"
	//s[0] = 'c'的写法是错误的
	//但是可以
	str := "hello"
	strr := []byte(str)
	strr[0] = 'c'
	str2 := string(strr)
	//'+'连接两个字符串
	str3 := y + z
	fmt.Println(y, z, str2, str3)

	//字符串切片
	stra := "xizhixiaokong"
	stra = "y" + stra[1:]
	fmt.Println(stra)

	//多行字符串，用 ` 声明
	strb := `	hello
	yizhixiaokong`
	fmt.Println(strb)

	//错误类型error ,处理错误信息的包errors
	err := errors.New("出错啦！")
	if err != nil {
		fmt.Print(err)
	}

}

var flag bool //默认值为false
var yes = true
