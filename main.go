package fusionsms

import (
	"fmt"
	"fusionsms/middleware"
	"fusionsms/route"
	"github.com/gin-gonic/gin"
)

func main() {
	service := gin.Default()     // 创建engine
	service.Use(middleware.Core) // 设置中间件
	route.Core(service)          // 设置路由
	if err := service.Run(":8888"); err != nil {
		fmt.Println("启动失败，原因：", err)
	}
}
