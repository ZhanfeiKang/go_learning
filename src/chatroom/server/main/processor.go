package main

import (
	"chatroom/common/message"
	"chatroom/server/processes"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

// 先创建一个 Processor 的结构体
type Processor struct {
	Conn net.Conn
}

// 编写一个ServerProcessMes 函数
// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (processor *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		// 创建一个UserProcess 实例
		up := &processes.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
		up := &processes.UserProcess{
			Conn: processor.Conn,
		}
		err = up.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (processor *Processor) process() (err error) {
	// 循环的读取客户端发送的信息
	for {

		// 这里我们将读取数据包直接封装成一个函数readPkg(), 返回Message, Err
		// 创建一个Transfer实例完成读包的任务
		tf := &utils.Transfer{
			Conn: processor.Conn,
		}
		mes, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出..")
				return err
			} else {
				fmt.Println("readPkg err: ", err)
				return err
			}
		}

		// fmt.Println("mes: ", mes)
		err = processor.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
