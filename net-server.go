package main
import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 启动 TCP 服务，监听 8080 端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()
	fmt.Println("Server is listening on :8080")

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}

		// 处理每个客户端连接（可以用 goroutine）
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 创建 gob 解码器，从连接中读取并解码数据
	decoder := gob.NewDecoder(conn)

	var p Person
	err := decoder.Decode(&p)
	if err != nil {
		log.Printf("Decode error: %v", err)
		return
	}

	// 打印接收到的数据
	fmt.Printf("Received from client: %+v\n", p)
}

