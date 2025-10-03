package main

import (
    "fmt"
    "log"
    "os"

    "gopkg.in/yaml.v3"
)

// 定义与 YAML 文件结构匹配的 Go 结构体
type Config struct {
    App      AppInfo      `yaml:"app"`
    Database DatabaseInfo `yaml:"database"`
    Features FeaturesInfo `yaml:"features"`
}

type AppInfo struct {
    Name string `yaml:"name"`
    Port int    `yaml:"port"`
}

type DatabaseInfo struct {
    Host     string `yaml:"host"`
    Port     int    `yaml:"port"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
}

type FeaturesInfo struct {
    EnableLogging bool `yaml:"enable_logging"`
    MaxUsers      int  `yaml:"max_users"`
}

func main() {
    // 准备一个配置对象
    config := Config{
        App: AppInfo{
            Name: "MyApp",
            Port: 3000,
        },
        Database: DatabaseInfo{
            Host:     "127.0.0.1",
            Port:     5432,
            Username: "dbuser",
            Password: "dbpass",
        },
        Features: FeaturesInfo{
            EnableLogging: true,
            MaxUsers:      500,
        },
    }

    // 1. 将结构体编码为 YAML 格式的 []byte
    yamlData, err := yaml.Marshal(&config)
    if err != nil {
        log.Fatalf("生成 YAML 失败: %v", err)
    }

    // 2. 写入到文件
    err = os.WriteFile("config.yaml", yamlData, 0644)
    if err != nil {
        log.Fatalf("写入 YAML 文件失败: %v", err)
    }

    fmt.Println("✅ YAML 文件已成功写入: config.yaml")
        // 1. 读取 YAML 文件内容
    
    data, err := os.ReadFile("config.yaml")
    if err != nil {
        log.Fatalf("无法读取 YAML 文件: %v", err)
    }
    
        // 2. 解析 YAML 到结构体
    //var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("解析 YAML 失败: %v", err)
    }
    
        // 3. 使用解析后的数据
    fmt.Printf("📦 应用名称: %s\n", config.App.Name)
    fmt.Printf("🔌 数据库地址: %s:%d\n", config.Database.Host, config.Database.Port)
    fmt.Printf("✨ 日志功能开启: %v\n", config.Features.EnableLogging)
}
    


