package main

import (
    "log"
    //"os"
    "time"

    "github.com/sevlyar/go-daemon"
)

func main() {
    // 1. 配置守护进程参数：日志、pid文件
    ctx := &daemon.Context{
        PidFileName: "app.pid",   // 记录进程ID，用于启停管理
        PidFilePerm: 0644,
        LogFileName: "app.log",   // 后台运行日志输出
        LogFilePerm: 0644,
        WorkDir:     "./",
        Umask:       027,
    }

    // 2. 启动守护进程
    d, err := ctx.Reborn()
    if err != nil {
        log.Fatalf("Failed to daemonize: %v", err)
    }
    if d != nil {
        // 父进程退出，子进程成为守护进程
        return
    }
    defer ctx.Release() // 程序退出时清理pid文件

    // 3. 你的业务逻辑（后台持续运行）
    log.Println("Daemon started successfully")
    for {
        log.Println("Running...", time.Now())
        time.Sleep(5 * time.Second)
    }
}



