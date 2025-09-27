package main

import (
  "fmt"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/widget"
)

func App() {
  myApp := app.New()               // 初始化APP
  myWindow := myApp.NewWindow("Fyne 测试窗口") // 创建窗口

  // 创建按钮：点击时打印日志
  btn := widget.NewButton("点击我", func() {
    fmt.Println("按钮被点击了！")
  })

  // 布局按钮并设置到窗口
  myWindow.SetContent(container.NewVBox(btn))
  myWindow.ShowAndRun() 
}