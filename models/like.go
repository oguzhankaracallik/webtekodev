package models

type Like struct {
	ID     int64 `gorm:"primary_key;auto_increment;column:like_id" json:"like_id"`
	UserID int64 `gorm:"column:user_id" json:"user_id"`
	PostID int64 `gorm:"column:post_id" json:"post_id"`
}
