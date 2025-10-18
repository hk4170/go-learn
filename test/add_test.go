package main

import "testing" // 必须导入 testing 包

// TestAdd 测试 Add 函数的正确性
// 函数名必须以 Test 开头，参数为 *testing.T 类型
func TestAdd(t *testing.T) {
    // 定义测试用例：输入（a,b）和期望输出
    testCases := []struct {
        name     string // 测试用例名称，便于定位问题
        a, b     int    // 输入参数
        expected int    // 期望结果
    }{
        {"positive numbers", 2, 3, 5},   // 正数相加 最后一个参数为验证结果
        {"negative numbers", -1, -2, -3}, // 负数相加
        {"error case", 0, 5, 6},     // 错误示例
    }

    // 遍历执行所有测试用例
    for _, tc := range testCases {
        // t.Run 用于分组测试，每个用例会单独显示结果
        t.Run(tc.name, func(t *testing.T) {
            result := Add(tc.a, tc.b)
            // 对比实际结果与期望结果
            if result != tc.expected {
                // t.Errorf 报告错误，但继续执行后续测试用例
                t.Errorf("Add(%d, %d) = %d, expected %d", tc.a, tc.b, result, tc.expected)
            }
        })
    }
}