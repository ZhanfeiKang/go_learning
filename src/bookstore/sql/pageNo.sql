-- 每页显示四条记录，第一页
SELECT * FROM books LIMIT 0,4

-- 第二页
SELECT * FROM books LIMIT 4,4

-- 第三页
SELECT * FROM books LIMIT 8,4

-- 结论: 假设当前页是pageNo，每页显示的条数是pageSize
-- limit (pageNo - 1) * pageSize, pageSize