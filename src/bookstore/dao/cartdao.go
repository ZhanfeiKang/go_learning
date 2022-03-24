package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

// AddCart 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	// 写sql语句
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	// 执行sql
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	// 获取购物车中的所有购物项
	cartItems := cart.CartItems
	// 遍历得到每一个购物项
	for _, cartItem := range cartItems {
		// 将购物项插入到数据库中
		AddCartItem(cartItem)
	}

	return nil
}

// GetCartByUserID 根据用户的id从数据库中查询对应的购物车
func GetCartByUserID(userID int) (*model.Cart, error) {
	// 写sql
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id = ?"
	// 执行sql
	row := utils.Db.QueryRow(sqlStr, userID)
	// 创建一个cart
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	// 获取当前购物车中的所有购物项
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	// 将所有的购物项设置到购物车中
	cart.CartItems = cartItems
	return cart, nil
}

// UpdateCart 更新购物车中图书的总数量和总金额
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count = ?, total_amount = ? where id = ?"
	// 执行
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartByCartID 删除购物车
func DeleteCartByCartID(cartID string) error {
	// 删除购物车之前需要先删除下面所有的购物项
	err := DeleteCartItemsByCartID(cartID)
	if err != nil {
		return err
	}
	// 写sql语句
	sqlStr := "delete from carts where id = ?"
	_, err = utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}
