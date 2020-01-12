package main

import (
	"fmt"
	"log"
	"os"
)

//os包，文件相关操作

func main() {
	file, err := os.Open("go-8.go") //打开文件
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1000)
	len, err := file.Read(buf) //Read方法读取文件信息到buf
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(buf[0:len]))

	//Stat方法查看文件状态
	f, _ := os.Stat(`go-8.go`)
	fmt.Println(f.Name(), f.IsDir(), f.ModTime())

	file.Close() //记得关闭文件
}
