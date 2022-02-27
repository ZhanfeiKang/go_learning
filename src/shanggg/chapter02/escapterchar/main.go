package main

import "fmt" // fmt包中提供格式化，输出，输入的函数。
func main() {
	fmt.Println("tom\tjack")                                // 制表符	\t
	fmt.Println("hello\nworld")                             // 换行符	\n
	fmt.Println("F:\\go\\shanggg\\chapter02\\escapterchar") // \\
	fmt.Println("tom说\"i love you\"")                       // \"
	// \r   回车,从当前行的最前面开始输出，覆盖掉以前内容
	fmt.Println("天龙八部雪山飞狐\r张飞")

	fmt.Println("helloworldhelloworldhelloworldhelloworldhelloworld",
		"\nkkitekkite")
}
