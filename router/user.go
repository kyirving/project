package router

import (
	"app/components/idgen"
	"app/components/logger"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRouter(r *gin.Engine, db *gorm.DB, logManager *logger.LoggerManager, gen idgen.Generator) {
	oauthRepo := repository.NewOAuthRepository(db)
	oauthSvc := service.NewOAuthService(oauthRepo, gen)
	handler := handler.NewOAuthHandler(logManager, oauthSvc)

	oauth := r.Group("/oauth")
	oauth.POST("/login", handler.Login)
	oauth.POST("/register", handler.Register)
}
