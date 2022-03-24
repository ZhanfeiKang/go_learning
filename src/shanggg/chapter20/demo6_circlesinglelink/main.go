package main

import "fmt"

// 定义猫的结构体结点
type CatNode struct {
	no   int // 猫猫的编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	// 判断是不是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head // 形成一个环状
		fmt.Println(newCatNode, "加入到环形的链表中")
		return
	}

	// 定义一个临时的变量，帮忙找到环形的最后结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

// 输出环形链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表的情况如下: ")
	temp := head
	if temp.next == nil {
		fmt.Println("空空如也的环形链表")
		return
	}
	for {
		fmt.Println("猫的信息为: ", temp, "->")
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

// 删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head
	// 空链表
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，无法删除")
		return head
	}
	// 如果只有一个结点
	if temp.next == head { // 只有一个结点
		temp.next = nil
		return head
	}

	// 将helper 定位到环形链表最后
	for {
		if helper.next == head {
			break
		}
		helper = head.next
	}

	// 如果有包含两个以上的结点
	flag := true
	for {
		if temp.next == head { // 如果到这，说明比较到最后一个【最后一个还没比较】
			break
		}
		if temp.no == id {
			if temp == head { // 说明删除的是头结点
				head = head.next

			}
			// 恭喜找到, 我们也可以这里直接删除
			helper.next = temp.next
			fmt.Println("猫猫：", id)
			flag = false
			break
		}
		temp = temp.next     // 移动【比较】
		helper = helper.next // 移动【价值: 一旦找到要删出去的结点 helper】
	}
	// 这里还要比较一次
	if flag { // 如果flag为true，则我们在上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Println("猫猫：", id)
		} else {
			fmt.Println("sorry，没有no: ", id)
		}
	}

	return head
}

func main() {
	// 这里我们初始化一个环形链表的头结点
	head := &CatNode{}

	// 创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	ListCircleLink(head)

	head = DelCatNode(head, 3)
	fmt.Println()
	ListCircleLink(head)
}
