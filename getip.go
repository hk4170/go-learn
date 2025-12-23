package main

import (
	"fmt"
	"net"
)

func getOutboundIP() (net.IP, error) {
	// 连接谷歌DNS，仅用于获取出口IP，不会实际发送数据
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

func main() {
	ip, err := getOutboundIP()
	if err != nil {
		fmt.Printf("获取出口IP失败: %v\n", err)
		return
	}
	fmt.Printf("本机出口IP: %s\n", ip.String())
}