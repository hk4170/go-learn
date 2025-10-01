package main

import (
	"fmt"
	"lib"
)

func main(){
	fmt.Println("hello this a workspace test ")
	fmt.Println("in workspace main dir you can run 'go run hello' run this  ")
	Test()
	fmt.Println(lib.SayHello("this a lib.hello test "))
}