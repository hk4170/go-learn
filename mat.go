package main

import (
    "fmt"
    "gonum.org/v1/gonum/mat" // 引入矩阵模块
)

func main() {
    // 1. 创建两个 2x2 矩阵（使用 []float64 按“行优先”存储数据）
    // 矩阵 A: [[1 2], [3 4]]
    A := mat.NewDense(2, 2, []float64{1, 2, 3, 4})
    // 矩阵 B: [[5 6], [7 8]]
    B := mat.NewDense(2, 2, []float64{5, 6, 7, 8})

    // 2. 计算矩阵乘法 C = A * B（需先创建空矩阵存储结果，尺寸为 A.Rows x B.Cols）
    C := mat.NewDense(2, 2, nil) // nil 表示先不初始化数据
    C.Mul(A, B) // 执行乘法

    // 3. 打印结果（使用 mat.Formatted 格式化输出，更易读）
    fmt.Println("矩阵 A:")
    fmt.Println(mat.Formatted(A))

    fmt.Println("\n矩阵 B:")
    fmt.Println(mat.Formatted(B))

    fmt.Println("\n矩阵 C = A * B:")
    fmt.Println(mat.Formatted(C))

    // 4. 额外常用操作：计算矩阵 A 的行列式
    det := mat.Det(A)
    fmt.Printf("\n矩阵 A 的行列式: %.2f\n", det)
}