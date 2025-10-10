package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// 定义要扫描的目标，比如扫描本机开放的常用端口
	target := "127.0.0.1"
	// 也可以扫描一个网段，如 "192.168.1.0/24"

	// 构造 nmap 命令，例如扫描常用端口
	cmd := exec.Command("nmap", "-A", target) // -F 表示快速扫描常用端口

	// 捕获命令的标准输出
	var out bytes.Buffer
	cmd.Stdout = &out

	// 执行命令
	err := cmd.Run()
	if err != nil {
		log.Fatalf("执行 nmap 失败: %v", err)
	}

	// 打印 nmap 输出结果
	fmt.Println(out.String())
}
