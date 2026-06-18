package router

import (
	"app/components/idgen"
	"app/components/logger"
	"app/internal/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, logManager *logger.LoggerManager, gen idgen.Generator) *gin.Engine {
	r := gin.Default()

	// Prometheus metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//存活侦探检查 判断容器是否还"活着"。如果探测失败，kubelet 会杀掉容器并根据 restartPolicy 重启
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Healthy",
		})
	})

	// 就绪探测 断容器是否准备好接收流量。如果探测失败，Service 的 Endpoint 会将该 Pod 摘除，不再分发请求给它， pod依然存在
	r.GET("/ready", func(c *gin.Context) {
		sqlDb, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}

		if err := sqlDb.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Ready",
		})
	})

	// 全局访问日志中间件
	r.Use(middleware.LoggerMiddleware(logManager.Access),
		middleware.MetricsMiddleware(),
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
