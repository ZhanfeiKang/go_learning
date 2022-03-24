package model

//  OrderItem 结构
type OrderItem struct {
	OrderItemID int64   // 订单项的id
	Count       int64   // 订单项中图书的数量
	Amount      float64 // 订单项中图书的金额小计
	Title       string  // 订单项中图书的书名
	Author      string  // 订单项中图书的作者
	Price       float64 // 订单项中图书的价格
	ImgPath     string  // 订单中图书的封面
	OrderId     string  // 订单项所属于的订单
}
