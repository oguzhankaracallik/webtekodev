package models

type Follower struct {
	Id             int `json:"id" gorm:"primary_key;column:follower_id"`
	FollowerUserId int `json:"follower_user_id" gorm:"column:follower_user_id"`
	FollowedUserId int `json:"followed_user_id" gorm:"column:followed_user_id"`
}
