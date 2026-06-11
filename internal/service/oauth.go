package service

import (
	"app/components/idgen"
	"app/components/jwt"
	"app/internal/model"
	"app/internal/repository"
	"app/utils/biz_error"
	"context"
	"time"
)

type OAuthService struct {
	oauthRepo *repository.OAuthRepository
	gen       idgen.Generator
}

func NewOAuthService(oauthRepo *repository.OAuthRepository, gen idgen.Generator) *OAuthService {
	return &OAuthService{oauthRepo: oauthRepo, gen: gen}
}

func (s *OAuthService) Login() {
	s.oauthRepo.Login()
}

type RegisterResult struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
	UserID uint64    `json:"user_id"`
}

func (s *OAuthService) Register(ctx context.Context, Username, Password string) (RegisterResult, error) {
	// 检查用户名是否存在
	userInfo, err := s.oauthRepo.FindByUsername(ctx, Username)
	if userInfo != nil {
		return RegisterResult{}, biz_error.ErrUserExist
	}

	uid := s.gen.NextID()
	// 先生成JWT
	jwtManager := jwt.NewJwtManager()
	token, expire, err := jwtManager.GenerateToken(int64(uid), Username)
	if err != nil {
		return RegisterResult{}, err
	}

	user := model.User{
		UserID:   uid,
		Username: Username,
		Password: Password,
	}
	err = s.oauthRepo.Register(ctx, &user)
	if err != nil {
		return RegisterResult{}, err
	}

	result := RegisterResult{
		Token:  token,
		Expire: expire,
		UserID: uid,
	}
	return result, nil
}
