package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"strconv"
)

// GetBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	// 写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	// 执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		// var book *model.Book  错
		book := &model.Book{}
		// 给book中的字段赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		// 将book添加到books中
		books = append(books, book)
	}
	return books, nil
}

// AddBook 向数据库中添加一本图书
func AddBook(book *model.Book) error {
	// 写sql语句
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBook 根据图书的id从数据库中删除一本图书
func DeleteBook(bookID string) error {
	// 写sql语句
	sqlStr := "delete from books where id = ?"
	// 执行
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil
}

// GetBookByID 根据图书的id从数据库中查询出一本图书
func GetBookByID(bookID string) (*model.Book, error) {
	// 写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, bookID)
	// 创建Book
	book := &model.Book{}
	// 为book中的字段赋值
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

// UpdateBook 根据图书的id更新图书信息
func UpdateBook(book *model.Book) error {
	// 写sql语句
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	// 执行
	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetPageBooks 获取带分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	// 将页码转换为int64
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	// 获取数据库中图书的总记录数
	sqlStr := "select count(*) from books"
	// 设置一个变量接收总记录数
	var totalRecord int64
	// 执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	// 设置每页只显示四条记录
	var pageSize int64 = 4
	// 设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	// 获取当前页中的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	// 执行
	var books []*model.Book
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		// 将book添加到books
		books = append(books, book)
	}
	// 创建page
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

// GetPageBooksByPrice 获取带分页和价格范围的图书信息
func GetPageBooksByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	// 将页码转换为int64
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	// 获取数据库中图书的总记录数
	sqlStr := "select count(*) from books where price between ? and ?"
	// 设置一个变量接收总记录数
	var totalRecord int64
	// 执行
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	// 设置每页只显示四条记录
	var pageSize int64 = 4
	// 设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	// 获取当前页中的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	// 执行
	var books []*model.Book
	rows, err := utils.Db.Query(sqlStr2, minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		// 将book添加到books
		books = append(books, book)
	}
	// 创建page
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}
