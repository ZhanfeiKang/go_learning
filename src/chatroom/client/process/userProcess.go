package process

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
	// 暂时不需要字段...
}

func (up *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// 1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 2. 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3.创建一个 registerMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	// 4. 将 registerMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 5. 把data赋给 mes.Data字段
	mes.Data = string(data)

	// 6. 将 mes 进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 创建一个 Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	// 发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误, err: ", err)
	}

	mes, err = tf.ReadPkg() // mes 就是
	if err != nil {
		fmt.Println("readPkg err: ", err)
		return
	}

	// 将mes的Data部分反序列化成 RegisterResMes
	var resgisterResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &resgisterResMes)
	if resgisterResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录..")
		os.Exit(0)
	} else {
		fmt.Println(resgisterResMes.Error)
		os.Exit(0)
	}

	return
}

// 给关联一个用户登录的方法
// 写一个函数，完成登录
func (up *UserProcess) Login(userId int, userPwd string) (err error) {

	// 下一个就要开始定协议..
	// fmt.Printf("userId = %d, userPwd = %s\n", userId, userPwd)

	// return nil
	// 1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 2. 准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 4. 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 5. 把data赋给 mes.Data字段
	mes.Data = string(data)

	// 6. 将 mes 进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 7. 到这个时候 data 就是我们要发送的消息
	// 7.1 先把 data 的长度发送给服务器
	// 先获取到 data 的长度 -> 转成一个表示长度的byte切片
	// var pkgLen uint32
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) err: ", err)
		return
	}

	fmt.Printf("客户端，发送消息的长度: %d, \n内容: %s\n", len(data), string(data))

	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err: ", err)
		return
	}

	// time.Sleep(20 * time.Second)
	// fmt.Println("休眠了20s")

	// 这里还需要处理服务器端返回的消息。
	// 创建一个 Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg() // mes 就是
	if err != nil {
		fmt.Println("readPkg err: ", err)
		return
	}

	// 将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// fmt.Println("登录成功")

		// 可以显示当前在线用户列表，遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {

			// 如果我们要求不显示自己在线，下面我们增加一个代码
			if v == userId {
				continue
			}

			fmt.Println("用户id: \t", v)
			// 完成 客户端的 onlineUsers 的初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("------------------------------------")

		// 这里我们还需要在客户端启动一个协程
		// 该协程保持喝服务器端的通讯.如果服务器有数据推送给客户端
		// 则接收并显示在客户端的终端.
		go serverProcessMes(conn)

		// 1.显示我们的登录成功的菜单[循环]..
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
