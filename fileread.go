package main

import (
    "os"
    "log"
    "fmt"
)

func main() {
    // 读取文件全部字节
    content, err := os.ReadFile("file.txt")
    if err != nil {
        log.Fatalf("读取文件失败: %v", err)
    }
    // 转换为字符串并打印
    fmt.Println("文件内容：")
    fmt.Println(string(content))
}