//go:build windows
package main

import (
	"os/exec"
)

func notice(title, message string) error {
	psCommand := `New-BurntToastNotification -Text "` + title + `", "` + message + `"`
	cmd := exec.Command("powershell", "-Command", psCommand)
	return cmd.Run()
}

func main() {
	err := notice("Go 通知", "这是通过 PowerShell 发出的通知")
	if err != nil {
		println("发送失败:", err.Error())
	} else {
		println("已尝试发送通知")
	}
}
