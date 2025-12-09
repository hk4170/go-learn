// main.go
package main

import (
    "fmt"
    _ "ext/processor/md5" // 导入扩展，下划线表示仅执行 init
    "ext/processor"
)

func main() {
    p, ok := processor.GetProcessor("md5")
    if !ok {
        fmt.Println("扩展不存在")
        return
    }
    res, _ := p.Process("hello")
    fmt.Println(res) // 输出 MD5 结果
}