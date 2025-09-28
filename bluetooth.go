package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/darwin" // macOS 专用
)

func scanDevices() {
	// 开始扫描（10秒后停止）
	err := ble.Scan(10 * time.Second, func(a ble.Advertisement) {
		fmt.Printf("发现设备: %s (RSSI: %d)\n", a.Addr(), a.RSSI())
		if len(a.LocalName()) > 0 {
			fmt.Printf("  设备名称: %s\n", a.LocalName())
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	adapter := darwin.NewDevice()
	ble.SetDefaultDevice(adapter)
	scanDevices()
}
