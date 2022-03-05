package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

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
	_, err = conn.Do("Set", "name", "tomjerry")
	if err != nil {
		fmt.Println("set err: ", err)
		return
	}

	// 3.通过 go 向 redis 读取数据 string [key-val]
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err: ", err)
		return
	}
	// 因为返回r是 interface{}
	// 因为 name 对应的值是string，因此我们需要转换
	// nameString := r.(string) 报错

	fmt.Println("succ...", r)
}
