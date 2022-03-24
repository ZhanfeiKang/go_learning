-- 创建订单表
CREATE TABLE orders(
	id VARCHAR(100) PRIMARY KEY,
	create_time DATETIME NOT NULL,
	total_count INT NOT NULL,
	total_amount DOUBLE(11,2) NOT NULL,
	state INT NOT NULL,
	user_id INT,
	FOREIGN KEY(user_id) REFERENCES users(id)
)

-- 创建订单项表
CREATE TABLE order_items(
	id INT PRIMARY KEY AUTO_INCREMENT,
	`count` INT NOT NULL,
	amount DOUBLE(11,2) NOT NULL,
	title VARCHAR(100) NOT NULL,
	author VARCHAR(100) NOT NULL,
	price DOUBLE(11,2) NOT NULL,
	img_path VARCHAR(100) NOT NULL,
	order_id VARCHAR(100) NOT NULL,
	FOREIGN KEY(order_id) REFERENCES orders(id)
)