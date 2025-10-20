package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. 打开 CSV 文件
	file, err := os.Open("data.csv") // 替换为你的 CSV 文件路径
	if err != nil {
		log.Fatal("无法打开文件:", err)
	}
	defer file.Close() // 确保文件关闭

	// 2. 创建 CSV Reader
	reader := csv.NewReader(file)

	// 3. 读取所有记录（返回 [][]string）
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("读取 CSV 失败:", err)
	}
	println("header")
	header := records[0]
    fmt.Println(header)
	
	println()
	// 4. 遍历并打印每一行
	//go无法删除切片内容 只能通过切片操作或者重新创建一个切片
	data := records[1:] //可通过切片操作删除第一行
	for _, v := range data {
		fmt.Println(v) // 每行是一个 []string
	}
}
