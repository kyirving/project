package model

type UserIndex struct {
	ID         uint   `gorm:"primaryKey,autoIncrement"`
	UserID     uint64 `gorm:"uniqueIndex:uk_user_id"`
	UserName   string `gorm:"size:128;index:uk_user_name,unique"`
	IndexValue int8   `gorm:"default:0"`
}

func (u UserIndex) TableName() string {
	return "user_index"
}
