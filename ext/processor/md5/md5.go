// processor/md5/md5.go MD5 处理扩展
package md5

import (
    "crypto/md5"
    "encoding/hex"
    "ext/processor"
)

type MD5Processor struct{}

func (m MD5Processor) Name() string {
    return "md5"
}

func (m MD5Processor) Process(data string) (string, error) {
    h := md5.New()
    h.Write([]byte(data))
    return hex.EncodeToString(h.Sum(nil)), nil
}

// init 函数在包加载时自动注册扩展
func init() {
    processor.Register(MD5Processor{})
    println("md5 processor init")
}