package router

import (
	"net/http"

	"mygithub/Gin/gin-restful-learn/apiserver/handler/sd"
	"mygithub/Gin/gin-restful-learn/apiserver/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers. 加载路由
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares. 设置 HTTP Header
	// 在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，
	// 需要通过 gin.Recovery() 来恢复 API 服务器
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handlers
	// sd 分组主要用来检查 API Server 的状态：健康状况、服务器硬盘、CPU 和内存使用量
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
