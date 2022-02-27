package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Score float64
}

type stuSlice []Student

func (stus stuSlice) Len() int {
	return len(stus)
}

func (stus stuSlice) Less(i, j int) bool {
	return stus[i].Score > stus[j].Score // 从大到小
}

func (stus stuSlice) Swap(i, j int) {
	stus[i], stus[j] = stus[j], stus[i]
}

func main() {

	var stus stuSlice

	for i := 0; i < 10; i++ {
		temp, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", rand.Float64()*100), 64)
		stu := Student{
			Name:  fmt.Sprintf("stu:%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: temp,
		}

		stus = append(stus, stu)

	}

	for _, stu := range stus {
		fmt.Println(stu)
	}

	sort.Sort(stus) // 排序
	fmt.Println("------------------排序后-----------------")
	for _, v := range stus {
		fmt.Println(v)
	}

}
