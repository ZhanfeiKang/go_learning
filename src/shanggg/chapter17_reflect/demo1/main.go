package main

import (
	"fmt"
	"reflect"
)

// 专门演示反射
func reflectTest01(b interface{}) {

	// 通过反射获取的传入的变量的 type ，kind，值
	// 1.先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rtype: ", rType)

	// 2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2: ", n2)
	fmt.Printf("rVal: %v \nrVal type: %T\n", rVal, rVal)

	// 3.下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	// 4. 将interface{}通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2: ", num2)
}

// 专门演示对结构体的反射
func reflectTest02(b interface{}) {

	// 通过反射获取的传入的变量的 type ，kind，值
	// 1.先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rtype: ", rType)

	// 2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	// 3.获取变量对应的Kind
	// (1) rVal.Kind()
	kind1 := rVal.Kind()
	// (2) rType.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind1: %v\nkind2: %v\n", kind1, kind2)

	// 4.下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV: %v \niV: %T\n", iV, iV)

	// 5. 将interface{}通过断言转成需要的类型
	// switch-type 断言形式  也可以更加灵活实现
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name: %v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {

	// 请编写一个案例，
	// 演示对(基本数据类型、interface{}、relect.Value)进行反射的基本操作

	// 1.先定义一个int
	var num int = 100
	reflectTest01(num)

	// 2.定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
