package main

import (
	"fmt"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {
	// 加载模型
	modelPath := "./model" // 模型目录路径
	model, err := tf.LoadSavedModel(modelPath, []string{"serve"}, nil)
	if err != nil {
		panic(fmt.Sprintf("无法加载模型: %v", err))
	}
	defer model.Session.Close()

	// 准备输入数据（示例：创建一个形状为 [1, 2] 的浮点张量）
	input, err := tf.NewTensor([][]float32{{1.0, 2.0}})
	if err != nil {
		panic(err)
	}

	// 运行模型（需要根据实际模型的输入输出名称修改）
	output, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			model.Graph.Operation("input").Output(0): input,
		},
		[]tf.Output{
			model.Graph.Operation("output").Output(0),
		},
		nil,
	)
	if err != nil {
		panic(err)
	}

	// 输出结果
	fmt.Printf("预测结果: %v\n", output[0].Value())
}
