package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int // 数组-> 模拟队列
	front   int    // 指向队列首部
	rear    int    // 表示指向队列的尾部
}

func (q *Queue) AddQueue(val int) (err error) {

	// 先判断队列是否已满
	if q.rear == q.maxSize-1 { // 重要的提示：rear 是队列尾部(含最后元素)
		return errors.New("queue full")
	}

	q.rear++ // rear 后移
	q.array[q.rear] = val
	return
}

// 从队列中取出数据
func (q *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if q.rear == q.front { // 队空
		return -1, errors.New("queue empty")
	}
	q.front++
	val = q.array[q.front]
	return val, err
}

// 显示队列, 找到队首，然后遍历到队尾
func (q *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	// q.front 不包含队首的元素
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.array[i])
	}
	fmt.Println()
}

func main() {

	// 先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入 add 表示添加数据到队列")
		fmt.Println("2. 输入 get 表示从队列获取数据")
		fmt.Println("3. 输入 show 表示显示队列")
		fmt.Println("4. 输入 exit 退出")

		fmt.Printf("请输入: ")
		fmt.Scanln(&key)

		switch key {
		case "add":
			fmt.Printf("请输入你要入队列数: ")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列 ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出的val: ", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
