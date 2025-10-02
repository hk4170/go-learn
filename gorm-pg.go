package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 定义一个 User 模型，对应数据库中的表
type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// PostgreSQL 连接字符串（DSN）
	// 格式：host= port= user= password= dbname= sslmode=
	dsn := "host=localhost user=postgres password= dbname=test port=5432 sslmode=disable"
	
	// 连接到 PostgreSQL 数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	fmt.Println("✅ 成功连接到数据库")

	// 自动迁移：根据 User 结构体创建 users 表（如果不存在）
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("自动迁移失败: %v", err)
	}
	fmt.Println("✅ 数据库迁移完成")

	// 创建一条新记录
	user := User{Name: "Alice", Email: "alice@example.com"}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatalf("插入数据失败: %v", result.Error)
	}
	fmt.Printf("✅ 插入成功，ID: %d\n", user.ID)
	data := db.Table("public.users").Find("name = ?","Alice")
	fmt.Printf("data: %v\n", data)
	println(db.Name())
}
