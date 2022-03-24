package main

import (
	"encoding/json"
	"fmt"
	"goweb/demo5_web_req/model"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求地址是: ", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求地址后的查询字符串是: ", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding信息是：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值是：", r.Header.Get("Accept-Encoding"))

	// // 获取请求体中内容的长度
	// len := r.ContentLength
	// // 创建byte切片
	// body := make([]byte, len)
	// // 将 请求体Body 中的内容读到 body中
	// r.Body.Read(body)
	// // 在浏览器中显示请求体中的内容
	// fmt.Fprintln(w, "请求体中的内容有：", string(body))

	// 注意：Body只能读一次！！！所以上面的需要注释掉
	// 解析表单, 在调用 r.Form 之前必须执行该操作
	// r.ParseForm()
	// 获取请求参数
	// 如果 form表单的action属性的URL地址中也有与form表单参数名相同的请求参数，
	// 那么两个种方式的参数值都可以得到，并且form表单中的参数值在URL参数值的前面
	// fmt.Fprintln(w, "请求参数有：", r.Form)
	// fmt.Fprintln(w, "POST请求的form表单中的请求参数有：", r.PostForm)

	// 通过直接调用 FormValue 方法和 PostFormValue 方法直接获取
	fmt.Fprintln(w, "URL中的user请求参数的值是：", r.FormValue("user"))
	fmt.Fprintln(w, "Form表单中的username请求参数的值是：", r.PostFormValue("username"))
}

func testJsonRes(w http.ResponseWriter, r *http.Request) {
	// 设置响应内容的类型
	w.Header().Set("Content-Type", "application/json")
	// 创建User
	user := model.User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Email:    "admin@kkite.com",
	}
	// 将 user 转换为 json 格式
	json, _ := json.Marshal(&user)
	// 将 json 格式的数据响应给客户端
	w.Write(json)
}

func testRedire(w http.ResponseWriter, r *http.Request) {
	// 设置响应头中的Location属性
	w.Header().Set("Location", "http://kkite.gitee.io/")
	// 设置响应的状态码
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/testJson", testJsonRes)
	http.HandleFunc("/testRedirect", testRedire)

	http.ListenAndServe(":8080", nil)
}
