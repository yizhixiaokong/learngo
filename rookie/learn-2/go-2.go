package main

//同时导多个包或声明多个变量时，可采用分组，用()括起多个

import (
	"fmt"
)

//关键字iota，这个关键字用来声明枚举的时候采用，它默认开始值是0，const中每增加一行加1
const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a = iota //a=0
	b = "B"
	c = iota //c=2

	//与空行或注释行无关

	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)

//go用大小写区分公有和私有，以大写字母开头的变量是公有变量，小写字母开头为私有变量，函数也是类似

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)

	//数组array
	var arr [10]int           //长度为10，默认均为0
	arr2 := [10]int{1, 2, 3}  //长度为10，前三个初始化，之后默认为0
	arr3 := [...]int{4, 5, 6} //根据元素个数自动计算
	//但数组长度是不可变的
	fmt.Println(arr, arr2, arr3)

	//嵌套数组(多维数组)
	doublearr := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	doublearr2 := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(doublearr, doublearr2)

	//切片slice
	//可理解为一个动态数组或声明不需要长度的数组
	//但其本质是一个引用类型，指向一个低层array
	var sli []int = make([]int, 3)
	sli1 := []int{1, 2, 3, 4}

	//slice可从一个数组或已存在的slice中再次声明，通过array[i:j]获取，
	//i是起始，j是结束，但不包括j，长度为j-i
	ar := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //一个长度10的数组
	var sli2, sli3 []int                        //两个slice
	sli2 = ar[2:5]
	sli3 = ar[3:5]
	fmt.Println(sli, sli1, sli2, sli3)

	//slice默认从0开始，ar[:n]等价于ar[0:n]
	//slice默认第二个数是数组长，ar[n:]等价于ar[n:len(ar)]
	//也可用sli:=arr[x:y:z]的形式指定slice的容量，容量为z-x

	//len()函数获得slice的长度，cap()函数获得slice的容量
	//未初始化的slice默认为nil,长度为0
	var nilsli []int
	if nilsli == nil {
		fmt.Println("空", len(nilsli), cap(nilsli), nilsli)
	}

	//字典(映射)map
	var ma map[string]int
	ma = make(map[string]int) //两种声明map的方式
	ma["one"] = 1
	ma["twe"] = 2
	//len()函数获取map中拥有的key的数量
	//delete()删除map元素
	fmt.Println(len(ma), ma)
	delete(ma, "one")
	fmt.Println(len(ma), ma)
	//map实质也是引用类型
}
