package main

import "github.com/spf13/viper"
import "fmt"


// 嵌套结构体，对应 JSON 中的层级
type AppConfig struct {
	Name  string `mapstructure:"name"`  // 与 JSON 的 "name" 对应
	Port  int    `mapstructure:"port"`  // 与 JSON 的 "port" 对应
	Debug bool   `mapstructure:"debug"` // 与 JSON 的 "debug" 对应
}
  
type DBConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}
  
  // 总配置结构体，包含所有子配置
type TotalConfig struct {
	App AppConfig `mapstructure:"app"` // 对应 JSON 的 "app" 层级
	DB  DBConfig  `mapstructure:"db"`  // 对应 JSON 的 "db" 层级
}

func loadconfig() (TotalConfig, error) {
	var config TotalConfig
  
	// 1. 设置配置文件路径和名称（无需写后缀，后续指定格式）
	viper.SetConfigName("config")         // 配置文件名（不含后缀）
	viper.SetConfigType("json")           // 配置文件格式（json/yaml/toml 等）
	viper.AddConfigPath(".")              // 配置文件所在路径（"." 表示项目根目录）
  
	// 2. 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
	  return config, fmt.Errorf("读取配置失败：%v", err)
	}
  
	// 3. 将配置绑定到结构体
	if err := viper.Unmarshal(&config); err != nil {
	  return config, fmt.Errorf("解析配置失败：%v", err)
	}
  
	return config, nil
}

func main() {
	config , err := loadconfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)
	fmt.Println(config.App.Name)
	fmt.Println(config.DB.Host)

}