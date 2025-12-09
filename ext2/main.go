// main.go
package main

import (
    "ext2/ext" // 导入扩展包
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    // 1. 创建 Fyne 应用
    a := app.New()
    w := a.NewWindow("扩展式 GUI 工具")

    // 2. 获取所有自动注册的扩展，生成下拉选项
    handlers := ext.GetAllHandlers()
    options := make([]string, 0, len(handlers))
    for name := range handlers {
        options = append(options, name)
    }
    selectExt := widget.NewSelect(options, nil)
    selectExt.SetSelectedIndex(0) // 默认选中第一个扩展

    // 3. GUI 控件：输入框、按钮、结果框
    input := widget.NewEntry()
    input.SetPlaceHolder("请输入要处理的文本...")
    result := widget.NewLabel("处理结果会显示在这里")
    btn := widget.NewButton("执行处理", func() {
        // 点击按钮时，获取选中的扩展并调用
        selectedName := selectExt.Selected
        handler, ok := ext.GetHandlerByName(selectedName)
        if !ok {
            result.SetText("无效的扩展！")
            return
        }
        res := handler.Handle(input.Text)
        result.SetText(res)
    })

    // 4. 布局并显示窗口
    w.SetContent(container.NewVBox(
        selectExt,
        input,
        btn,
        result,
    ))
    w.Resize(fyne.NewSize(400, 300))
    w.ShowAndRun()
}