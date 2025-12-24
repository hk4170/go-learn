package main

import (
    "fmt"
    "github.com/yalue/onnxruntime_go"
  //  "unsafe"
)

func main() {
    // 1. 初始化 ONNX Runtime 环境
    env, err := onnxruntime_go.NewEnvironment()
    if err != nil {
        panic(err)
    }
    defer env.Destroy()

    // 2. 加载本地 ONNX 模型文件
    model, err := env.LoadModelFromFile("model.onnx")
    if err != nil {
        panic(err)
    }
    defer model.Destroy()

    // 3. 准备输入张量（示例：float32 类型，形状 [1, 3, 224, 224]）
    inputShape := []int64{1, 3, 224, 224}
    inputData := make([]float32, 1*3*224*224)
    // 填充输入数据（需替换为实际的预处理后数据）
    for i := range inputData {
        inputData[i] = 0.0
    }
    inputTensor, err := onnxruntime_go.NewTensor(inputShape, inputData)
    if err != nil {
        panic(err)
    }
    defer inputTensor.Destroy()

    // 4. 运行模型推理
    outputs, err := model.Run(map[string]onnxruntime_go.Tensor{"input_name": *inputTensor})
    if err != nil {
        panic(err)
    }
    defer func() {
        for _, v := range outputs {
            v.Destroy()
        }
    }()

    // 5. 解析输出结果
    outputTensor := outputs["output_name"]
    outputData := outputTensor.GetData().([]float32)
    fmt.Println("输出结果长度:", len(outputData))
}