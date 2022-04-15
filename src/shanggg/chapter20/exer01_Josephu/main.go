package main

import "fmt"

/*
	Josephu 问题：
	Josephu 问题为:	设编号为1,2，... n 的n个人围坐一圈，
					约定编号为k（1<=k<=n）的人从1开始报数，
				   	数到 m 的那个人出列，
					它的下一位又从1开始报数，
					数到m的那个人又出列，以此类推，直到所有人出列位置，由此产生一个出队编号的序列。

	提示：
		用一个不带头结点的循环列表来处理Josephu问题：
		先构建一个有n个结点的单循环链表，
		然后由k个结点起从1开始计数，记到m时，对应结点从链表中删除，
		然后再从被删除结点的下一个结点又从1开始计数，
		直到最后一个结点丛链表中删除，算法结束。
*/

// 小孩的结构体
type Boy struct {
	No   int  // 编号
	Next *Boy // 指向下一个小孩的指针[默认值是nil]
}

// 1.编写一个函数，构成单向的环形链表
// num : 表示小孩的个数
// *Boy : 返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {

	first := &Boy{}  // 空结点
	curBoy := &Boy{} // 空结点

	// 判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}

	// 循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}

		// 分析构成循环链表，需要一个辅助指针[帮忙的]
		// 1.因为第一个小孩比较特殊
		if i == 1 {
			first = boy // 不要动
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first // 构成环形链表
		}
	}

	return first
}

// 2.显示单向的环形链表[遍历]
func ShowBoy(first *Boy) {

	// 处理一下如果环形链表为空的情况
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩...")
		return
	}

	// 创建一个指针，帮助遍历[说明至少有一个小孩]
	curBoy := first
	for {
		fmt.Printf("No: %d -> ", curBoy.No)

		// 退出的条件
		if curBoy.Next == first {
			break
		}

		// curBoy 移动到下一个
		curBoy = curBoy.Next
	}
	fmt.Println()

}

// 分析思路
// 1. 编写一个函数，PlayGame(first *Boy, startNo int, countNum int)
// 2. 最后使用一个算法，按照要求，在环形链表中留下最后一个人
func PlayGame(first *Boy, startNo int, countNum int) {

	// 1.空链表单独处理
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	// 留一个，判断 startNo <= 小孩的总数
	// 2.需要定义一个辅助指针，帮助我们删除小孩
	tail := first
	// 3.让tail指向环形链表的最后一个小孩，这个非常的重要
	// 因为tail在删除小孩的时候需要用到
	for {
		if tail.Next == first { // 说明tail到了最后的小孩
			break
		}

		tail = tail.Next
	}

	// 4.让 first 移动到 startNo [后面我们删除小孩，就以first为准]
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	// 5.开始数countNum下，然后就删除first指向的小孩
	for {
		// 开始数 countNum-1次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈->\n", first.No)
		// 删除first指向的结点
		first = first.Next
		tail.Next = first

		// 判断如果 tail == first，圈中只有一个小孩。
		if tail == first {
			break
		}
	}
	fmt.Printf("最后出圈的小孩：%d\n", first.No)
}

func main() {

	first := AddBoy(500)
	// 显示
	ShowBoy(first)
	PlayGame(first, 20, 31)
}
