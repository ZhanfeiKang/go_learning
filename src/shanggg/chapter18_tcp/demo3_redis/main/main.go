package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 对哈希的操作
// 一个一个放入，一个一个读取

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
	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("hset err: ", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("hset err: ", err)
		return
	}

	// 3.通过 go 向 redis 读取数据 string [key-val]
	name, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("hget err: ", err)
		return
	}

	age, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("hget err: ", err)
		return
	}

	fmt.Println("name: ", name, ", age: ", age)

	fmt.Println("succ...")
}
