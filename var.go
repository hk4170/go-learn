/*
var 语句用于声明一系列变量。和函数的参数列表一样，类型在最后。

如例中所示，var 语句可以出现在包或函数的层级。
在函数中，短赋值语句 := 可在隐式确定类型的 var 声明中使用。

函数外的每个语句都 必须 以关键字开始（var、func 等），因此 := 结构不能在函数外使用

package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

*/
package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
