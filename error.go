package main


func risky() {
	//定义错误处理函数
    defer func() {
		println("处理错误...")
        if err := recover(); err != nil {
            println("恢复恐慌",err)
		}
    }()

    panic("发生严重错误") // 触发恐慌
}

func main() {
    risky()
    println("程序继续运行") // 会执行
}