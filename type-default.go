/*
没有明确初始化的变量声明会被赋予对应类型的 零值。

零值是：

数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）。
*/
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
/*
表达式 T(v) 将值 v 转换为类型 T。

一些数值类型的转换：

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
或者，更加简短的形式：

i := 42
f := float64(i)
u := uint(f)
与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。
试着移除例子中的 float64 或 uint 的类型转换，看看会发生什么。
