package main

import (
	"fmt"
	"reflect"
)

/*
	练习：
	给一个变量 var v float64 = 1.2
	请使用反射来得到它的reflect.Value，然后获取对应的 Type，Kind和值 ，
	并将 reflect.Value 转换成 interface{}, 再将interface{}转换成float64.
*/

func reflectTest(b interface{}) {
	rVal := reflect.ValueOf(b)

	rType := rVal.Type()
	rKind := rVal.Kind()
	rf := rVal.Float()

	fmt.Println("rType: ", rType)
	fmt.Println("rKind: ", rKind)
	fmt.Println("rf: ", rf)

	iVal := rVal.Interface()

	num := iVal.(float64)
	fmt.Println("num: ", num)

}

func main() {
	var v float64 = 1.2
	reflectTest(v)
}
