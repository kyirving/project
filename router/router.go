package router

import (
	"app/components/idgen"
	"app/components/logger"
	"app/internal/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, logManager *logger.LoggerManager, gen idgen.Generator) *gin.Engine {
	r := gin.Default()
	// 全局访问日志中间件
	r.Use(middleware.LoggerMiddleware(logManager.Access),
		middleware.ExceptionMiddleware(),
		middleware.TimeoutMiddleware(5*time.Second),
		middleware.LimiterMiddleware(),
		middleware.CORSMiddleware(),
	)

	// 方法不存在处理
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    http.StatusMethodNotAllowed,
			"message": "Method Not Allowed",
		})
	})

	// 全局路由不存在处理
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Not Found",
		})
	})

	RegisterUserRouter(r, db, logManager, gen)
	return r
}
