package repository

import (
	"app/internal/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

const (
	TABLE  = "user_%d"
	CCOUNT = 16
)

type OAuthRepository struct {
	db *gorm.DB
}

func NewOAuthRepository(db *gorm.DB) *OAuthRepository {
	return &OAuthRepository{db: db}
}

func (r *OAuthRepository) Login() {
}

func (r *OAuthRepository) Register(ctx context.Context, user *model.User) error {
	tbName := TableName(user.UserID)

	userIdx := model.UserIndex{
		UserID:     user.UserID,
		UserName:   user.Username,
		IndexValue: TableSuffix(user.UserID),
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Create(&userIdx).Error; err != nil {
			// will rollback
			return err
		}

		// if create user success, then return nil will commit else rollback
		return tx.WithContext(ctx).Table(tbName).Create(user).Error
	})
}

func (r *OAuthRepository) FindByUsername(ctx context.Context, username string) (*model.UserIndex, error) {
	var userIdx model.UserIndex
	if err := r.db.WithContext(ctx).Where("user_name = ?", username).First(&userIdx).Error; err != nil {
		return nil, err
	}

	return &userIdx, nil
}

func TableName(id uint64) string {
	// 用id的高32位作为表名的索引
	// 低32位作为表名的索引
	return fmt.Sprintf("user_%d", (id^(id>>32))%16)
}

func TableSuffix(id uint64) int8 {
	return int8((id ^ (id >> 32)) % 16)
}
