package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//b, _ := ioutil.ReadFile("F:\\GoGinWeb\\src\\demo1\\hello.txt")
	b, err := ioutil.ReadFile(".\\hello.txt")
	if err != nil {
		fmt.Println("read file err: ", err)
	}
	fmt.Println(string(b))
	_, err = fmt.Fprintln(w, string(b))
	if err != nil {
		fmt.Println("err: ", err)
	}
	//fmt.Fprintln(w, "hello")
}

func main() {
	http.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server failed,err: ", err)
		return
	}
}
