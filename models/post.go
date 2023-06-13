package models

type Post struct {
	PostID              int64  `json:"post_id" gorm:"primary_key;column:post_id"`
	UserID              int64  `json:"user_id" gorm:"column:user_id"`
	PostType            int8   `json:"post_type" gorm:"column:post_type"`
	PostCreatedDatetime string `json:"post_created_datetime" gorm:"column:post_created_datetime"`
	PostLike            int64  `json:"post_like" gorm:"column:post_like"`
	PostComment         int64  `json:"post_comment" gorm:"column:post_comment"`
	PostText            string `json:"post_text" gorm:"column:post_text"`
}
