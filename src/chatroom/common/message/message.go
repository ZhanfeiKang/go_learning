package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

// 这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string `json:"type"` // 消息的类型
	Data string `json:"data"` // 消息的类型
}

// 定义两个消息..后面需要再增加

type LoginMes struct {
	UserId   int    `json:"userId"`   // 用户id
	UserPwd  string `json:"userPwd"`  // 用户密码
	UserName string `json:"userName"` // 用户名
}

type LoginResMes struct {
	Code int `json:"code"` // 返回状态码
	// 500 表示用户未注册
	// 200 表示登陆成功
	UsersId []int  // 增加字段，保存用户id的切片
	Error   string `json:"error"` // 返回错误信息
}

type RegisterMes struct {
	User User `json:"user"` // 类型就是User结构体.
}

type RegisterResMes struct {
	Code int `json:"code"` // 返回状态码
	// 400 表示用户已经占用
	// 200 表示注册成功
	Error string `json:"error"` // 返回错误信息
}

// 为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}