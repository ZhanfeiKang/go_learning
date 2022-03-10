package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 这里将这些方法关联到结构体中
type Transfer struct {
	// 分析它应该有哪些字段
	Conn net.Conn
	Buf  [4096]byte // 这是传输时，使用的 缓冲
}

func (tf *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 4096)
	fmt.Println("读取客户端发送的数据...")
	// conn.Read 在 conn 没有被关闭的情况下，才会阻塞
	// 如果客户端关闭了 conn 则，不会阻塞
	_, err = tf.Conn.Read(tf.Buf[:4])
	if err != nil {
		// fmt.Println("conn.Read err: ", err)
		// err = errors.New("read pkg header error")
		return
	}

	// 根据buf[:4] 转成一个uint32类型
	// var pkgLen uint32
	pkgLen := binary.BigEndian.Uint32(tf.Buf[0:4])

	// 根据pkgLen 读取消息内容
	n, err := tf.Conn.Read(tf.Buf[:pkgLen]) // 意思是从 conn 套接字中读 pkgLen 个字节扔到 buf 中去
	if n != int(pkgLen) || err != nil {
		// fmt.Println("conn.Read fail err: ", err)
		// err = errors.New("read pkg body error")
		return
	}

	// 把pkg反序列化成 -> message.Message
	// 注意：&mes !!!!!!
	err = json.Unmarshal(tf.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err: ", err)
		return
	}

	return
}

func (tf *Transfer) WritePkg(data []byte) (err error) {

	// 先发送一个长度给对方
	pkgLen := uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(tf.Buf[0:4], pkgLen)
	// 发送长度
	n, err := tf.Conn.Write(tf.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) err: ", err)
		return
	}

	// 发送data本身
	n, err = tf.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(buf) err: ", err)
		return
	}
	return
}
