package main

import (
	"encoding/gob"
	"log"
	"os"
)

func main(){
	type host struct {
		Host string
		Port int
		proxy string
		Proxy string
		Os string
	}
    var hosts []host
	var asd host
	asd.Host = "localhost"
	asd.Port = 3386
	asd.Os = "linux"
	asd.proxy = "localhost"//首字母必须大写 否则无法访问
	asd.Proxy = "localhost"
    hosts = append(hosts, asd)
	file , err := os.Create("test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encode := gob.NewEncoder(file)
	encode.Encode(hosts)
	println("Done")

	//fmt.Printf("data : %+v",asd)

}