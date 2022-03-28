package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("parse template failed,err: ", err)
		return
	}
	msg := "小王子"
	// 渲染模板
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./home.html")
	if err != nil {
		fmt.Println("parse template failed,err: ", err)
		return
	}
	msg := "小王子"
	// 渲染模板
	t.Execute(w, msg)
}
func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板	(模板继承的方式)
	// 解析模板
	t, err := template.ParseFiles("./templates/base.html", "./templates/index2.html")
	if err != nil {
		fmt.Println("parse template failed,err: ", err)
		return
	}
	msg := "小王子"
	// 渲染模板
	t.ExecuteTemplate(w, "index2.html", msg)
}

func home2(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./templates/base.html", "./templates/home2.html")
	if err != nil {
		fmt.Println("parse template failed,err: ", err)
		return
	}
	msg := "kkite"
	// 渲染模板
	t.ExecuteTemplate(w, "home2.html", msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err: ", err)
		return
	}
}
