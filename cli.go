package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "greet-demo",
		Usage: "A demo for args and flag in urfave/cli",
		Commands: []*cli.Command{
			{
				Name:  "greet",
				Usage: "greet with a message and flag",
				// 定义 flag
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "suffix", // 长选项
						Aliases: []string{"s"}, // 短选项
						Value:   "!", // 默认值
						Usage:   "suffix of the greeting message",
					},
				},
				// 核心逻辑：获取位置参数 + flag
				Action: func(c *cli.Context) error {
					// 1. 获取位置参数（greet 后面的 hello）
					if c.Args().Len() == 0 {
						return fmt.Errorf("missing message argument, usage: greet [message] --suffix [value]")
					}
					// 获取第一个位置参数
					message := c.Args().Get(0)

					// 2. 获取 flag 的值
					suffix := c.String("suffix")

					// 输出结果
					fmt.Printf("Greeting: %s %s\n", message, suffix)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}