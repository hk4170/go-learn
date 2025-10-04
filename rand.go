package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
)

func main() {
	// 方法1：生成一个 [0, 100) 的随机整数
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}
	fmt.Println("crypto/rand 随机整数 (0~99):", n)

	// 方法2：生成随机字节，再转成整数（更底层）
	var buf [4]byte
	_, err = rand.Read(buf[:])
	if err != nil {
		panic(err)
	}
	num := binary.BigEndian.Uint32(buf[:])
	fmt.Println("crypto/rand 随机 uint32:", num)
}
