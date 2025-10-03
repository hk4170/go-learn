/*
switch 语句是编写一连串 if - else 语句的简便方法。
它运行第一个 case 值 值等于条件表达式的子句。

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只会运行选定的 case，而非之后所有的 case。 
在效果上，Go 的做法相当于这些语言中为每个 case 后面自动添加了所需的 break 语句。在 Go 中，除非以 fallthrough 语句结束，否则分支会自动终止。 
Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不限于整数。
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
	"runtime"
)

func main() {
	fmt.Print("Go 运行的系统环境：")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}