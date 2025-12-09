// ext/sha1.go
package ext

import (
    "crypto/sha1"
    "encoding/hex"
)

type SHA1Handler struct{}

func (s SHA1Handler) Name() string {
    return "sha1"
}

func (s SHA1Handler) Handle(data string) string {
    h := sha1.New()
    h.Write([]byte(data))
    return hex.EncodeToString(h.Sum(nil))
}

// init 函数：包加载时自动注册扩展
func init() {
    Register(SHA1Handler{})
}