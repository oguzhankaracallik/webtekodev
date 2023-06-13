package models

import "time"

type Comment struct {
	CommentID   int       `json:"comment_id" gorm:"primary_key;column:comment_id"`
	UserID      int       `json:"user_id" gorm:"column:user_id"`
	PostID      int       `json:"post_id" gorm:"column:post_id"`
	CommentText string    `json:"comment_text" gorm:"column:comment_text"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"`
}
