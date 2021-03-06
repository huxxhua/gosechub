package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {

	// 初始化 Gin 实例
	r := gin.Default()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {
		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World",
		})
	})

	// 处理404请求
	r.NoRoute(func(c *gin.Context) {

		// 获取标头信息的 Accept 信息
		accept := c.Request.Header.Get("Accept")

		if strings.Contains(accept, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON 格式
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

	err := r.Run()
	if err != nil {
		return
	}
}
