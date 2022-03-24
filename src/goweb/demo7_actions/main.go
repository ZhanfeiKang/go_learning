package main

import (
	"goweb/demo7_actions/model"
	"html/template"
	"net/http"
)

// 测试if
func testIf(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	t := template.Must(template.ParseFiles("if.html"))
	age := 17
	// 执行
	t.Execute(w, age)
}

// 测试range
func testRange(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	t := template.Must(template.ParseFiles("range.html"))
	var emps []*model.Employee
	emp1 := model.Employee{
		ID:       1,
		LastName: "tom",
		Email:    "tom@tom.com",
	}
	emp2 := model.Employee{
		ID:       2,
		LastName: "jerry",
		Email:    "jerry@jerry.com",
	}
	emp3 := model.Employee{
		ID:       3,
		LastName: "mary",
		Email:    "mary@mary.com",
	}
	emps = append(emps, &emp1)
	emps = append(emps, &emp2)
	emps = append(emps, &emp3)
	// 执行
	t.Execute(w, emps)
}

// 测试with
func testWith(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	t := template.Must(template.ParseFiles("with.html"))
	// 执行
	t.Execute(w, "狸猫")
}

// 测试testTemplate
func testTemplate(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	t := template.Must(template.ParseFiles("template1.html", "template2.html"))
	// 执行
	t.Execute(w, "我能在两个文件中显示吗？")
}

// 测试define
func testDefine(w http.ResponseWriter, r *http.Request) {
	// 解析模板文件
	t := template.Must(template.ParseFiles("define.html"))
	// 执行
	t.ExecuteTemplate(w, "model", "")
}

// 测试define2
func testDefine2(w http.ResponseWriter, r *http.Request) {
	age := 17
	var t *template.Template
	if age < 18 {
		// 解析模板文件
		t = template.Must(template.ParseFiles("define2.html"))
	} else {
		// 解析模板文件
		t = template.Must(template.ParseFiles("define2.html", "content1.html"))
	}

	// 执行
	t.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/testIf", testIf)
	http.HandleFunc("/testRange", testRange)
	http.HandleFunc("/testWith", testWith)
	http.HandleFunc("/testTemplate", testTemplate)
	http.HandleFunc("/testDefine", testDefine)
	http.HandleFunc("/testDefine2", testDefine2)

	http.ListenAndServe(":8080", nil)
}
