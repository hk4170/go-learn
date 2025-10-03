package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

// Person 结构体（必须与服务器端一致）
type Person struct {
	Name string
	Age  int
}

func main() {
	// 连接到服务端
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// 创建要发送的 Person 对象
	person := Person{Name: "Alice", Age: 30}

	// 创建 gob 编码器，将数据编码后发送到连接
	encoder := gob.NewEncoder(conn)

	err = encoder.Encode(person)
	if err != nil {
		log.Fatalf("Encode error: %v", err)
	}

	fmt.Println("Sent person:", person)
}
