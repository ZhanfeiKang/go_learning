package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int    // 4
	array   [5]int // 数组
	head    int    // 指向队首
	tail    int    // 指向队尾
}

// 入队列 AddQueue(push)
func (cq *CircleQueue) Push(val int) (err error) {
	if cq.IsFull() {
		return errors.New("queue full")
	}
	// 最后一个位置空了，实际大小只有 maxSize - 1
	cq.array[cq.tail] = val // 把值给尾部
	cq.tail = (cq.tail + 1) % cq.maxSize
	return
}

// 出队列
func (cq *CircleQueue) Pop() (val int, err error) {
	if cq.IsEmpty() {
		return 0, errors.New("queue empty")
	}

	// 取出, head 指向队首，并且含队首的元素
	val = cq.array[cq.head]
	cq.head = (cq.head + 1) % cq.maxSize
	return
}

// 显示队列
func (cq *CircleQueue) ListQueue() {

	fmt.Println("环形队列情况如下:")
	// 取出当前队列有多少个元素
	size := cq.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	// 设计一个辅助变量，指向head
	tempHead := cq.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, cq.array[tempHead])
		tempHead = (tempHead + 1) % cq.maxSize
	}
	fmt.Println()
}

// 判断环形队列为满
func (cq *CircleQueue) IsFull() bool {
	return (cq.tail+1)%cq.maxSize == cq.head
}

// 判断是否为空
func (cq *CircleQueue) IsEmpty() bool {
	return cq.tail == cq.head
}

// 取出环形队列有多少个元素
func (cq *CircleQueue) Size() int {
	// 这是一个关键的算法
	return (cq.tail + cq.maxSize - cq.head) % cq.maxSize
}

func main() {
	// 先创建一个环形队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1. 输入 push 表示添加数据到队列")
		fmt.Println("2. 输入 pop 表示从队列获取数据")
		fmt.Println("3. 输入 show 表示显示队列")
		fmt.Println("4. 输入 exit 退出")

		fmt.Printf("请输入: ")
		fmt.Scanln(&key)

		switch key {
		case "push":
			fmt.Printf("请输入你要入队列数: ")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列 ok")
			}
		case "pop":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出的val: ", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
