package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("测试bookdao中的方法")
	m.Run()
}

func TestUser(t *testing.T) {
	// fmt.Println("测试userdao中的函数")
	// t.Run("验证用户名或密码：", testLogin)
	// t.Run("验证用户名：", testRegist)
	// t.Run("添加用户：", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取的用户信息是：", user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取的用户信息是：", user)
}

func testSave(t *testing.T) {
	SaveUser("admin3", "123456", "admin@kkite.com")
}

func TestBook(t *testing.T) {
	// fmt.Println("测试bookdao中的相关函数")
	// t.Run("测试获取所有图书", testGetBooks)
	// t.Run("测试添加图书", testAddBook)
	// t.Run("测试删除图书", testDeleteBook)
	// t.Run("测试获取一本图书", testGetBook)
	// t.Run("测试更新图书", testUpdateBook)
	// t.Run("测试带分页的图书", testGetPageBooks)
	// t.Run("测试带分页和价格范围的图书", testGetPageBooksByPrice)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	// 遍历得到每一本书
	for k, v := range books {
		fmt.Printf("第%v本书的信息是：%v\n", k+1, v)
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.8,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	// 调用添加图书的函数
	AddBook(book)
}

func testDeleteBook(t *testing.T) {
	// 调用删除图书的函数
	DeleteBook("33")
}

func testGetBook(t *testing.T) {
	// 调用删除图书的函数
	book, _ := GetBookByID("31")
	fmt.Println("获取的图书信息是：", book)
}

func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      31,
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   66.66,
		Sales:   10000,
		Stock:   10,
		ImgPath: "/static/img/default.jpg",
	}
	// 调用更新图书的函授
	UpdateBook(book)
}

func testGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("9")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.TotalRecord)
	fmt.Println("当前页中的图书有：")
	for _, v := range page.Books {
		fmt.Println("图书有：", v)
	}
}

func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("3", "10", "30")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.TotalRecord)
	fmt.Println("当前页中的图书有：")
	for _, v := range page.Books {
		fmt.Println("图书有：", v)
	}
}

func TestSession(t *testing.T) {
	// fmt.Println("测试Session相关函数")
	// t.Run("测试添加Session", testAddSession)
	// t.Run("测试删除Session", testDeleteSession)
	// t.Run("测试获取Session", testGetSession)
}

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "13838381438",
		UserName:  "admin4_zs",
		UserID:    4,
	}
	AddSession(sess)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("13838381438")
}

func testGetSession(t *testing.T) {
	sess, _ := GetSession("480735d8-ac77-4e0d-4b06-09060f1d31f8")
	fmt.Println("Session的信息是：", sess)
}

func TestCart(t *testing.T) {
	// fmt.Println("测试购物车的相关函数")
	// t.Run("测试添加购物车", testAddCart)
	// t.Run("测试根据图书的id获取对应的购物项", testGetCartItemByBookID)
	// t.Run("测试根据购物车的id获取所有的购物项", testGetCartItemsByCartID)
	// t.Run("测试根据用户的id获取购物车", testGetCartByUserID)
	// t.Run("测试根据图书的id和购物车的id以及输入的图书的数量更新购物项", testUpdateBookCount)
	// t.Run("测试根据购物车的id删除购物车", testDeleteCartByCartID)
	// t.Run("测试删除购物项", testDeleteCartItemByID)
}

func testAddCart(t *testing.T) {
	// 设置要买的书
	book1 := &model.Book{
		ID:    1,
		Price: 27.20,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}
	// 创建一个购物项切片
	var cartItems []*model.CartItem
	// 创建两个购物项
	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  10,
		CartID: "66668888",
	}
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "66668888",
	}
	cartItems = append(cartItems, cartItem1)
	cartItems = append(cartItems, cartItem2)
	// 创建购物车
	cart := &model.Cart{
		CartID:    "66668888",
		CartItems: cartItems,
		UserID:    1,
	}
	// 将购物车插入到数据库中
	AddCart(cart)
}

func testGetCartItemByBookID(t *testing.T) {
	cartItem, _ := GetCartItemByBookIDAndCartID("1", "66668888")
	fmt.Println("图书id=1的购物项的信息是: ", cartItem)
}

func testGetCartItemsByCartID(t *testing.T) {
	cartItems, _ := GetCartItemsByCartID("66668888")
	for k, v := range cartItems {
		fmt.Printf("第%v个购物项是: %v\n", k+1, v)
	}
}

func testGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(1)
	fmt.Println("id为1的用户的购物车信息是：", cart)
}

func testUpdateBookCount(t *testing.T) {
	// UpdateBookCount(cartItem)
}

func testDeleteCartByCartID(t *testing.T) {
	DeleteCartByCartID("231fb9b5-8890-4e3b-5bdc-d885030fbcac")
}

func testDeleteCartItemByID(t *testing.T) {
	DeleteCartItemByID("27")
}

func TestOrder(t *testing.T) {
	fmt.Println("测试订单相关函数")
	// t.Run("测试添加订单和订单项", testAddOrder)
	// time := time.Now()
	// lo := time.Local()
	// fmt.Println("现在的时间是: ", time)
	// fmt.Println("本地的时间是: ", lo)

	// t.Run("测试获取所有的订单", testGetOrders)
	// t.Run("测试获取所有的订单项", testGetOrderItems)
	// t.Run("测试获取我的订单", testGetMyOrders)
	t.Run("测试发货和收货", testUpdateOrderState)
}

func testAddOrder(t *testing.T) {
	// 生成订单号
	orderID := "24424242424242"
	// 创建订单
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  2,
		TotalAmount: 400,
		State:       0,
		UserID:      1,
	}
	// 创建订单项
	orderItem1 := &model.OrderItem{
		Count:   1,
		Amount:  300,
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   300,
		ImgPath: "/static/img/default.jpg",
		OrderId: orderID,
	}
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "西游记",
		Author:  "吴承恩",
		Price:   100,
		ImgPath: "/static/img/default.jpg",
		OrderId: orderID,
	}
	// 保存订单
	AddOrder(order)
	// 保存订单项
	AddOrderItem(orderItem1)
	AddOrderItem(orderItem2)
}

func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, v := range orders {
		fmt.Println("订单信息是: ", v)
	}
}

func testGetOrderItems(t *testing.T) {
	orderItems, _ := GetOrderItemsByOrderID("6d1c01df-9b69-4f3f-6a04-b8f2dfb7f20e")
	for _, v := range orderItems {
		fmt.Println("订单项信息是: ", v)
	}
}

func testGetMyOrders(t *testing.T) {
	orders, _ := GetMyOrders(1)
	for _, v := range orders {
		fmt.Println("我的订单有：", v)
	}
}

func testUpdateOrderState(t *testing.T) {
	UpdateOrderState("6d1c01df-9b69-4f3f-6a04-b8f2dfb7f20e", 1)
}
