//go:build darwin
package main
import (
	"fmt"
	"os/exec"
)

func notice(title, message string) error {
	script := fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)
	cmd := exec.Command("osascript", "-e", script)
	return cmd.Run()
}

func main() {
	err := notice("Hello from Go", "这是一个来自 Go 程序的通知！")
	if err != nil {
		fmt.Println("发送通知失败:", err)
	} else {
		fmt.Println("通知已发送")
	}
}
