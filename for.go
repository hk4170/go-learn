package main

import "fmt"

func main(){
	var list []string 

	list = append(list, "nihao ")
	list = append(list, "xinshijie")
	list = append(list, "nihao")
	list = append(list, "work")
	for i:=0; i<len(list); i++{
		fmt.Println(list[i])
	}
//死循环
	var env int = 0
	for {
		if env == 10{
			break
		}
		println("run is ",env)
		env++
	}
}