package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// AddOrder 向数据库插入订单
func AddOrder(order *model.Order) error {
	// 写sql语句
	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}

// GetOrders 获取数据库中所有的订单
func GetOrders() ([]*model.Order, error) {
	// 写sql语句
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders"
	// 执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}

// GetMyOrders 获取我的订单
func GetMyOrders(userID int) ([]*model.Order, error) {
	// 写sql语句
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id = ?"
	// 执行
	rows, err := utils.Db.Query(sqlStr, userID)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}

// UpdateOrderState 更新订单状态，即发货和收货
func UpdateOrderState(orderID string, state int64) error {
	// 写sql语句
	sqlStr := "update orders set state = ? where id = ?"
	// 执行
	_, err := utils.Db.Exec(sqlStr, state, orderID)
	if err != nil {
		return err
	}
	return nil
}
