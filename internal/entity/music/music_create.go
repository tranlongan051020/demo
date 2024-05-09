package music

type MusicCreate struct {
	Name           string  `json:"name"             gorm:"column:name"`
	AuthorUserId   uint32  `json:"author_user_id"   gorm:"column:author_user_id"`
	ResourceUrl    string  `json:"resource_url"     gorm:"column:resource_url"`
	Description    *string `json:"description"      gorm:"column:description"`
	AvatarUrl      *string `json:"avatar_url"       gorm:"column:avatar_url"`
	AvatarThumbUrl *string `json:"avatar_thumb_url" gorm:"column:avatar_thumb_url"`
	OriginUrl      *string `json:"origin_url"       gorm:"column:origin_url"`
	LockKey        *string `json:"lock_key"         gorm:"column:lock_key"`
}

func (MusicCreate) TableName() string {
	return TableName
}
