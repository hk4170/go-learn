// server.go
package main

import (
	"net"
	"net/rpc"
	"time"
	"log"
	"fmt"
)

// 1. 定义 RPC 服务结构体（承载业务函数）
type UserService struct{}

// 2. 定义请求/响应结构体（Gob 要求字段首字母大写，否则无法序列化）
// 需求：根据用户ID查询用户信息
type GetUserRequest struct {
	UserID int64 // 请求参数：用户ID
}

type GetUserResponse struct {
	ID       int64     // 响应参数：用户ID
	Name     string    // 用户名
	Age      int       // 年龄
	CreateAt time.Time // 创建时间（Gob 天然支持 time.Time 等Go原生类型）
}

// 3. 实现 RPC 业务函数（必须满足：首字母大写、2个参数、1个error返回）
func (s *UserService) GetUser(req GetUserRequest, res *GetUserResponse) error {
	// 模拟数据库查询：根据 UserID 返回用户信息
	switch req.UserID {
	case 1001:
		*res = GetUserResponse{
			ID:       1001,
			Name:     "Alice",
			Age:      25,
			CreateAt: time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC),
		}
	case 1002:
		*res = GetUserResponse{
			ID:       1002,
			Name:     "Bob",
			Age:      30,
			CreateAt: time.Date(2024, 2, 20, 14, 15, 0, 0, time.UTC),
		}
	default:
		return fmt.Errorf("用户ID %d 不存在", req.UserID)
	}
	return nil // 无错误返回 nil
}

func main() {
	// 4. 注册 RPC 服务：将 UserService 实例注册为 "UserService" 服务
	userService := new(UserService)
	err := rpc.RegisterName("UserService", userService)
	if err != nil {
		log.Fatalf("注册服务失败：%v", err)
	}

	// 5. 启动 TCP 监听（端口 9090，Gob 基于 TCP 传输）
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("启动监听失败：%v", err)
	}
	defer listener.Close()
	log.Println("Gob 传输的 RPC 服务已启动，地址：tcp://localhost:9090")

	// 6. 循环处理客户端请求（阻塞运行）
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("接收连接失败：%v", err)
			continue
		}
		// 每个连接启动一个 goroutine 处理（支持并发）
		go rpc.ServeConn(conn)
	}
}