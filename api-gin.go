package main
import (
	"github.com/gin-gonic/gin"
	"os"
)

var apiport = os.Getenv("apiport")
type Message struct {
	Text string `json:"text"`
}

func api() {
	r := gin.Default() // 创建一个默认的 Gin 路由，包含 Logger 和 Recovery 中间件

	// 定义路由和处理函数
	r.GET("/hello", func(c *gin.Context) {
		// 返回 JSON 响应
		c.JSON(200, Message{Text: "Hello, 世界！这是 Gin 框架的 Web API"})
	})

	// 启动服务，监听 :8080
	if apiport == "" {
		apiport = "8081"
	}
	addr := ":" + apiport
	r.Run(addr)
}
func main(){
	api()
}
