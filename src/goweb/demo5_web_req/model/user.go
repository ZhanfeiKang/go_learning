package model

import (
	"fmt"
	"goweb/demo4_db/utils"
)

// User结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// AddUser 添加User的方法 	方法1
func (user *User) AddUser() (err error) {
	// 1.写 sql 语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	// 2.预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常 err", err.Error())
		return
	}
	// 3.执行
	_, err = inStmt.Exec("admin", "123456", "admin@kkite.com")
	if err != nil {
		fmt.Println("执行出现异常 err", err.Error())
		return
	}

	return
}

// AddUser2 添加User的方法 	方法2
func (user *User) AddUser2() (err error) {
	// 1.写 sql 语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	// 2.执行
	_, err = utils.Db.Exec(sqlStr, "admin2", "666666", "admin2@tom.com")
	if err != nil {
		fmt.Println("执行出现异常 err", err.Error())
		return
	}

	return
}

// GetUerById 根据用户id从数据库中查询一条记录
func (user *User) GetUerById() (*User, error) {
	// 写 sql 语句
	sqlStr := "select id,username,password,email from users where id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, user.ID)
	// 声明
	var (
		id       int
		username string
		password string
		email    string
	)

	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return u, nil
}

// GetUsers 获取数据库中所有的记录
func (user *User) GetUsers() ([]*User, error) {
	// 写SQL语句
	sqlStr := "select id,username,password,email from users"
	// 执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	// 创建User切片
	var users []*User
	// 声明
	var (
		id       int
		username string
		password string
		email    string
	)

	for rows.Next() {
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, u)
	}
	return users, nil
}
