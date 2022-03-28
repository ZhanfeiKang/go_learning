package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.New("index.html").
		Delims("{[", "[}").
		ParseFiles("./index.html")
	if err != nil {
		fmt.Println("parse template failed,err: ", err)
		return
	}
	// 渲染模板
	name := "kkite"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("execute template failed,err: ", err)
		return
	}
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	// 解析模板之前定义一个自定义的函授safe
	t, err := template.New("xss.html").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.html")
	if err != nil {
		fmt.Println("parse template failed,err: ", err)
		return
	}
	// 渲染模板
	str1 := "<script>alert(123)</script>"
	str2 := "<a href='http://kkite.gitee.io/'>kkite的blog</a>"
	err = t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
	if err != nil {
		fmt.Println("execute template failed,err: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err: ", err)
		return
	}
}
