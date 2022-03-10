package processes

import "fmt"

// 因为 UserMgr 实例在服务器端有且只有一个
// 因为在很多的地方，都会使用到，因此，
// 我们将其定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers添加
func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserId] = up
}

// 删除
func (um *UserMgr) DelOnlineUser(userId int) {
	delete(um.onlineUsers, userId)
}

// 获取当前所有在线用户
func (um *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return um.onlineUsers
}

// 根据id返回对应的值
func (um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	// 如何从map中取出一个值，带检测方式
	up, ok := um.onlineUsers[userId]
	if !ok { // 说明，要查找的这个用户，当前不在线
		err = fmt.Errorf("用户 %d 不在线", userId)
		return
	}

	return
}
