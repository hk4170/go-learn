package main

import (
	"fmt"
	"os"
)


func main(){
	uid := os.Getuid()
	fmt.Println(uid)

	fmt.Println(os.Hostname())

	fmt.Println(os.Getwd())
    
    fmt.Println(os.Getenv("home"))

	os.Mkdir("test",0777)
	
	file ,_ := os.OpenFile("config.json", os.O_CREATE, 0777)
	file.Read([]byte("test"))
    fmt.Println(file.Name())
	
}