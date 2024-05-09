package music

type MusicUpdate struct {
	Name           string  `json:"name"             gorm:"column:name"`
	Description    *string `json:"description"      gorm:"column:description"`
	AvatarUrl      *string `json:"avatar_url"       gorm:"column:avatar_url"`
	AvatarThumbUrl *string `json:"avatar_thumb_url" gorm:"column:avatar_thumb_url"`
	OriginUrl      *string `json:"origin_url"       gorm:"column:origin_url"`
	LockKey        *string `json:"lock_key"         gorm:"column:lock_key"`
}

func (MusicUpdate) TableName() string {
	return TableName
}
