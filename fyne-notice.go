package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/container"
)

func main() {
    myApp := app.New()
    win := myApp.NewWindow("Notification Demo")

    btn := widget.NewButton("Send Notification", func() {
        // 创建通知（标题+内容）
        note := fyne.NewNotification("Hello", "This is a Fyne notification!")
        // 发送系统通知
        myApp.SendNotification(note)
    })

    win.SetContent(container.NewVBox(btn))
    win.Resize(fyne.NewSize(300, 200))
    win.ShowAndRun()
}