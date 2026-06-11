package handler

import (
	"app/utils/biz_error"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS        = 0
	FAILED         = 500
	BAD_REQUEST    = 400
	NOT_AUTHORIZED = 401
	NOT_FOUND      = 404
	NOT_ACCEPTABLE = 406
	INTERNAL_ERROR = 500
	NETWORK_ERROR  = 502
)

var respMsg = map[int]string{
	SUCCESS:        "ok",
	FAILED:         "failed",
	BAD_REQUEST:    "bad request",
	NOT_AUTHORIZED: "not authorized",
	NOT_FOUND:      "not found",
	NOT_ACCEPTABLE: "not acceptable",
	NETWORK_ERROR:  "network error",
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
type BaseHandler struct {
}

func (b BaseHandler) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    SUCCESS,
		Message: respMsg[SUCCESS],
		Data:    data,
	})
}

func (b BaseHandler) Error(c *gin.Context, code int) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: respMsg[code],
	})
}

func (b BaseHandler) BizError(c *gin.Context, err *biz_error.BizError) {
	c.JSON(http.StatusOK, Response{
		Code:    err.Code,
		Message: err.Message,
	})
}
