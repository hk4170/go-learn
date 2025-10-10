package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	go func(){
		fmt.Printf("hello2")//后台运行 不输出
	}()
	fmt.Println("hello3")
}