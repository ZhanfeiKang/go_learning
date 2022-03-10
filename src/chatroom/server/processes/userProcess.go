package processes

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	// 字段？
	Conn net.Conn
	// 增加一个字段，表示该conn是哪个用户
	UserId int
}

// 这里我们编写通知所有在线的用户的方法
// userId int 要通知其他在线用户，我上线
func (up *UserProcess) NotifyOthersOnlineUser(userId int) {

	// 遍历 onlineUsers，然后一个一个的发送 NotifyUserStatusMes
	for id, upOther := range userMgr.onlineUsers {
		// 过滤掉自己
		if id == userId {
			continue
		}
		// 开始通知【单独的写一个方法】
		upOther.NotifyMeOnline(userId)
	}
}

// 通知 userId 已上线
func (up *UserProcess) NotifyMeOnline(userId int) {
	// 组装我们的消息NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	// 将 notifyUserStatusMes 序列化
	data, err := json.Marshal(&notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}
	// 将序列化后的 notifyUserStatusMes 赋值给 mes.Data
	mes.Data = string(data)

	// 对 mes 序列化, 准备发送
	data, err = json.Marshal(&mes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 发送，创建 Transfer 实例，发送
	tf := &utils.Transfer{
		Conn: up.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline WritePkg err: ", err)
		return
	}
}

func (up *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	// 1.先从mes中取出mes.Data, 并直接反序列化成registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Umarshal fail err: ", err)
		return
	}

	// 1.先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	// 2.再声明一个 registerResMes, 并完成赋值
	var registerResMes message.RegisterResMes

	// 我们需要到redis数据库去完成注册
	// 1. 使用model.MyUserDao 到 redis 去验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
	}

	// 3.将 registerResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 4.将data赋值给 resMes
	resMes.Data = string(data)

	// 5.对 resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 6.发送 data， 我们将其封装到writePkg函数
	// 因为使用分层模式（MVC），我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)

	return
}

// 编写一个函数 serverProcessLogin 函数，专门处理登录请求
func (up *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 核心代码..
	// 1.先从mes中取出mes.Data, 并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Umarshal fail err: ", err)
		return
	}

	// 1.先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 2.再声明一个LoginResMes, 并完成赋值
	var loginResMes message.LoginResMes

	// 我们需要到redis数据库去完成验证
	// 1. 使用model.MyUserDao 到 redis 去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}

	} else {
		loginResMes.Code = 200
		// 这里，因为用户登录成功，我们就把该登录成功的user放入到 userMgr 中
		// 将登录成功的用后的userId 赋给 up
		up.UserId = loginMes.UserId
		userMgr.AddOnlineUser(up)
		// 通知其他在线的用户，我上线了
		up.NotifyOthersOnlineUser(loginMes.UserId)
		// 将当前在线用户的id放入到 loginResdMes.UserId
		// 遍历 userMgr.onlineUsers
		for id := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user.UserName, "登录成功")
	}

	// fmt.Println(loginResMes.Error)

	// // 如果用户的id为100，密码等于123456，认为合法，否则不合法
	// if LoginMes.UserId == 100 && LoginMes.UserPwd == "123456" {
	// 	// 合法
	// 	LoginResMes.Code = 200

	// } else {
	// 	// 不合法
	// 	LoginResMes.Code = 500 // 500状态码表示该用户不存在
	// 	LoginResMes.Error = "该用户不存在，请注册再使用..."
	// }

	// 3.将 LoginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 4.将data赋值给 resMes
	resMes.Data = string(data)

	// 5.对 resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	// 6.发送 data， 我们将其封装到writePkg函数
	// 因为使用分层模式（MVC），我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return
}
