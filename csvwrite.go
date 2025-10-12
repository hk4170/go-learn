package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// 1. 创建或打开 CSV 文件（追加模式）
	file, err := os.Create("data.csv") // 或 os.OpenFile("output.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("无法创建文件:", err)
	}
	defer file.Close()

	// 2. 创建 CSV Writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // 确保数据写入磁盘

	// 3. 写入单行数据（[]string）
	err = writer.Write([]string{"Name", "Age", "City"})
	if err != nil {
		log.Fatal("写入标题失败:", err)
	}

	// 4. 写入多行数据
	data := [][]string{
		{"Alice", "25", "New York"},
		{"Bob", "30", "London"},
		{"Charlie", "35", "Tokyo"},
	}
	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			log.Fatal("写入行失败:", err)
		}
	}
}
