package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // 导入但不直接使用，仅用于注册驱动
)

func main() {
	// 设置数据库连接信息
	// 格式：用户名:密码@tcp(主机:端口)/数据库名?charset=utf8mb4&parseTime=True
	dsn := "root:@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close() // 确保程序退出前关闭连接

	// 测试连接是否有效（可选但推荐）
	err = db.Ping()
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
	fmt.Println("✅ 成功连接到 MySQL 数据库！")

	// 示例：查询数据
	rows, err := db.Query("SELECT id, name FROM users WHERE age > ?", 18)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
	}
	defer rows.Close()

	// 遍历查询结果
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatalf("读取行数据失败: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		log.Fatalf("遍历结果出错: %v", err)
	}
}
