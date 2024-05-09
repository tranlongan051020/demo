package user

type UserUpdate struct {
	HashPassword   string  `json:"hash_password"    gorm:"column:hash_password"    validate:"min=1"`
	Name           string  `json:"name"             gorm:"column:name"             validate:"min=1"`
	Description    *string `json:"description"      gorm:"column:description"      validate:"min=1"`
	AvatarUrl      *string `json:"avatar_url"       gorm:"column:avatar_url"       validate:"min=1"`
	AvatarThumbUrl *string `json:"avatar_thumb_url" gorm:"column:avatar_thumb_url" validate:"min=1"`
}
