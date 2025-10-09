// client.go
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

// 注意：客户端的请求/响应结构体，必须与服务端 **完全一致**（字段名、类型、首字母大写）
type GetUserRequest struct {
	UserID int64
}

type GetUserResponse struct {
	ID       int64
	Name     string
	Age      int
	CreateAt time.Time
}

func main() {
	// 1. 连接 RPC 服务端（指定 TCP 地址，Gob 自动生效）
	client, err := rpc.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("连接服务失败：%v", err)
	}
	defer client.Close()
	log.Println("成功连接 RPC 服务")

	// 2. 构造请求参数（查询 UserID=1001 的用户）
	req := GetUserRequest{UserID: 1001}
	var res GetUserResponse // 接收响应结果

	// 3. 调用服务端函数：格式为 "服务名.函数名"
	err = client.Call("UserService.GetUser", req, &res)
	if err != nil {
		log.Fatalf("调用 GetUser 失败：%v", err)
	}

	// 4. 打印响应结果（Gob 已自动反序列化 time.Time 等类型）
	fmt.Println("=== 查询用户信息 ===")
	fmt.Printf("用户ID：%d\n", res.ID)
	fmt.Printf("用户名：%s\n", res.Name)
	fmt.Printf("年龄：%d\n", res.Age)
	fmt.Printf("创建时间：%s\n", res.CreateAt.Format("2006-01-02 15:04:05"))
}