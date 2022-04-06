package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// context.withValue

type TraceCode string
type UserID string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")           // 类型转换
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
	userIdKey := UserID("USER_ID")
	UserId, ok := ctx.Value(userIdKey).(int64) // 在子goroutine中获取userID
	if !ok {
		fmt.Println("invalid userid code")
	}
	log.Printf("%s worker func...", traceCode)
	log.Printf("userid: %d worker func...", UserId)
LOOP:
	for {
		fmt.Printf("worker,trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code 传递给后续启动的gorountine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12334545")
	ctx = context.WithValue(ctx, UserID("USER_ID"), int64(24365524464334))
	log.Printf("%s main 函数", "12334545")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
