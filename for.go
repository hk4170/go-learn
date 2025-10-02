package main

import "fmt"

var list []string 
func main(){
    list = append(list, "nihao ")
    list = append(list, "xinshijie")
    list = append(list, "nihao")
    list = append(list, "work")
    
	//test()
	test2()
	//test3()


}

func test2(){
	//go 写法
	//for v := range(list){//单个参数只能获取到i
	//for i,v := range(list){//获取i 跟v 
	for _,v := range list{//或者使用此方式获取v
		//fmt.Println(i)
		fmt.Println(v)
	}
}

func test(){//c++ 写法
	for i:=0; i<len(list); i++{
		fmt.Println(list[i])
	}
}

func test3(){
	var env int = 0
	for {//死循环
		//if env == 10{
			//break //也可以添加break退出
		//}
		println("run is ",env)
		env++
	}
}