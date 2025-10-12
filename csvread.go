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
    fmt.Println(records[0])
	println(len(records))
	// 4. 遍历并打印每一行
	for _, record := range records {
		fmt.Println(record) // 每行是一个 []string
	}
}
