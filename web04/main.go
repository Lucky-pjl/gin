package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter,r *http.Request)  {
	// 解析模板
	t,err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed,err:%v",err)
		return
	}
	// 渲染模板
	err = t.Execute(w,"pjl")
	if err != nil {
		fmt.Printf("execute template failed,err:%v",err)
		return
	}
}

func main() {
	http.HandleFunc("/hello",sayHello)
	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		fmt.Printf("http server failed,err:%v",err)
		return
	}
}
