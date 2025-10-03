/*
switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。

例如，

switch i {
case 0:
case f():
}
在 i==0 时，f 不会被调用。）

*/


package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("周六是哪天？")
	today := time.Now().Weekday()
	day  := time.Now().Day()
	println(time.Now().Date())
	println(time.Saturday)
    println(day)
	switch time.Saturday {
	case today + 0:
		fmt.Println("今天。")
	case today + 1:
		fmt.Println("明天。")
	case today + 2:
		fmt.Println("后天。")
	default:
		fmt.Println("很多天后。")
	}
}
