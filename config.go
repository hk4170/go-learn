package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Host string "yaml: host"
	Port int "yaml: port"
	Proxy string "yaml: proxy"
}

func config(configfile string ) Server {
	var config Server
	file, err := os.ReadFile(configfile)
    if err != nil {
        config  = Server{
			Host: "127.0.0.1",
			Port: 8080,
			Proxy: "127.0.0.1:8081",
		}
		
    } else{
		yaml.Unmarshal(file,&config)
	}
	yamlData, _ := yaml.Marshal(&config)
	os.WriteFile(configfile, yamlData, 0644)
	return config
}
func main(){
	conf := config("config.yaml")
	fmt.Println(conf.Host)
}


