package main

import (
    "os"
)

func main() {
    // 定义变量
    name := "write test"
    
    // 写入文件（文件不存在则创建，存在则覆盖）
    err := os.WriteFile("file.txt",[]byte(name), 0644)
    if err != nil {
        panic(err)
    }
}