package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    // 读取整个文件为 []byte
    data, err := os.ReadFile("test.bin") 
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("读取到的二进制数据（十六进制）: %x\n", data)
}
