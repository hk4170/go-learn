package main

import (
	"fmt"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func main() {
	// 创建一个图
	s := op.NewScope()
	
	// 定义两个常量张量
	a := op.Const(s, int64(2))
	b := op.Const(s, int64(3))
	
	// 定义加法操作
	c := op.Add(s, a, b)
	
	// 创建图
	graph, err := s.Finalize()
	if err != nil {
		panic(err)
	}
	
	// 创建会话
	sess, err := tf.NewSession(graph, nil)
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	
	// 运行计算
	result, err := sess.Run(nil, []tf.Output{c}, nil)
	if err != nil {
		panic(err)
	}
	
	// 输出结果
	fmt.Printf("2 + 3 = %v\n", result[0].Value()) // 输出: 2 + 3 = 5
}
