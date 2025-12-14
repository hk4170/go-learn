package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
	"os"
	"log"
)

type ImageProject struct {
	ID          string
	Width, Height int
	PixelFormat  string
	ImageData   []byte
	CreateTime  time.Time
}

func main() {
	// --- 先执行序列化，得到 byteData（复用之前的逻辑）---
	project := ImageProject{
		ID:          "img_001",
		Width:       1920,
		Height:      1080,
		PixelFormat: "RGBA",
		ImageData:   []byte("模拟图片二进制数据"),
		CreateTime:  time.Now(),
	}

	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(project); err != nil {
		fmt.Println("编码失败:", err)
		return
	}
	byteData := buf.Bytes()
    err := os.WriteFile("gob-byte.bin", byteData,0644)
	if err != nil {
		log.Fatal(err)
	}
	
	// --- 重点：反序列化还原 struct ---
	var restoredProject ImageProject // 声明目标 struct 实例
	// 1. 创建 bytes.Reader 读取 byteData
	reader := bytes.NewReader(byteData)
	// 2. 创建 gob 解码器
	decoder := gob.NewDecoder(reader)
	// 3. 解码到目标 struct 实例
	if err := decoder.Decode(&restoredProject); err != nil {
		fmt.Println("解码失败:", err)
		return
	}

	// 验证结果
	fmt.Printf("byteData", byteData)
	println()
	fmt.Printf("还原后的项目 ID: %s\n", restoredProject.ID)
	fmt.Printf("图片尺寸: %dx%d\n", restoredProject.Width, restoredProject.Height)
	fmt.Printf("图片数据: %s\n", restoredProject.ImageData)
}