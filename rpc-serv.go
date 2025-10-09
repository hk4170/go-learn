// server.go（RPC 服务端代码）
package main

import (
	"net"
	"net/http"
	"net/rpc"
)

// 1. 定义 RPC 服务结构体（承载要暴露的函数）
type Calculator struct{}

// 2. 定义请求/响应参数结构体（字段首字母大写，确保可序列化）
type AddRequest struct {
	A int // 第一个加数
	B int // 第二个加数
}

type AddResponse struct {
	Result int // 相加结果
}

// 3. 实现 RPC 服务函数（必须满足 RPC 函数规则）
// 功能：计算两个数的和
func (c *Calculator) Add(req AddRequest, res *AddResponse) error {
	res.Result = req.A + req.B
	return nil // 无错误返回 nil
}

func main() {
	// 4. 注册 RPC 服务：将 Calculator 实例注册为名为 "Calculator" 的服务
	calc := new(Calculator)
	err := rpc.RegisterName("Calculator", calc)
	if err != nil {
		panic("注册 RPC 服务失败：" + err.Error())
	}

	// 5. 将 RPC 服务绑定到 HTTP 传输（方便客户端通过 HTTP 调用）
	rpc.HandleHTTP()

	// 6. 启动服务监听（端口设为 8080）
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("启动监听失败：" + err.Error())
	}
	defer listener.Close()

	println("RPC 服务已启动，地址：localhost:8080")
	// 7. 处理 HTTP 请求（阻塞运行）
	err = http.Serve(listener, nil)
	if err != nil {
		panic("服务运行失败：" + err.Error())
	}
}