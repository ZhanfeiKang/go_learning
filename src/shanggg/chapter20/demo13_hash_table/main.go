package main

import (
	"fmt"
	"os"
)

// 题目：
// 有一个公司，当有新的员工来报道时，要求将员工的信息加入(id, 性别, 年龄, 住址..)，
// 当输入该员工的id时，要求查找到该员工的所有信息。

// 要求:
// 1.不使用数据库，尽量节省内存，速度越快越好 => 哈希表（散列）
// 2.添加时，保证按照雇员的id从低到高插入

// 思路分析：
// 1.使用链表来实现哈希表，该链表不带表头
// 2.思路分析并画出示意图
// 3.代码实现[增删改查(显示所有员工，按id查询)]

// 定义Emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 方法待定
func (e *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员id=%d\n", e.Id%7, e.Id)
}

// 定义 EmpLink
// 我们这里的EmpLink 不带表头，即第一个结点就存放雇员
type EmpLink struct {
	Head *Emp
}

// 方法待定
// 1.添加员工的方法,保证添加时，编号是从小到大
func (el *EmpLink) Insert(emp *Emp) {

	cur := el.Head     // 辅助指针
	var pre *Emp = nil // 辅助指针 pre 在 cur 前面
	// 如果当前的 EmpLink 就是一个空链表
	if cur == nil {
		el.Head = emp // 完成
		return
	}
	// 如果不是一个空链表，给 emp 找到对应的位置并插入
	// 思路是让 cur 和 emp 比较，然后让 pre 保持在 cur 前面
	for {
		if cur != nil {
			// 比较
			if cur.Id > emp.Id {
				// 找到位置
				break
			}
			pre = cur // 保证同步
			cur = cur.Next
		} else {
			break
		}
	}

	// 退出时，我们看下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

// 显示链表的信息
func (el *EmpLink) ShowLink(no int) {
	if el.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}

	// 遍历当前链表并显示数据
	cur := el.Head // 辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表%d,雇员id=%d 名字=%s -> ", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

// 根据id查找对应的雇员，如果没有就返回nil
func (el *EmpLink) FindById(id int) *Emp {
	cur := el.Head

	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}

	return nil
}

// 定义hashTable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 给 HashTable 编写 Insert 雇员 的方法。
func (ht *HashTable) Insert(emp *Emp) {
	// 使用散列函数，确定将该雇员添加到哪个链表
	linkNo := ht.HashFun(emp.Id)
	// 使用对应的链表添加
	ht.LinkArr[linkNo].Insert(emp)
}

// 编写方法，显示hashTable所有的雇员
func (ht *HashTable) ShowAll() {
	for i := 0; i < len(ht.LinkArr); i++ {
		ht.LinkArr[i].ShowLink(i)
	}
}

// 编写一个散列方法
func (ht *HashTable) HashFun(id int) int {
	return id % 7 // 得到一个值，就是对应的链表的下标
}

// 编写一个方法，完成查找
func (ht *HashTable) FindById(id int) *Emp {
	// 使用散列函授，确定该雇员在哪个链表
	linkNo := ht.HashFun(id)
	return ht.LinkArr[linkNo].FindById(id)
}

func main() {

	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("=================雇员系统菜单=================")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出系统")
		fmt.Print("请输入你的选择: ")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Printf("输入雇员id: ")
			fmt.Scanln(&id)
			fmt.Printf("输入雇员name: ")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Print("请输入id号: ")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				// 编写一个方法，显示雇员信息
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误...")
		}
	}
}
