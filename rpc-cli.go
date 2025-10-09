// client.go（RPC 客户端代码）
package main

import (
	"fmt"
	"net/rpc"
)

// 注意：客户端的请求/响应结构体，必须和服务端完全一致（字段名、类型、首字母大写）
type AddRequest struct {
	A int
	B int
}

type AddResponse struct {
	Result int
}

func main() {
	// 1. 连接 RPC 服务端（地址与服务端一致）
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		panic("连接 RPC 服务失败：" + err.Error())
	}
	defer client.Close()

	// 2. 构造请求参数
	req := AddRequest{A: 10, B: 20}
	var res AddResponse // 用于接收响应结果

	// 3. 调用服务端函数：格式为 "服务名.函数名"
	err = client.Call("Calculator.Add", req, &res)
	if err != nil {
		panic("调用 Add 函数失败：" + err.Error())
	}

	// 4. 打印结果
	fmt.Printf("10 + 20 = %d\n", res.Result) // 输出：10 + 20 = 30
}