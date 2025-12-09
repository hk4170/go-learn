// processor/processor.go 定义扩展接口
package processor

// DataProcessor 所有扩展必须实现的接口
type DataProcessor interface {
    Name() string       // 扩展名称
    Process(data string) (string, error) // 核心处理逻辑
}

// 全局注册器：管理所有扩展
var processors = make(map[string]DataProcessor)

// Register 供扩展调用，注册自身
func Register(p DataProcessor) {
    processors[p.Name()] = p
}

// GetProcessor 获取已注册的扩展
func GetProcessor(name string) (DataProcessor, bool) {
    p, ok := processors[name]
    return p, ok
}