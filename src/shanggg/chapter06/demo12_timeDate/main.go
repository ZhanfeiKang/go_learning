package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now_Type=%T\n", now, now)

	// 2.通过now可以获取到 年月日时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 3.格式化日期时间
	// 方式1：
	fmt.Printf("当前年月日：%02d-%02d-%02d %02d:%02d:%02d \n",
		now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d \n",
		now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dateStr=%v", dateStr)
	// 方式2：
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))

	// 4.每隔1毫秒打印一个数字，打印到20是就退出
	// 		每隔0.1秒打印一个数字，打印到20就退出
	i := 0
	for {
		i++
		fmt.Println(i)
		// 休眠
		// time.Sleep(time.Microsecond)
		time.Sleep(time.Microsecond * 100) // 0.1s

		if i == 20 {
			break
		}
	}

	// 5.获取当前 unix时间戳 和 unixnano 时间戳
	fmt.Printf("unix时间戳=%v, unixnano时间戳=%v", now.Unix(), now.UnixNano())

}
