package main

import (
	"encoding/gob"
	"log"
	"os"
)
type host struct {
	Host string
	Port int
	proxy string
	Proxy string
	Os string
	Os_version string
}
var hosts []host
func main(){
	if len(os.Args) != 2{
		log.Fatal("gob-encode <file>")
	}

	var asd host
	asd.Host = "localhost"
	asd.Port = 3321
	asd.Os = "linux"
	asd.Os_version = "ubuntu-2204"
	asd.proxy = "localhost"//首字母必须大写 否则无法访问
	asd.Proxy = "localhost"
    hosts = append(hosts, asd)

	file , err := os.Create(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encode := gob.NewEncoder(file)
	encode.Encode(hosts)
	println("Done")

	//fmt.Printf("data : %+v",asd)

}