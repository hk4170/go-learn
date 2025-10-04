package main

import (
	"database/sql"
	"fmt"
	"log"
    _"github.com/mattn/go-sqlite3" // 导入驱动，仅用于注册
)

func main() {
	// SQLite 不需要服务器，直接打开一个数据库文件
	// 如果文件不存在，会自动创建
	db, err := sql.Open("sqlite3", "./test.db") // 当前目录下 test.db 文件
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}
	defer db.Close() // 确保关闭

	// 测试连接（可选）
	err = db.Ping()
	//println("test:",err.Error())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✅ 成功连接到 SQLite 数据库！")

	// 创建表（如果不存在）
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER
		)
	`)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}
	fmt.Println("✅ 表已创建或已存在")

	// 插入数据
	result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)
	if err != nil {
		log.Fatal("插入失败:", err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("✅ 插入成功，ID: %d\n", id)

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

