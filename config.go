package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host string "yaml: host"
	Port int "yaml: port"
	Proxy string "yaml: proxy"
}

var defaultConfig = Config{
   Host : "localhost",
   Port : 8080,
   Proxy: "",
}

func confRead(configfile string ) Config {
	var config Config
	file, err := os.ReadFile(configfile)
    if err != nil {
        config  = defaultConfig
    } else{
		yaml.Unmarshal(file,&config)
	}
	return config
}

func confWrite(configfile string, config Config){
    yamldata ,_:= yaml.Marshal(&config)
	os.WriteFile(configfile,yamldata,0644)

}

func main(){
	conf := confRead("config.yaml")
	confWrite("config.yaml",conf)
	println("host:",conf.Host)
	println("port:",conf.Port)
	println("proxy:",conf.Proxy)
}


