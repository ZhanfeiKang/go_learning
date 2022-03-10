package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 我们在服务器启动后就初始化一个 userDao 实例
// 把它做成全局变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

// 定义一个 UserDao 结构体
// 完成对 User 结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 1.根据一个用户id，返回一个 User实例+err
func (ud *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	// 通过给定 id 去 redis 查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// 错误！
		if err == redis.ErrNil { // 表示在 users 哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	// 这里我们需要把res反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err)
		return
	}

	return
}

// 完成登录的校验 Login
// 1. Login 完成对用户的验证
// 2. 如果用后的id和pwd都正确，则返回一个user实例
// 3. 如果用户的id或pwd有错误，则返回对应的错误信息
func (ud *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	// 先从 UserDao 的连接池中取出一根连接

	conn := ud.pool.Get()
	defer conn.Close()
	user, err = ud.getUserById(conn, userId)
	if err != nil {
		return
	}

	// 这时证明用户是获取到了.
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}

	return
}

func (ud *UserDao) Register(user *message.User) (err error) {

	// 先从 UserDao 的连接池中取出一根连接

	conn := ud.pool.Get()
	defer conn.Close()
	_, err = ud.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	// 这时说明 id 在redis还没有，则可以完成注册
	data, err := json.Marshal(user) // 序列化
	if err != nil {
		return
	}

	// 入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err: ", err)
		return
	}

	return
}
