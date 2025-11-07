package main
import "C"
import (
	"strconv"
	"strings"
)
// go build -buildmode=c-shared -o clib.so clib.go
// 注意：//export 必须紧贴函数，中间不能有空行！
// 功能：接收 C 字符串（逗号分隔的数字），返回计算后的总和（C.int 类型）
//export CalculateSum
func CalculateSum(cInput *C.char) C.int {
	// 1. 将 C 字符串转为 Go 字符串
	goInput := C.GoString(cInput)
	// 2. 分割字符串为数字列表（如 "10,20,30" → ["10","20","30"]）
	numStrs := strings.Split(goInput, ",")
	
	// 3. 计算总和
	sum := 0
	for _, s := range numStrs {
		num, err := strconv.Atoi(s)
		if err != nil {
			continue // 忽略非法数字
		}
		sum += num
	}
	
	// 4. 将 Go int 转为 C.int 返回
	return C.int(sum)
}

// 必须有 main 函数（即使空实现），否则编译动态库会报错
func main() {}