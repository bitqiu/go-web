package article

import (
	"go-web/app/models"
	"go-web/app/models/tag"
	"go-web/app/models/user"
)

type Article struct {
	models.BaseModel

	Title    string `json:"title" gorm:"not null;index:idx_title_sub"`
	SubTitle string `json:"sub_title" gorm:"index:idx_title_sub"`
	Content  string `json:"content" gorm:"text"`

	// 用户关联
	UserID string    `json:"user_id" gorm:"index"`
	User   user.User `json:"user"`

	// 标签关联
	TagID string  `json:"tag_id" gorm:"index"`
	Tag   tag.Tag `json:"tag"`

	models.CommonTimestampsField
	models.CommonSoftDeletesField
}
