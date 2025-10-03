/*
Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( )，而大括号 { } 则是必须的。
和 for 一样，if 语句可以在条件表达式前执行一个简短语句。

该语句声明的变量作用域仅在 if 之内。


*/

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}