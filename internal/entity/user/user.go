package user

import (
	"time"
)

type User struct {
	Id             uint32    `json:"id"               gorm:"<-:create,column:id"`
	Email          string    `json:"email"            gorm:"<-:create,column:email"`
	Role           UserRole  `json:"role"             gorm:"<-:create,column:role"`
	HashPassword   string    `json:"hash_password"    gorm:"column:hash_password"`
	Name           string    `json:"name"             gorm:"column:name"`
	Description    *string   `json:"description"      gorm:"column:description"`
	AvatarUrl      *string   `json:"avatar_url"       gorm:"column:avatar_url"`
	AvatarThumbUrl *string   `json:"avatar_thumb_url" gorm:"column:avatar_thumb_url"`
	CreatedAt      time.Time `json:"created_at"       gorm:"<-:create;autoCreateTime:true;column:created_at"`
	UpdatedAt      time.Time `json:"updated_at"       gorm:"autoUpdateTime:true;column:updated_at"`
}

func (User) String() string {
	return "users"
}

func (User) TableName() string {
	return "users"
}
