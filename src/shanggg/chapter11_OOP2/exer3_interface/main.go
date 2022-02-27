package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1.声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2.声明一个Hero结构体的切片类型
type HeroSlice []Hero

// 3.实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less()方法就是决定你使用什么标准进行排序
// 1.按hero的年龄从小到大（升序）排序!!
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp

	// 上面三行 等价于：
	hs[i], hs[j] = hs[j], hs[i]
}

// 1.声明Student结构体 如果Student实现了接口，则也可以进行排序
// type Student struct {
// 	Name  string
// 	Age   int
// 	Score float64
// }

func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	// 要求对intSlice切片进行排序
	// 1.冒泡排序
	// 2.使用系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 对结构体切片进行排序
	// 1.冒泡排序
	// 2.使用系统提供的方法

	// 测试看看我们是否可以对结构体切片排序
	var heroes HeroSlice

	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		// 将hero append到heros切片中
		heroes = append(heroes, hero)
	}

	// 看看排序前的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

	// 调用sort.Sort
	sort.Sort(heroes) // 因为Hero实现了Interface接口三个方法，所以才能放进来

	fmt.Println("---------------排序后----------------")
	for _, v := range heroes {
		fmt.Println(v)
	}
}
