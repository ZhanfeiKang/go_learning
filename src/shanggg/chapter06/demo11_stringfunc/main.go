package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1)统计字符串的长度，按字节 len(str)
	str := "hello"

	str1 := "hello北" // len()=8
	fmt.Println(len(str))
	fmt.Println(len(str1))

	// 2)遍历字符串
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 3)字符串转ASCII
	n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换成的结果是", n)
	}

	// 4)整数ASCII转字符串
	str = strconv.Itoa(12345)
	fmt.Println("str=", str)

	// 5)字符串转[]byte: var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	// 6)[]byte 转 字符串：str = string([]byte{97,98,99})
	str = string([]byte{97, 98, 99})
	fmt.Println(str)

	// 7)10进制转2,8,16进制：str = strconv.FormatInt(123,2), 返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Println("123对应的二进制是：", str)
	str = strconv.FormatInt(123, 16)
	fmt.Println("123对应的16进制是：", str)

	// 8)查找子字符串是否在指定的字符串中：strings.Contains("seefood","foo")	// true
	b := strings.Contains("seefood", "foo")
	fmt.Println("b=", b)

	// 9)统计一个字符串有几个指定的子串：strings.Count("chinese","e")   // 4
	num := strings.Count("chisss", "ss") // 1
	fmt.Println("num=", num)

	// 10)不区分字母大小写的字符串比较(== 是区分字母大小写的)：fmt.Println(strings.EqualFold("abc","AbC"))
	fmt.Println(strings.EqualFold("abc", "AbC"))

	// 11)返回子串在字符串第一次出现的index值，如果没有返回-1：
	// strings.Index("NLT_abc", "abc")
	index := strings.Index("NLT_abc", "abc") // 4
	fmt.Println("index=", index)

	// 12)返回子串在字符串最后一次出现的index，如果没有返回-1
	index = strings.LastIndex("go golang sss", "go")
	fmt.Println("index=", index)

	// 13)将指定的子串替换成另外一个子串：
	// strings.Replace("go go hello", "go", "golang",n)
	// n可以指定你希望替换几个，如果n=-1表示全部替换
	str = strings.Replace("go go hello", "go", "golang", 1)
	fmt.Println(str)

	// 14)按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：
	// strings.Split("hello,world,ok", ",")
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Printf("strArr=%v\n", strArr)

	// 15)将字符串的字母进行大小写转换：
	str = strings.ToLower("Go")
	fmt.Println("ToLower:", str)
	str = strings.ToUpper("Go")
	fmt.Println("ToUpper:", str)

	// 16)将字符串左右两边的空格去掉
	str = strings.TrimSpace("   tn a loing   asd    ")
	fmt.Printf("str=%q\n", str)

	// 17)将字符串左右两边指定的字符去掉
	// strings.Trim("! he!llo! "," !")  // "he!llo" 将左右两边的 "!" 和 " " 去掉
	str = strings.Trim("! he!llo! ", " !")
	fmt.Printf("str=%q\n", str)

	// 18)将字符串左边指定的字符去掉
	// strings.TrimLeft("! hello!"," !")
	// 19)将字符串右边指定的字符去掉
	// strings.TrimRight("! hello!"," !")

	// 20)判断字符串是否以指定的字符串开头：
	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Println(b)

	// 21)判断字符串是否以指定的字符串结尾：
	b = strings.HasSuffix("brid.jpg", ".jpg")
	fmt.Println(b)

}
