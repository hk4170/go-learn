package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // 导入 PostgreSQL 驱动（仅注册）
)

func main() {
	// PostgreSQL 连接字符串 DSN
	// 格式：host=127.0.0.1 port=5432 user=postgres password=123456 dbname=testdb sslmode=disable
	dsn := "host=localhost port=5432 user=postgres password=123456 dbname=testdb sslmode=disable"

	// 打开数据库连接
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	defer db.Close()

	// 测试连接是否有效
	err = db.Ping()
	if err != nil {
		log.Fatal("无法连接到 PostgreSQL:", err)
	}
	fmt.Println("✅ 成功连接到 PostgreSQL！")

	// 创建表（如果不存在）
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			age INTEGER
		)
	`)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}
	fmt.Println("✅ 表已创建或已存在")

	// 插入数据
	result, err := db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", "Alice", 25)
	if err != nil {
		log.Fatal("插入失败:", err)
	}
	rowsAffected, _ := result.RowsAffected()
	//lastInsertID := -1
	// PostgreSQL 不直接支持 LastInsertId()，但可通过 RETURNING 或查询 currval 获取
	// 这里简单打印影响了多少行
	fmt.Printf("✅ 插入成功，影响行数: %d\n", rowsAffected)

	// 查询数据
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal("查询失败:", err)
	}
	defer rows.Close()

	fmt.Println("📋 用户列表：")
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("读取行失败:", err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// 检查遍历错误
	if err = rows.Err(); err != nil {
		log.Fatal("遍历出错:", err)
	}
}
