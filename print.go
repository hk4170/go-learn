/*
函数	             特点	                       是否格式化	   适用场景
fmt.Print/print	    直接输出内容，不换行，不格式化	   ❌ 不格式化	 简单输出，不需要格式控制
fmt.Println/println 直接输出内容，自动换行，不格式化   ❌ 不格式化	 简单输出，多个值之间用空格分隔，最后换行
fmt.Printf/printf	支持格式化字符串 不换行          ✅ 支持格式化	需要按指定格式输出，比如控制数字、对齐、浮点精度等

核心重点：格式化输出 可以输出到io  例如./test > 1.txt
fmt.Printf
使用 格式化字符串 来控制输出的样式，后面可以跟若干个参数，按照格式字符串中的占位符依次填充。

fmt.Fprintf(os.Stdout, "Name: %s\n", "Alice")

格式化字符串 但是不输出
fmt.Sprintf
基本语法：

fmt.Printf("格式化字符串", 参数1, 参数2, ...)
占位符	说明	示例
\n  换行
%v	默认格式的值	fmt.Printf("%v", 42) → 42
%+v	打印结构体时带上字段名	适用于结构体
%#v	Go 语法表示的值	如 main.Person{Name:"Alice"}
%T	值的类型	fmt.Printf("%T", 42) → int
%d	十进制整数	fmt.Printf("%d", 42) → 42
%b	二进制	fmt.Printf("%b", 42) → 101010
%o	八进制	fmt.Printf("%o", 42) → 52
%x, %X	十六进制（小写/大写）	%x: 2a, %X: 2A
%f	浮点数（默认精度）	fmt.Printf("%f", 3.14159) → 3.141590
%F	同 %f
%e, %E	科学计数法（小写 e / 大写 E）	%e: 3.141590e+00
%s	字符串	fmt.Printf("%s", "hi") → hi
%q	带双引号的字符串（适合打印字符串字面量）	"hi" → "hi"
%c	Unicode 字符（ASCII / rune）	'A' → A
%p	指针地址	fmt.Printf("%p", &x) → 0xc0000...
%%	打印一个百分号 %	fmt.Printf("%%") → %
*/
package main

import (
	"fmt"

)

//import "os"

func main() {
    name := "Alice"
    age := 30
    height := 1.65

    // 使用 Printf 格式化输出                 %.2f 1.65
    fmt.Printf("Name: %s, Age: %v, Height: %.1f\n", name, age, height)
	x := 42
    y := 3.14
    z := "hello"
    println("打印值跟类型")//不会输出到

    fmt.Printf("x 的值: %v, 类型: %T\n", x, x)
    fmt.Printf("y 的值: %v, 类型: %T\n", y, y)
    fmt.Printf("z 的值: %v, 类型: %T\n", z, z)
    println("打印结构体")
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Bob", Age: 25}

    fmt.Printf("%v\n", p)    // {Bob 25}
    fmt.Printf("%+v\n", p)   // {Name:Bob Age:25}
    fmt.Printf("%#v\n", p)   // main.Person{Name:"Bob", Age:25}
	println("格式化支持输出到io: go run print.go > 1.txt")
	//fmt.Fprintf(os.Stdout, "Name: %s\n", "Alice")

}
