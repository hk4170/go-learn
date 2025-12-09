// ext/ext.go
package ext

import "fmt"

// DataHandler 所有扩展必须实现的接口
type DataHandler interface {
    Name() string       // 扩展名称（如 md5/sha1）
    Handle(data string) string // 核心处理逻辑
}

var handlers = make(map[string]DataHandler) // 全局扩展注册表

// Register 供扩展调用，注册自身
func Register(h DataHandler) {
    if _, exists := handlers[h.Name()]; exists {
        fmt.Printf("扩展 %s 已存在，跳过注册\n", h.Name())
        return
    }
    handlers[h.Name()] = h
}

// GetAllHandlers 获取所有已注册的扩展
func GetAllHandlers() map[string]DataHandler {
    return handlers
}

// GetHandlerByName 根据名称获取单个扩展
func GetHandlerByName(name string) (DataHandler, bool) {
    h, ok := handlers[name]
    return h, ok
}