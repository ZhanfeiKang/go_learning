package main

import (
	"errors"
	"fmt"
)

// 使用数组模拟栈的使用

type Stack struct {
	MaxTop int    // 表示栈最大可以存放的个数
	Top    int    // 表示栈顶
	arr    [5]int // 数组模拟栈
}

// Push: 放入数据
func (s *Stack) Push(val int) (err error) {

	// 先判断栈是否满了
	if s.Top == s.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}

	s.Top++
	// 放入数据
	s.arr[s.Top] = val
	return
}

// 出栈
func (s *Stack) Pop() (val int, err error) {
	// 判断栈是否为空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}

	// 先取值，再 this.Top--
	val = s.arr[s.Top]
	s.Top--
	return val, nil
}

// 遍历栈，注意需要从栈顶开始遍历
func (s *Stack) List() {
	// 先判断栈是否为空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}

	fmt.Println("栈的情况如下：")
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.arr[i])
	}
}

func main() {

	stack := &Stack{
		MaxTop: 5,  // 表示最多存放5个数到栈中
		Top:    -1, // -1 表示空栈
	}

	// 入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	// 显示
	stack.List()

	val, _ := stack.Pop()
	fmt.Println("出栈val: ", val)

	// 显示
	stack.List()
}
