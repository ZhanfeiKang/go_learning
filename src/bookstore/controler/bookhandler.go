package controler

import (
	"bookstore/dao"
	"bookstore/model"
	"net/http"
	"strconv"
	"text/template"
)

// IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// 获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	// 调用bookdao中获取带分页图书的函数
	page, _ := dao.GetPageBooks(pageNo)
	// 解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	// 执行
	t.Execute(w, page)
}

// GetPageBooks获取所有图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	// 获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	// 调用bookdao中获取带分页图书的函数
	page, _ := dao.GetPageBooks(pageNo)
	// 解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	// 执行
	t.Execute(w, page)
}

// GetPageBooksByPrice获取带分页和价格范围的图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	// 获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}

	// 获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		// 调用bookdao中获取带分页图书的函数
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		// 调用bookdao中获取带分页图书的函数
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		// 将价格范围设置到page中
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}

	// 获取Cookie
	// 调用 IsLogin 函数判断是否已经登录
	flag, session := dao.IsLogin(r)

	if flag {
		// 已经登录, 设置page中的IsLogin字段和Username的字段值
		page.IsLogin = true
		page.Username = session.UserName
	}

	// 解析模板文件
	t := template.Must(template.ParseFiles("views/index.html"))
	// 执行
	t.Execute(w, page)
}

// GetBooks获取所有图书
// func GetBooks(w http.ResponseWriter, r *http.Request) {
// 	// 调用bookdao中获取所有图书的函数
// 	books, _ := dao.GetBooks()
// 	// 解析模板文件
// 	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
// 	// 执行
// 	t.Execute(w, books)
// }

// // AddBook 添加图书
// func AddBook(w http.ResponseWriter, r *http.Request) {
// 	// 获取图书信息
// 	title := r.PostFormValue("title")
// 	author := r.PostFormValue("author")
// 	price := r.PostFormValue("price")
// 	sales := r.PostFormValue("sales")
// 	stock := r.PostFormValue("stock")
// 	// 将price、sales、stock进行转换
// 	fPrice, _ := strconv.ParseFloat(price, 64)
// 	iSales, _ := strconv.ParseInt(sales, 10, 0)
// 	iStock, _ := strconv.ParseInt(stock, 10, 0)
// 	// 创建Book
// 	book := &model.Book{
// 		Title:   title,
// 		Author:  author,
// 		Price:   fPrice,
// 		Sales:   int(iSales),
// 		Stock:   int(iStock),
// 		ImgPath: "/static/img/default.jpg",
// 	}
// 	// 调用bookdao中添加图书的函数
// 	dao.AddBook(book)
// 	// 调用GetBook处理器函数再查一遍数据库
// 	GetBooks(w, r)
// }

// 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// 获取言删除的图书ID
	bookID := r.FormValue("bookId")
	// 调用bookdao中的删除图书的函数
	dao.DeleteBook(bookID)
	// 调用GetBook处理器函数再查一遍数据库
	GetPageBooks(w, r)
}

// ToUpdateBookPage 去更新或者添加图书的页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	// 获取要更新的图书ID
	bookID := r.FormValue("bookId")
	// 调用bookdao中获取图书的函数
	book, _ := dao.GetBookByID(bookID)
	if book.ID > 0 {
		// 在更新图书
		// 解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		// 执行
		t.Execute(w, book)
	} else {
		// 在添加图书
		// 解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		// 执行
		t.Execute(w, "")
	}

}

// UpdateOrAddBook 更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	// 获取图书信息
	bookID := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	// 将price、sales、stock进行转换
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	ibookID, _ := strconv.ParseInt(bookID, 10, 0)

	// 创建Book
	book := &model.Book{
		ID:      int(ibookID),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	if book.ID > 0 {
		// 在更新图书
		// 调用bookdao中更新图书的函数
		dao.UpdateBook(book)
	} else {
		// 在添加图书
		// 调用bookdao中添加图书的函数
		dao.AddBook(book)
	}

	// 调用GetBook处理器函数再查一遍数据库
	GetPageBooks(w, r)
}
