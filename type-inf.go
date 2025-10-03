/*
在声明一个变量而不指定其类型时（即使用不带类型的 := 语法 var = 表达式语法），变量的类型会通过右值推断出来。

当声明的右值确定了类型时，新变量的类型与其相同：

var i int
j := i // j 也是一个 int
不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int、float64 或 complex128 了，这取决于常量的精度：

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
试着修改示例代码中 v 的初始值，并观察它是如何影响类型的。
*/
package main

import "fmt"

func main() {
	v := 56.7 // 修改这里看看！
	fmt.Printf("v is of type %T\n", v)
}
