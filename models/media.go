package models

type Media struct {
	ID        int64  `json:"id" gorm:"column:media_id"`
	Path      string `json:"path" gorm:"column:media_path"`
	OwnerID   int64  `json:"owner_id" gorm:"column:media_owner_id"`
	OwnerType int8   `json:"owner_type" gorm:"column:media_owner_type"`
}
