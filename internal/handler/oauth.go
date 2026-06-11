package handler

import (
	"app/components/logger"
	"app/internal/service"
	"app/utils/biz_error"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OAuthHandler struct {
	BaseHandler
	oauthSvc *service.OAuthService
	logger   *logger.LoggerManager
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewOAuthHandler(logManager *logger.LoggerManager, oauthSvc *service.OAuthService) *OAuthHandler {
	return &OAuthHandler{
		logger:      logManager,
		oauthSvc:    oauthSvc,
		BaseHandler: BaseHandler{},
	}
}

func (h *OAuthHandler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error.Error("login_req_bind_error", zap.Error(err))
		h.Error(c, BAD_REQUEST)
		return
	}

	h.Success(c, nil)
}

func (h *OAuthHandler) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error.Error("register_req_bind_error", zap.Error(err))
		h.Error(c, BAD_REQUEST)
		return
	}

	result, err := h.oauthSvc.Register(c.Request.Context(), req.Username, req.Password)
	if err == nil {
		h.Success(c, result)
		return
	}

	var bizErr *biz_error.BizError
	if errors.As(err, &bizErr) {
		h.BizError(c, bizErr)
		return
	}

	h.logger.Error.Error("register_error", zap.Error(err))
	h.Error(c, FAILED)
}
