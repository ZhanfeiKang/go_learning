package main

import "fmt"

// 编写一个学生考试系统

type Student struct {
	Name  string
	Age   int
	Score int
}

// 将Pupil和Graduate共有的方法也绑定到Student
func (stu *Student) showInfo() {
	fmt.Printf("学生姓名:%v\n年龄:%v\n成绩:%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	// 业务判断..省略
	stu.Score = score
}

// 小学生
type Pupil struct {
	Student //嵌入了Student匿名结构体
}

// Pupil特有的方法保留
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中.....")
}

// 大学生
type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中.....")
}

func main() {
	// 当我们对结构体嵌入了匿名结构体后，使用方法发生了变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.Age = 8

	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.showInfo()

	//------
	graduate := &Graduate{}
	graduate.Student.Name = "mary~"
	graduate.Student.Age = 20

	graduate.testing()
	graduate.Student.SetScore(90)
	graduate.Student.showInfo()
}
