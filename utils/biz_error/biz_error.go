package biz_error

type BizError struct {
	Code    int
	Message string
}

// 预定义错误
var (
	ErrUserExist        = &BizError{Code: 400, Message: "用户已存在"}
	ErrUserNotExist     = &BizError{Code: 404, Message: "用户不存在"}
	ErrNotAuthenticated = &BizError{Code: 401, Message: "未认证"}
	ErrNotFound         = &BizError{Code: 404, Message: "未找到资源"}
)

func NewBizError(code int, message string) *BizError {
	return &BizError{Code: code, Message: message}
}

func (e *BizError) Error() string {
	return e.Message
}
