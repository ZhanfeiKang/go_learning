package processes

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	// ..[暂时不需要字段]
}

// 写方法转发消息
func (sp *SmsProcess) SendGroupMes(mes *message.Message) {

	// 遍历服务器端的 onlineUsers map[int]*UserProcess
	// 将消息转发出去

	// 取出 mes 的内容 SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err: ", err.Error())
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err: ", err.Error())
		return
	}

	for id, up := range userMgr.onlineUsers {
		// 这里我们还需要过滤掉自己，即不要再发给自己
		if id == smsMes.UserId {
			continue
		}
		sp.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (sp *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	// 创建一个Transfer实例，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err: ", err)
	}
}
