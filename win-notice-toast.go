//go:build windows
package main

import (
	"github.com/go-toast/toast"
)

func main() {
	notification := toast.Notification{
		AppID:   "MyGoApp", // 应用标识（建议固定）
		Title:   "你好，Windows！",
		Message: "这是一条来自 Go 程序的通知。",
		Icon:    "icon.png", // 可选，通知图标（放在程序同目录下，或全路径）
		Actions: []toast.Action{
			{Type: "protocol", Label: "查看详情", Arguments: "https://example.com"},
		},
	}

	// 发送通知
	err := notification.Push()
	if err != nil {
		println("发送通知失败:", err.Error())
	} else {
		println("通知已发送！")
	}
}
