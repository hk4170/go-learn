package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func generateHex(length int) (string, error) {
	// 创建一个指定长度的字节切片
	b := make([]byte, length)
	// 从 crypto/rand 中读取安全的随机字节
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// 转为 16 进制字符串
	return hex.EncodeToString(b), nil
}

func main() {
	token, err := generateHex(6) 
	if err != nil {
		panic(err)
	}
	fmt.Println("安全随机 Token (Hex):", token)
}
