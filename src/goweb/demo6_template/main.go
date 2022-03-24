package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	// t, _ := template.ParseFiles("index.html")
	// 通过 Must 函数让 Go 帮我们自动处理异常
	t := template.Must(template.ParseFiles("index.html", "index2.html"))
	// 执行
	// t.Execute(w, "Hello Template")
	// 将响应数据在 index2.html 文件中显示
	t.ExecuteTemplate(w, "index2.html", "Hello index2.html")
}

func main() {
	http.HandleFunc("/testTemplate", testTemplate)

	http.ListenAndServe(":8080", nil)
}
