package model

import (
	"fmt"
	"time"
)

type User struct {
	ID           uint
	UserID       uint64
	Username     string
	Password     string
	Mobile       string
	Nikename     string
	RegisterType int8
	Balance      float64
	IsDeleted    int8
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u User) TableName() string {
	return fmt.Sprintf("user_%d", u.UserID)
}

func (User) GetTableName(index uint) string {
	return fmt.Sprintf("user_%d", index)
}
