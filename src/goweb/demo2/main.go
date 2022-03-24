package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的处理器处理请求！", r.URL.Path)
}

func main() {
	myHandler := MyHandler{}
	http.Handle("/myHandler", &myHandler)
	// server := http.Server{
	// 	Addr:        ":8080",
	// 	Handler:     &myHandler,
	// 	ReadTimeout: 2 * time.Second,
	// }

	http.ListenAndServe(":8080", nil)
	// http.ListenAndServe()
}
