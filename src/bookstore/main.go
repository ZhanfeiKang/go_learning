package main

import (
	"bookstore/controler"
	"net/http"
)

func main() {
	// 设置处理静态资源
	// /static/会匹配 以 /static/ 开头的路径，当浏览器请求index.html页面中的style.css文件时，
	// static前缀会被替换为views/static。然后去 views/static/css目录中去查找style.css文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	// 直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/main", controler.GetPageBooksByPrice)

	// 去登录
	http.HandleFunc("/login", controler.Login)
	// 去注销
	http.HandleFunc("/logout", controler.Logout)
	// 去注册
	http.HandleFunc("/regist", controler.Regist)

	// 通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controler.CheckUserName)

	// 获取带分页的图书信息
	http.HandleFunc("/getPageBooks", controler.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controler.GetPageBooksByPrice)
	// 添加图书
	http.HandleFunc("/addBook", controler.UpdateOrAddBook)
	// 删除图书
	http.HandleFunc("/deleteBook", controler.DeleteBook)
	// 去更新图书的页面
	http.HandleFunc("/toUpdateBookPage", controler.ToUpdateBookPage)
	// 更新图书
	http.HandleFunc("/updateOrAddBook", controler.UpdateOrAddBook)

	// 添加图书到购物车
	http.HandleFunc("/addBook2Cart", controler.AddBook2Cart)
	// 获取购物车信息
	http.HandleFunc("/getCartInfo", controler.GetCartInfo)
	// 清空购物车
	http.HandleFunc("/deleteCart", controler.DeleteCart)
	// 删除购物项
	http.HandleFunc("/deleteCartItem", controler.DeleteCartItem)
	// 更新购物项
	http.HandleFunc("/updateCartItem", controler.UpdateCartItem)

	// 去结账
	http.HandleFunc("/checkout", controler.Checkout)

	// 获取所有订单
	http.HandleFunc("/getOrders", controler.GetOrders)
	// 获取订单详情 ， 即订单所对应的所有订单项
	http.HandleFunc("/getOrderInfo", controler.GetOrderInfo)
	// 获取我的订单
	http.HandleFunc("/getMyOrder", controler.GetMyOrders)
	// 发货
	http.HandleFunc("/sendOrder", controler.SendOrder)
	// 确认收货
	http.HandleFunc("/takeOrder", controler.TakeOrder)

	http.ListenAndServe(":8080", nil)
}
