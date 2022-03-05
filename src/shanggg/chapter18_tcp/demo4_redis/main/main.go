package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 对哈希的操作
// 一起放入，一起读取

func main() {
	// 通过 go 向 redis 写入数据和读取数据
	// 1.连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err: ", err)
		return
	}

	defer conn.Close()

	// 2.通过 go 向 redis 写入数据 string [key-val]
	_, err = conn.Do("HMSet", "user02", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("hset err: ", err)
		return
	}

	// 3.通过 go 向 redis 读取数据 string [key-val]
	r, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("hget err: ", err)
		return
	}

	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
	// fmt.Println("r: ", r) // 切片
	// fmt.Println("name: ", name, ", age: ", age)

	fmt.Println("succ...")
}
