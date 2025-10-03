/*
Go 只有一种循环结构：for 循环。

基本的 for 循环由三部分组成，它们用分号隔开：

初始化语句：在第一次迭代前执行
条件表达式：在每次迭代前求值
后置语句：在每次迭代的结尾执行
初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。

一旦条件表达式求值为 false，循环迭代就会终止。

注意：和 C、Java、JavaScript 之类的语言不同，Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 { } 则是必须的。

初始化语句和后置语句是可选的。
func main() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}
此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。
func main() {
	for {
	}
}


*/
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