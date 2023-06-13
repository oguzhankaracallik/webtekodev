package models

import "time"

type User struct {
	ID          int64     `json:"id" gorm:"primary_key;column:user_id"`
	Username    string    `json:"username" gorm:"column:user_name"`
	Name        string    `json:"name" gorm:"column:name"`
	Surname     string    `json:"surname" gorm:"column:surname"`
	UserMail    string    `json:"user_mail" gorm:"column:user_mail"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time"`
}
