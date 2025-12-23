package main

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/maxminddb-golang"
)

// 定义结构体，与 MMDB 中的字段映射（按需选择字段）
type GeoRecord struct {
	Country struct {
		ISOCode string `maxminddb:"iso_code"` // 国家 ISO 代码
		Name    string `maxminddb:"en"`       // 国家英文名称
	} `maxminddb:"country"`
	City struct {
		Name string `maxminddb:"en"` // 城市英文名称
	} `maxminddb:"city"`
	Location struct {
		Latitude  float64 `maxminddb:"latitude"`  // 纬度
		Longitude float64 `maxminddb:"longitude"` // 经度
	} `maxminddb:"location"`
}

func main() {
	// 1. 打开 MMDB 数据库文件
	db, err := maxminddb.Open("GeoIP.mmdb")
	if err != nil {
		log.Fatalf("打开数据库失败: %v", err)
	}
	defer db.Close() // 延迟关闭数据库

	// 2. 待查询的 IP 地址
	ip := net.ParseIP("223.160.229.115") // 示例 IP：Google DNS

	// 3. 初始化结果结构体
	var record GeoRecord

	// 4. 执行查询
	err = db.Lookup(ip, &record)
	if err != nil {
		log.Fatalf("查询 IP 失败: %v", err)
	}

	// 5. 输出查询结果
	fmt.Printf("IP: %s\n", ip)
	fmt.Printf("国家代码: %s\n", record.Country.ISOCode)
	fmt.Printf("国家名称: %s\n", record.Country.Name)
	fmt.Printf("城市名称: %s\n", record.City.Name)
	fmt.Printf("经纬度: %.4f, %.4f\n", record.Location.Latitude, record.Location.Longitude)
}