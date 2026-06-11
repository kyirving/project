package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// JwtSecret JWT 密钥
	JwtSecret = "admin@123456"
)

type JwtManager struct {
	secret string
}

// CustomClaims 自定义声明
type CustomClaims struct {
	UserID               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 嵌套标准声明
}

func NewJwtManager() *JwtManager {
	return &JwtManager{
		secret: JwtSecret,
	}
}

func (m *JwtManager) GenerateToken(userID int64, username string) (string, time.Time, error) {

	duration := 24 * time.Hour
	now := time.Now()
	exp := now.Add(duration)

	// 设置声明
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),   // 过期时间
			IssuedAt:  jwt.NewNumericDate(now),   // 签发时间
			NotBefore: jwt.NewNumericDate(now),   // 生效时间
			Issuer:    "wuhao",                   // 签发者
			Subject:   fmt.Sprintf("%d", userID), // 主题
			ID:        "",                        // 建议使用UUID，用于后续黑名单功能
		},
	}

	// 选择签名算法并生成token字符串，这里使用HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(m.secret) // 生产环境必须从环境变量读取！
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", exp, err
	}
	return tokenString, exp, nil
}

func (m *JwtManager) ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 关键步骤：验证签名算法是否符合预期，防止alg=none攻击
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.secret), nil
	})

	// 处理解析错误
	if err != nil {
		return nil, err
	}

	// 验证token有效性并进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
