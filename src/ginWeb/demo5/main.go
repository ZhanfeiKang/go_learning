package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数 kua
	// 要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	kua := func(name string) (string, error) {
		return name + "帅~", nil
	}
	// 定义模板
	t := template.New("f.html") // 创建一个模板对象, 名字一定要与模板的名字能对应上
	// 告诉模板引擎，我现在付哦了一个自定义的函数kua
	t.Funcs(template.FuncMap{
		"kua99": kua,
	})
	// 解析模板
	_, err := t.ParseFiles("./f.html")
	// t, err := template.New("f").ParseFiles("./f.html")
	if err != nil {
		fmt.Println("parse template failed,err: ", err)
		return
	}
	name := "小王子"
	// 渲染模板
	t.Execute(w, name)
}

func demo1(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./t.html", "./ul.html")
	if err != nil {
		fmt.Println("parse template failed,err: ", err)
		return
	}
	name := "小王子"
	// 渲染模板
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmpDemo", demo1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err: ", err)
		return
	}
}
