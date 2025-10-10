package main

import (
	"fmt"
	"log"

	"github.com/Ullaakut/nmap"
)

func main() {
	// 创建一个 nmap 扫描器
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("127.0.0.1"),     // 目标 IP
		//nmap.WithPorts("22,80,443"), // 指定端口
		nmap.WithServiceInfo(),           // 获取服务信息
		nmap.WithOSDetection(),         // 可选：检测操作系统
	)
	if err != nil {
		log.Fatalf("无法创建 nmap 扫描器: %v", err)
	}

	// 执行扫描
	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("扫描失败: %v", err)
	}
	if warnings != nil {
		log.Printf("扫描警告: %v", warnings)
	}

	// 处理扫描结果
	for _, host := range result.Hosts {
		if len(host.Addresses) == 0 {
			continue
		}
		fmt.Printf("Host: %v\n", host.Addresses[0])

		for _, port := range host.Ports {
			if port.State.State == "open" {
				service := port.Service
				fmt.Printf("  Port %d/%s %s %s\n",
					port.ID,
					port.Protocol,
					service.Name,
					service.Product,
				)
			}
		}
	}
}
