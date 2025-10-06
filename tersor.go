package main

import (
    "fmt"
    "github.com/gorgonia/tensor"
    "github.com/gorgonia/tensor/device"
)

func main() {
    // 1. 定义张量数据（float32 类型，GPU 计算常用该类型）
    data1 := []float32{1.0, 2.0, 3.0, 4.0}
    data2 := []float32{5.0, 6.0, 7.0, 8.0}

    // 2. 创建 2x2 形状的 GPU 张量（关键：指定 WithEngine(device.CUDADevice{})）
    tensor1, err := tensor.New(
        tensor.WithShape(2, 2),    // 张量形状：2行2列
        tensor.WithData(data1),    // 绑定数据
        tensor.WithEngine(device.CUDADevice{}), // 启用 GPU
    )
    if err != nil {
        panic(err)
    }

    tensor2, err := tensor.New(
        tensor.WithShape(2, 2),
        tensor.WithData(data2),
        tensor.WithEngine(device.CUDADevice{}),
    )
    if err != nil {
        panic(err)
    }

    // 3. 在 GPU 上执行张量加法
    result := tensor.New(tensor.WithEngine(device.CUDADevice{}))
    if err := tensor.Add(result, tensor1, tensor2); err != nil {
        panic(err)
    }

    // 4. 打印结果（数据会从 GPU 传回 CPU 显示）
    fmt.Println("GPU 张量加法结果：")
    fmt.Println(result)
}