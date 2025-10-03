/*
什么是方法？

方法是依附于某个具体类型的函数。方法本质上也是函数，但它有一个“接收者”（receiver），这个接收者就是它所依附的类型。

你可以理解为：方法是绑定到某个类型的特殊函数。

方法的定义格式：
func (接收者) 方法名(参数列表) 返回值列表 {
    // 方法体
}
其中，(接收者) 是关键，它指定了这个方法是针对哪个类型的。

👉 在这个例子中：

• Area 是一个 方法，不是普通函数。

• 它的 接收者是 Rectangle 类型，即 (r Rectangle)。

• 所以，只有 Rectangle 类型的变量（比如 rect）才能调用这个 Area 方法。

• 方法调用形式是：变量.方法名()，如 rect.Area()。
*/


package main

import "fmt"

// 定义一个类型
type Rectangle struct {
    Width  float64
    Height float64
}

// 定义一个方法，接收者是 Rectangle 类型
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}


func main() {
    rect := Rectangle{Width: 10, Height: 5}
	rect.Width = 22
    fmt.Printf("Width: %v",rect.Width)
    println()
    fmt.Printf("Heght: %v",rect.Height)
    println()
    area := rect.Area() // 调用方法
    fmt.Println(area)   // 输出：50
}
