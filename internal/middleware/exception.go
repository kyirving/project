package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 打印堆栈信息方便调试
				debug.PrintStack()
				// 构造内部错误响应
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "服务器内部错误",
				})
			}
		}()

		// 执行下一个处理器
		c.Next()

		// 检查是否有未处理的错误
		if len(c.Errors) > 0 {
			// 取最后一个错误（业务通常只抛出一个）
			err := c.Errors.Last().Err
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}
	}
}
