package tag

import (
	"go-web/app/models"
	"go-web/app/models/user"
)

type Tag struct {
	models.BaseModel

	Name string `json:"name"`

	// 用户关联
	UserID string    `json:"user_id" gorm:"index"`
	User   user.User `json:"user"`

	models.CommonTimestampsField
	models.CommonSoftDeletesField
}