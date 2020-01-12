package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数
	fmt.Println(r.Form) //输出到服务端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World! %s", time.Now()) //输出到客户端的信息


	name:=r.FormValue("name")
	pass:=r.FormValue("pass")
	w.Write([]byte("\n"+name+"   "+pass))//也是输出到客户端页面
}

func main() {
	http.HandleFunc("/", greet)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
