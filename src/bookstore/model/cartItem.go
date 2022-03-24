package model

// CartItem 购物项
type CartItem struct {
	CartItemID int64   // 购物项的id
	Book       *Book   // 图书信息
	Count      int64   // 图书数量
	Amount     float64 // 图书金额小计，通过计算得到
	CartID     string  // 当前购物项术语哪一个购物车
}

// GetAmount 获取购物项中图书的金额小计，由图书的价格和数量计算得到
func (cartItem *CartItem) GetAmount() float64 {
	// 获取当前购物项中图书的价格
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}
